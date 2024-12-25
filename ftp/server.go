package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "LIST":
			c.list()
		case "PWD":
			c.pwd()
		case "SIZE":
			c.size(args)
		case "CWD":
			c.cwd(args)
		case "RETR":
			c.retr(args)
		case "PORT":
			c.port(args)
		case "QUIT":
			c.quit()
		case "USER":
			c.user(args)
		case "SYST":
			c.respond("215 UNIX Type: L8")
		case "FEAT":
			c.respond("211-Features:")
			c.respond(" SIZE")   
			c.respond(" binary") 
			c.respond("211 End")
		default:
			c.respond(status502)
		}
	}
}
