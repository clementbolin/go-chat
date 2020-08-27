package main;

import (
	"os"
	"pkg/pkg"
);

func main()  {
	var args = os.Args[1:];
	if (len(args) != 1) {
		os.Exit(84);
	}
	if (args[0] == "--server") {
		pkg.Server();
	} else if (args[0] == "--client") {
		pkg.Client();
	}
}