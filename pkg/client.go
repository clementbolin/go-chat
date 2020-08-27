package pkg;

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"sync"
);

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

func choiceName() string {
	scan := bufio.NewScanner(os.Stdin);
	fmt.Print("Choice your pseudo: ");
	scan.Scan();
	return (scan.Text())
}

func Client() {
	var wg sync.WaitGroup;
	var pseudo string = choiceName(); // pseudo user
	// Connect to server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", IP, PORT));
	manageErrorClient(err, 0);
	wg.Add(2)

	// Infinite loop, where user can send message.
	go func() {
		defer wg.Done()
		for {
			// User Input
			reader := bufio.NewReader(os.Stdin);
			fmt.Print("> ");
			input, err := reader.ReadString('\n');
			manageErrorClient(err, 0);
			input = pseudo + ": " + input;
			// Sends message to server
			conn.Write([]byte(input))
		}
	}()

	// Goroutine for receive message
	go func() {
		for {
	
			message, err := bufio.NewReader(conn).ReadString('\n');
			manageErrorClient(err, 0);
	
			// Display server message
			fmt.Println(message);
		}
	}();
	wg.Wait();
}
