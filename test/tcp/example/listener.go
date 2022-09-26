package main

import (
	"log"
	"net"
)

func main() {
	// tcp server 생성
	// port를 0으로 설정시 랜덤으로 지정
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	// defer를 통해 자연스럽게 Close되도록 구현
	defer func() {
		closeErr0 := listener.Close()
		if closeErr0 != nil {
			log.Println(closeErr0)
		}
	}()
	log.Printf("bound to %q", listener.Addr())

	for {
		// 수신 연결 요청을 수락
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Fatal(connErr)
		}

		// 해당 연결을 처리
		go func(c net.Conn) {
			defer func() {
				closeErr1 := c.Close()
				if closeErr1 != nil {
					log.Println(closeErr1)
				}
			}()
		}(conn)
	}
}
