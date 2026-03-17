package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type HealthStatus struct {
	Redis    string `json:"redis"`
	RabbitMQ string `json:"rabbitmq"`
	MQTT     string `json:"mqtt"`
}

func checkRedis(ctx context.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()
	return rdb.Ping(ctx).Err()
}

func checkRabbitMQ() error {
	// RabbitMQ driver doesn't natively use context for Dial, so we manual timeout
	conn, err := amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{
		Dial: amqp.DefaultDial(2 * time.Second),
	})
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func checkMQTT() error {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetConnectTimeout(2 * time.Second)
	client := mqtt.NewClient(opts)

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	client.Disconnect(250)
	return nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Using a map for quick status gathering
	statuses := make(map[string]string)
	isAllUp := true

	// Check Redis
	if err := checkRedis(ctx); err != nil {
		statuses["redis"] = "DOWN"
		isAllUp = false
	} else {
		statuses["redis"] = "UP"
	}

	// Check RabbitMQ
	if err := checkRabbitMQ(); err != nil {
		statuses["rabbitmq"] = "DOWN"
		isAllUp = false
	} else {
		statuses["rabbitmq"] = "UP"
	}

	// Check MQTT
	if err := checkMQTT(); err != nil {
		statuses["mqtt"] = "DOWN"
		isAllUp = false
	} else {
		statuses["mqtt"] = "UP"
	}

	// Response Logic
	w.Header().Set("Content-Type", "application/json")
	if !isAllUp {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	fmt.Fprintf(w, `{"redis":"%s", "rabbitmq":"%s", "mqtt":"%s"}`,
		statuses["redis"], statuses["rabbitmq"], statuses["mqtt"])
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Health check service running on :8080...")
	http.ListenAndServe(":8080", nil)
}
