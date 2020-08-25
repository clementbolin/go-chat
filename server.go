package main;

import (
	"fmt"
	"net"
);

// Global Variable
const (
	IP = "127.0.0.01" // localhost adress
	PORT = 8080 // Port use
);

func manageError(err error, errorType int) {
	if (errorType == 0) {
		if (err != nil) {
			panic(err);
		}
	} else if (errorType == 1) {
		if (err != nil) {
			fmt.Println("Client disconnect");
		}
	}
}

func main() {
	fmt.Println("Start Server...");

	// Listen Server
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, PORT));
	manageError(err, 0);
	// Accept Connection
	conn, err := ln.Accept();
	manageError(err, 0);
	// Display Client Info
	fmt.Println("Client connected : ", conn.RemoteAddr());
	for {
		var buffer []byte = make([]byte, 4096); // 4096 is maximun size for message
		var message string = "";
		length, err := conn.Read(buffer); // Read Message
		message = string(buffer[:length]); // Convert message in String
		manageError(err, 1);
		fmt.Println("Mesage: ", message);
		conn.Write([]byte(message + "\n"));
	}
}