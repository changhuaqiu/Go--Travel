package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "P", 8000, "端口号，默认为8000")
	flag.Parse()
	Port := ":" + strconv.Itoa(port)

	listener, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
