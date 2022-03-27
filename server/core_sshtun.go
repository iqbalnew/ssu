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
	tunnel := sshtunnel.NewSSHTunnel(
		// User and host of tunnel server, it will default to port 22
		// if not specified.
		fmt.Sprintf("%s@%s", getEnv("SSH1_USER", ""), getEnv("SSH1_HOST", "")),

		// Pick ONE of the following authentication methods:
		// sshtunnel.PrivateKeyFile("path/to/private/key.pem"), // 1. private key
		ssh.Password(getEnv("SSH1_PASSWORD", "")), // 2. password

		// The destination host and port of the actual server.
		getEnv("SSH1_DEST1", ""),

		"5432",
	)

	tunnel.Log = log.New(os.Stdout, "Tunnel 1: ", log.Ldate|log.Lmicroseconds)
	go tunnel.Start()
	time.Sleep(1000 * time.Millisecond)

	logrus.Println(tunnel.Local.Host)
	logrus.Println(tunnel.Local.Port)

	return tunnel.Local.Port

}
