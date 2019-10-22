package main

func main() {
	// consumer
	go initConsumer()

	// producer
	go initProducer()

	// document REST api
	initApi()
}
