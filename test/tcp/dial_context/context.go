package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	//deadline
	dl := time.Now().Add(5 * time.Second)

	//new cotext
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer

	// delay time
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}

	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		conn.Close()
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

	if ctx.Err() != context.DeadlineExceeded {
		log.Printf("expected deadline exceeded; actual: %v", ctx.Err())
	}
}
