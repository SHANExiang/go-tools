package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	//uri          = "amqp://openstack:A123456z@10.50.7.108:5672/"
	uri          = "amqp://openstack:a06NXQ6eyKUpohuKCRJNJtIkw3kCzthIR1A7ypiu@10.50.1.57:5672/"
	exchange     = "nova_test"
	exchangeType = "topic"
	queue        = "ecs_voneyun_topic.error"
	bindingKey   = ""
	consumerTag  = ""
)


func main() {
	_, err := NewConsumer(uri, exchange, exchangeType, queue, bindingKey, consumerTag)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("running forever")
	select {}

}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	//if err = c.channel.ExchangeDeclare(
	//	exchange,     // name of the exchange
	//	exchangeType, // type
	//	false,         // durable
	//	false,        // delete when complete
	//	false,        // internal
	//	false,        // noWait
	//	nil,          // arguments
	//); err != nil {
	//	return nil, fmt.Errorf("Exchange Declare: %s", err)
	//}
	//log.Printf("declared Exchange, declaring Queue %q", queueName)
	//queue, err := c.channel.QueueDeclare(
	//	queueName, // name of the queue
	//	false,      // durable
	//	false,     // delete when unused
	//	false,     // exclusive
	//	false,     // noWait
	//	nil,       // arguments
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("Queue Declare: %s", err)
	//}
	//
	//log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
	//	queue.Name, queue.Messages, queue.Consumers, key)
	//
	//if err = c.channel.QueueBind(
	//	queue.Name, // name of the queue
	//	key,        // bindingKey
	//	exchange,   // sourceExchange
	//	false,      // noWait
	//	nil,        // arguments
	//); err != nil {
	//	return nil, fmt.Errorf("Queue Bind: %s", err)
	//}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue, // name
		c.tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, c.done)

	return c, nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}


