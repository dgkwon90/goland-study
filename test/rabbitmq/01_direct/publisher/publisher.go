package publisher

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Pub struct {
	Url      string
	Name     string
	Exchange string
	Conn     *amqp.Connection
	Channel  *amqp.Channel
}

func NewPub(url, name, exchange string) *Pub {
	pub := new(Pub)
	pub.Url = url
	pub.Name = name
	pub.Exchange = exchange
	fmt.Printf("[%v] New Pub :%v\n", pub.Name, pub)
	return pub
}

func (p *Pub) Connection() error {
	conn, connErr := amqp.Dial(p.Url)
	if connErr != nil {
		fmt.Printf("[%v] Rebbit MQ Connection Fail %v\n", p.Name, connErr.Error())
		return connErr
	}

	p.Conn = conn
	return nil
}

func (p *Pub) OpenChannel() error {
	ch, chErr := p.Conn.Channel()
	if chErr != nil {
		fmt.Printf("[%v] Channel Fail %v\n", p.Name, chErr.Error())
		return chErr
	}
	p.Channel = ch
	return nil
}

func (p *Pub) Close() {
	if p.Channel != nil {
		p.Channel.Close()
	}
	if p.Conn != nil {
		p.Conn.Close()
	}
}

func (p *Pub) Publish(routingKey string, msgHeader map[string]interface{}, msgBody []byte) error {
	mandatory := false
	immediate := false
	publishErr := p.Channel.PublishWithContext(
		context.Background(), // context
		p.Exchange,           // exchange
		routingKey,           // routing key
		mandatory,            // mandatory
		immediate,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Headers:     msgHeader,
			Body:        msgBody,
		})
	if publishErr != nil {
		fmt.Printf("[%v] publish Error %v\n", p.Name, publishErr)
		return publishErr
	}
	fmt.Printf("[%v] push message: %v => \n", p.Name, msgHeader)
	return nil
}
