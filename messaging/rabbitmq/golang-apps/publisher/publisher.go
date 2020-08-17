package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var rabbitHost = os.Getenv("RABBIT_HOST")
var rabbitPort = os.Getenv("RABBIT_PORT")
var rabbitUser = os.Getenv("RABBIT_USERNAME")
var rabbitPassword = os.Getenv("RABBIT_PASSWORD")

func main() {

	router := httprouter.New()

	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		submit(w, r, p)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func submit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	message := p.ByName("message")

	fmt.Println("Received message: " + message)

	conn, err := amqp.Dial("amqp://" + rabbitUser + ":" + rabbitPassword + "@" + rabbitHost + ":" + rabbitPort + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"publisher", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}

	fmt.Println("publish success!")
}
