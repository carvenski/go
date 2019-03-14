package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		panic(msg)
	}
}

func callHTTPHandler(msg amqp.Delivery) (body []byte) {
	defer func() {
		// if exception:
		if err := recover(); err != nil {
			//don't use log.Fatalf() !! it will exit after log...
			log.Printf("====> [ERROR] in go callHTTPHandler: %s <====", err)

			//decode json in go
			var m map[string]interface{}
			json.Unmarshal(msg.Body, &m)

			requestId := m["requestId"].(string)
			requestCode := m["requestCode"].(string)
			providerId := m["providerId"].(string)

			x := `{"code": 3, "msg": "call HTTP Exception in go !", "result": ""`
			y := `,"requestId":"` + requestId + `","requestCode":"` + requestCode + `","providerId":"` + providerId + `"}`
			body = []byte(x + y) //can't use := here, cause it will make body local

		}
	}()

	// call python handler by HTTP GET
	url := "http://localhost:9191?req=" + string(msg.Body)
	resp, err := http.Get(url)
	failOnError(err, "Failed to connect http://localhost:9191")
	body, err = ioutil.ReadAll(resp.Body)
	failOnError(err, "Failed to read HTTP body")
	defer resp.Body.Close()

	return body

}

func handleMessage(msg amqp.Delivery, ch *amqp.Channel) {
	log.Printf("handling message: %d ...", msg.DeliveryTag)

	// call handler
	body := callHTTPHandler(msg)

	// send message
	log.Printf("send message: %d ...", msg.DeliveryTag)
	err := ch.Publish(
		"itom_resource_backend_exchange",   // exchange
		"itom_resource_backend_routingkey", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	// ack
	log.Printf("ack message: %d ...", msg.DeliveryTag)
	msg.Ack(false)

	log.Printf("#### End message: %d ####", msg.DeliveryTag)
}

var connURL string

func Init() {
	// add -connURL arg:
	// go run compute/vmvare/main.go -connURL "amqp://admin:admin@localhost:5672/%2F"
	flag.StringVar(&connURL, "connURL", "amqp://admin:admin@localhost:5672/%2F", "rabbitmq connURL: amqp://admin:admin@localhost:5672/%2F")
}

func main() {
	Init()
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile) //add log lineno

	log.Printf("connecting RabbitMQ: %s", connURL)
	conn, err := amqp.Dial(connURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"itom_resource_front_exchange", // name of the exchange
		"direct",                       // type
		true,                           // durable
		false,                          // delete when complete
		false,                          // internal
		false,                          // noWait
		nil,                            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	q, err := ch.QueueDeclare(
		"itom_resource_front_queue", // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // name of the queue
		"itom_resource_front_routingkey", // bindingKey
		"itom_resource_front_exchange",   //exchange
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to bind queue to exchange")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf("[****] Waiting for messages...")
	for msg := range msgs {
		log.Printf("#### Received message: %d ####\n%s ", msg.DeliveryTag, msg.Body)
		go handleMessage(msg, ch)
	}
	log.Printf("[****] Stopping consumer...")
}

