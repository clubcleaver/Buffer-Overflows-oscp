package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//define Target IP and Port
	IP := "192.168.1.105" //change this
	PORT := "1985" //change this
	
	//define buffer
	bufferAlphabet := "A"
	maxToSend := 500 //max number of bytes to send
	minToSend := 100 //minimum bytes to send
	
	//Establish Connection
	Target := IP + ":" + PORT
	conn, err := net.Dial("tcp", Target)
	if err != nil {
		fmt.Println("could not connect ...")
		os.Exit(0)
	}
	//Loop to send bytes
	i := minToSend
	for i <= maxToSend {
		buffer := strings.Repeat(bufferAlphabet, i)
		finalBuffer := buffer + "\x00" //the return ending character was added for a specific usecase. can be removed if needs be.
		bytesWritten, err := conn.Write([]byte(finalBuffer))
		if err != nil {
			fmt.Println("could not write payload")
		}
		fmt.Println("Bytes Written: ", bytesWritten)
		time.Sleep(5 * time.Second)
		i += 100 //change the increment value if needed here
	}
}
