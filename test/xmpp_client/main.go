package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"gosrc.io/xmpp"
	"gosrc.io/xmpp/stanza"
)

// test
var (
	ClientJid = "test-xmpp-client@localhost/XMPPConn1"
	TargetJid = "test-xmpp-client0000@localhost/XMPPConn1"
)

type ConnectionRequest struct {
	XMLName  xml.Name `xml:"urn:broadband-forum-org:cwmp:xmppConnReq-1-0 connectionRequest"`
	UserName string   `xml:"username,omitempty"`
	Password string   `xml:"password,omitempty"`
}

func (c ConnectionRequest) Namespace() string {
	return c.XMLName.Space
}

func (c ConnectionRequest) GetSet() *stanza.ResultSet {
	return nil
}
func init() {
	stanza.TypeRegistry.MapExtension(stanza.PKTIQ, xml.Name{Space: "urn:broadband-forum-org:cwmp:xmppConnReq-1-0"}, ConnectionRequest{})
}

func connectionRequest(client xmpp.Sender) {
	// Craft a roster request
	req, err := stanza.NewIQ(stanza.Attrs{
		Id:   "cr001",
		From: ClientJid,
		To:   TargetJid,
		Type: stanza.IQTypeGet,
		Lang: "en",
	})
	if err != nil {
		log.Println(err)
	}

	req.Payload = ConnectionRequest{
		XMLName:  xml.Name{Space: "urn:broadband-forum-org:cwmp:xmppConnReq-1-0"},
		UserName: "test-xmpp-client0000",
		Password: "test1234",
	}
	//req.Payload := .

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	c, err := client.SendIQ(ctx, req)
	if err != nil {
		log.Println(err)
	}

	// if err := ctx.Err(); err != nil {
	// 	fmt.Println("ERROR!", err)
	// 	return
	// }

	//for {
	select {
	case <-ctx.Done():
		log.Println("timeout")
	case serverResp := <-c:
		log.Println(serverResp)
	}
	//}
}

func main() {
	host := "localhost"
	pass := "test1234"
	domain := "localhost"
	config := xmpp.Config{
		TransportConfiguration: xmpp.TransportConfiguration{
			Address: host + ":" + "5222",
		},
		Jid:          "test-xmpp-client@" + domain,
		Credential:   xmpp.Password(pass),
		StreamLogger: os.Stdout,
		Insecure:     true,
		// TLSConfig: tls.Config{InsecureSkipVerify: true},
	}

	router := xmpp.NewRouter()
	router.HandleFunc("message", handleMessage)

	client, err := xmpp.NewClient(&config, router, errorHandler)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// If you pass the client to a connection manager, it will handle the reconnect policy
	// for you automatically.
	cm := xmpp.NewStreamManager(client, nil)
	go func() {
		log.Fatal(cm.Run())
	}()

	time.Sleep(time.Second * 5)
	connectionRequest(client)

	// ==========================
	// Client connection
	// if err = client.Connect(); err != nil {
	// 	log.Fatal(err)
	// }
	time.Sleep(time.Second * 5)
	client.Disconnect()
}

func handleMessage(s xmpp.Sender, p stanza.Packet) {
	msg, ok := p.(stanza.Message)
	if !ok {
		_, _ = fmt.Fprintf(os.Stdout, "Ignoring packet: %T\n", p)
		return
	}

	log.Printf("Body = %s - from = %s\n", msg.Body, msg.From)

	// _, _ = fmt.Fprintf(os.Stdout, "Body = %s - from = %s\n", msg.Body, msg.From)
	// reply := stanza.Message{Attrs: stanza.Attrs{To: msg.From}, Body: msg.Body}
	// _ = s.Send(reply)
}

func errorHandler(err error) {
	fmt.Println(err.Error())
}
