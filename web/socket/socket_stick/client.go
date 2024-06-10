package main

import (
	"fmt"
	"go-tools/web/socket/socket_stick/proto"
	"net"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:30000")
    if err != nil {
    	fmt.Println("Failed to dial", err)
    	return
	}
	defer conn.Close()

    for i := 0; i < 20;i++ {
    	msg := `Hello, world, how are you`
    	encodeMsg, err := proto.Encode(msg)
    	if err != nil {
    		fmt.Println("Failed to encode msg", err)
    		return
		}
    	conn.Write(encodeMsg)
	}
}
