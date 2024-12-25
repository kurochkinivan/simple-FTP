package ftp

import (
	"fmt"
	"log"
)

const (
	status150    = "150 File status okay; about to open data connection."
	status200    = "200 Command okay."
	status220    = "220 Service ready for new user."
	status221    = "221 Service closing control connection."
	status226    = "226 Closing data connection. Requested file action successful."
	status230    = "230 User logged in, proceed."
	status250    = "250 Directory successfully changed."
	status425    = "425 Can't open data connection."
	status426    = "426 Connection closed; transfer aborted."
	status501    = "501 Syntax error in parameters or arguments."
	status502    = "502 Command not implemented."
	status504    = "504 Cammand not implemented for that parameter."
	status550    = "550 Requested action not taken. File unavailable."
	status550msg = "550 %s"
)

func (c *Conn) sendData(s string) {
	log.Print(">> sending data")
	_, err := fmt.Fprint(c.activeConn, s, c.EOL())
	if err != nil {
		log.Printf("failed to send response, err: %v", err)
	}
}

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Printf("failed to send response, err: %v", err)
	}
}

func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binar:
		return "\n"
	default:
		return "\n"
	}
}
