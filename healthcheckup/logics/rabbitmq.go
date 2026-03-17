

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