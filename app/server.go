package main

import (
	"fmt"
	"net"
	"os"
)

const HTTP_VERSION = "HTTP/1.1"

type header struct {
	name  string
	value string
}
type statusLine struct {
	httpVersion string
	statusCode  string
	statusText  string
}

func response(headers []header, status statusLine, body string) []byte {
	status_data := HTTP_VERSION + " " + status.statusCode + " " + status.statusText
	header_data := ""
	body_data := ""

	resp := status_data + "\r\n" + header_data + "\r\n" + body_data
	return []byte(resp)
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(response([]header{}, statusLine{statusCode: "200", statusText: "OK"}, ""))
	if err != nil {
		fmt.Println("Error writing response", err.Error())
		os.Exit(1)
	}
}
