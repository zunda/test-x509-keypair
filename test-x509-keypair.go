package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
)

func usage() string {
	return fmt.Sprintf("usage: %s <server-certificate.pem> <private-key.pem>\ntries to load a x509 key pair with crypto/tls.\n", filepath.Base(os.Args[0]))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print(usage())
		os.Exit(1)
	}
	certFile := os.Args[1]
	pkeyFile := os.Args[2]

	// https://pkg.go.dev/crypto/tls#example-LoadX509KeyPair
	_, err := tls.LoadX509KeyPair(certFile, pkeyFile)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Print("OK\n")
	}
}
