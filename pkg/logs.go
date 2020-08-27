package pkg

import (
	"os"
	"fmt"
	"time"
	"strings"
);

// Create log File Server
func SetupLogServer(path string) {
	if _, err := os.Stat(path); err == nil { return }
	err := os.Mkdir("./logs", os.ModePerm);
	fd, err := os.Create(path);
	defer fd.Close();
	if (err != nil) {
		fmt.Println("Error: ", err);
		return;
	}
}

func OpenLogsFile(path string) *os.File {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600);
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error, OpenFile function in logs.go\nError Content: %s", err));
		return nil;
	}
	return fd;
}

// Write log Server in "logs/logServer.txt"
func WriteLogsServer(fd *os.File, content string, user string) {
	var log string = "";
	var currentTime = time.Now();

	content = strings.Replace(content, "\n", "", 1);
	log = fmt.Sprintf("Time: [%s] User: [%s] content: [%s]\n", currentTime.String(), content, user); // Create log mesage
	fd.WriteString(log); // Write log message in file
}