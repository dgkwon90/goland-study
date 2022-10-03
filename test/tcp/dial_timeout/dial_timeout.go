package main

import (
	"log"
	"net"
	"syscall"
	"time"
)

func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	// net.Dialer 인터페이스의 Control 함수를 오버라이딩 한다.
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timeout",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}

	return d.Dial(network, address)
}

func main() {
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		c.Close()
		log.Println("connection did not time out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		log.Printf("not netError: %v", nErr)
	}
	if !nErr.Timeout() {
		log.Printf("error is not a timeout: %v", nErr)
	} else {
		log.Printf("timeout: %v", nErr)
	}
}
