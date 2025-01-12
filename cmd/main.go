package main

import (
	"log"
	"net"
	"path/filepath"

	"github.com/kurochkinivan/ftp_server/ftp"
)

const (
	hostname = "127.0.0.1"
	port     = "8080"
	rootDir  = "../public"
)

func main() {
	ln, err := net.Listen("tcp", net.JoinHostPort(hostname, port))
	if err != nil {
		log.Fatalf("failed to create listener, err: %v", err)
	}
	log.Println("server started on", ln.Addr().String())
	log.Printf("server is ready to accept connections")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept connection, err: %v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	abs, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatalf("failed to get abs path, err: %v", err)
	}
	ftpConn := ftp.NewConn(conn, abs)
	ftp.Serve(ftpConn)
}
