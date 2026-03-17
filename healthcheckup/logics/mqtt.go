

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