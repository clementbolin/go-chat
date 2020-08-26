package pkg

import (
	"os"
	"fmt"
);

// Create log File Server
func SetupLogServer(path string) bool {
	err := os.Mkdir("./logs", os.ModePerm);
	fd, err := os.Create(path);
	defer fd.Close();
	fmt.Println(err);
	if (err != nil) {
		fmt.Println("Error: ", err);
		return false;
	}
	return (true)
}

// Write log Server in "logs/logServer.txt"
func LogsServer() {
	fmt.Println("In Development");
}