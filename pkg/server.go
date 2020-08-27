package pkg

import (
	"fmt"
	"bufio"
	"net"
);

// Global Variable
const (
	IP = "127.0.0.01" // localhost adress
	PORT = 8080 // Port use
	PATHLOG = "./logs/logServer.txt" // Log path
);

// Manage Error 
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

func sendClientConnect(client net.Conn, allClient []net.Conn) {
	var message string = "[INFO] New connected. Welcome " + client.RemoteAddr().String() + "\n";
	for _, c := range allClient {
		c.Write([]byte(message))
	}
}

func Server() {
	SetupLogServer(PATHLOG); // Setup logs server
	var fdLog = OpenLogsFile(PATHLOG);
	defer fdLog.Close();
	fmt.Println("Start Server...");
	var clients []net.Conn; // Client Array

	// Listen Server
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, PORT));
	manageError(err, 0);

	// Accept Connection
	for {
		conn, err := ln.Accept();
		manageError(err, 0);
		sendClientConnect(conn, clients);
		clients = append(clients, conn);
		fmt.Println("Client connected : ", conn.RemoteAddr());
		// Display Client Info

		go func() {
			buf := bufio.NewReader(conn);
			for {
				name, err := buf.ReadString('\n');
				if (err != nil) {
					fmt.Println("Client disconnected.");
					break;
				}
				for _, c := range clients {
					c.Write([]byte(name))
				}
				WriteLogsServer(fdLog, name, conn.RemoteAddr().String());
			}
		}()
	}
}