package pkg;

import (
	"fmt"
	"net"
	"bufio"
	"os"
);

// Global Variable
// const (
// 	IP = "127.0.0.01" // localhost adress
// 	PORT = 8080 // Port use
// );

// Manage Error 
func manageErrorClient(err error, errorType int) {
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

func Client() {
	// Connect to server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", IP, PORT));
	manageErrorClient(err, 0);
	// Infinite loop, where user can send message.
	for {
		// User Input
		reader := bufio.NewReader(os.Stdin);
		fmt.Print("User : ");
		input, err := reader.ReadString('\n');
		manageErrorClient(err, 0);
		// Sends message to server
		conn.Write([]byte(input));
		message, err := bufio.NewReader(conn).ReadString('\n');
		manageErrorClient(err, 0);
		// Display server message
		fmt.Println("Server : ", message);
	}
}