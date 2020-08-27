package pkg;

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"sync"
	"strings"
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

// Choice Pseudo
func choiceName() string {
	var pseudo string = "";
	for (len(pseudo) == 0 || len(pseudo) > 20) {
		scan := bufio.NewScanner(os.Stdin);
	
		fmt.Print("Choice your pseudo: ");
		scan.Scan();
		pseudo = scan.Text();
	}
	return (pseudo)
}

func checkUserDisconnect(input string, pseudo string, conn net.Conn) {
	input = strings.Replace(input, "\n", "", 1);
	if input == "Quit" {
		conn.Write([]byte("\033[1;31m[INFO]\033[1;37m " + pseudo + " leave the room Bye Bye\n"));
		fmt.Println("Bye bye", pseudo, "see you later");
		os.Exit(0);
	}
}

func Client() {
	var wg sync.WaitGroup;
	var pseudo string = choiceName(); // pseudo user
	// Connect to server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", IP, PORT));
	manageErrorClient(err, 0);
	fmt.Printf("%s connected to the room\nWrite Quit for leave room\n\n", pseudo);
	wg.Add(2)

	// Infinite loop, where user can send message.
	go func() {
		defer wg.Done()
		for {
			// User Input
			reader := bufio.NewReader(os.Stdin);
			input, err := reader.ReadString('\n');
			manageErrorClient(err, 0);
			checkUserDisconnect(input, pseudo, conn);
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
			fmt.Print(message);
		}
	}();
	wg.Wait();
}
