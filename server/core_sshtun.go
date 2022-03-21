package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

func startSSHTunnel() int {
	logrus.Println(fmt.Sprintf("%s@%s", getEnv("SSH1_USER", ""), getEnv("SSH1_HOST", "")))
	tunnel1 := sshtunnel.NewSSHTunnel(
		// User and host of tunnel server, it will default to port 22
		// if not specified.
		fmt.Sprintf("%s@%s", getEnv("SSH1_USER", ""), getEnv("SSH1_HOST", "")),

		// Pick ONE of the following authentication methods:
		// sshtunnel.PrivateKeyFile("path/to/private/key.pem"), // 1. private key
		ssh.Password(getEnv("SSH1_PASSWORD", "")), // 2. password

		// The destination host and port of the actual server.
		getEnv("SSH1_DEST1", ""),

		"8443",
	)

	tunnel1.Log = log.New(os.Stdout, "Tunnel 1: ", log.Ldate|log.Lmicroseconds)
	go tunnel1.Start()
	time.Sleep(1000 * time.Millisecond)

	tunnel2 := sshtunnel.NewSSHTunnel(
		// User and host of tunnel server, it will default to port 22
		// if not specified.
		fmt.Sprintf("%s@%s:%d", getEnv("SSH2_USER", ""), getEnv("SSH2_HOST", ""), tunnel1.Local.Port),

		// Pick ONE of the following authentication methods:
		// sshtunnel.PrivateKeyFile("path/to/private/key.pem"), // 1. private key
		ssh.Password(getEnv("SSH2_PASSWORD", "")), // 2. password

		// The destination host and port of the actual server.
		getEnv("SSH1_DEST2", ""),

		"5000",
	)

	tunnel2.Log = log.New(os.Stdout, "Tunnel 2: ", log.Ldate|log.Lmicroseconds)

	go tunnel2.Start()
	time.Sleep(100 * time.Millisecond)

	logrus.Println(tunnel2.Local.Host)
	logrus.Println(tunnel2.Local.Port)

	return tunnel1.Local.Port

}
