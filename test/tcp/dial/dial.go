package main

import (
	"io"
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

	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
			log.Println("1 done")
		}()

		for {
			// 수신 연결 요청을 수락
			conn, connErr := listener.Accept()
			if connErr != nil {
				log.Println(connErr)
			}

			// 해당 연결을 처리
			go func(c net.Conn) {
				defer func() {
					closeErr1 := c.Close()
					if closeErr1 != nil {
						log.Println(closeErr1)
					}
					done <- struct{}{}
					log.Println("2 done")
				}()

				buf := make([]byte, 1024)
				for {
					n, readErr := c.Read(buf)
					if readErr != nil {
						if readErr != io.EOF {
							log.Println(readErr)
						}
						return
					}
					log.Printf("received: %q\n", buf[:n])
				}
			}(conn)
		}
	}()

	clientConn, clientConnErr := net.Dial("tcp", listener.Addr().String())
	if clientConnErr != nil {
		log.Fatal(clientConnErr)
	}

	clientConn.Close()
	<-done

	listener.Close()
	<-done
	// // defer를 통해 자연스럽게 Close되도록 구현
	// defer func() {
	// 	closeErr0 := listener.Close()
	// 	if closeErr0 != nil {
	// 		log.Println(closeErr0)
	// 	}
	// }()
	// log.Printf("bound to %q", listener.Addr())

}
