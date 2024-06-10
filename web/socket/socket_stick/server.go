package main

import (
	"bufio"
	"fmt"
	"go-tools/web/socket/socket_stick/proto"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed", err)
			break
		}
		fmt.Println("received data:", msg)
	}
}

func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:30000")
    if err != nil {
    	fmt.Println("listen failed", err)
    	return
	}
    defer listen.Close()

    for {
    	conn, err := listen.Accept()
    	if err != nil {
    		fmt.Println("accept failed", err)
    		continue
		}
		go process(conn)
	}
}

