package check

import (
	"context"
	"net/http"
	"time"
	
)

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
