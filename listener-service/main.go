package main

import (
	"fmt"
	"listener-service/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go" //advance messaging protocol replaces community service
)

func main() {

	//try to connect to rubbitmq
	rabbitConn, err := connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	//start listening for the messages

	log.Println("Listening for and consuming rubbit mq messageges......")

	//create consumer
	consumer, err := event.NewConsumer(rabbitConn)

	if err != nil {
		panic(err)
	}

	//watch the queue and consume event
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

	if err != nil {
		log.Println(err)
	}

}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	//dont continue untill rubbit mq is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")

		if err != nil {
			fmt.Println("RubbitMQ not yet ready ....")
			counts++
		} else {
			log.Println("Connected to RubbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backingoff")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
