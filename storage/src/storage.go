package main

func main() {
	// consumer
	go initConsumer()

	// producer
	initProducer()
}
