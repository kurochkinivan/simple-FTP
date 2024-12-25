package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (c *Conn) list() {
	if c.activeConn == nil {
		c.respond(status425)
		return
	}
	c.setDataType(ascii)

	c.respond(status150)
	files, err := getFiles(c.getCurPath())
	if err != nil {
		c.respond(status550)
		return
	}
	if files == "" {
		files = "Directory is empty"
	}
	c.sendData(files)

	c.activeConn.Close()
	c.respond(status226)
}

func getFiles(filepath string) (string, error) {
	paths, err := os.ReadDir(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to getFiles, err: %w", err)
	}

	var builder strings.Builder
	for i, path := range paths {
		info, err := path.Info()
		if err != nil {
			log.Printf("Error getting info for %s: %v", path.Name(), err)
			continue
		}
		line := fmt.Sprintf("%-10s %10d %12s %s", info.Mode(), info.Size(), info.ModTime().Format("02-01-2006"), info.Name())
		builder.WriteString(line)

		if i < len(paths)-1 {
			builder.WriteString("\r\n")
		}
	}

	return builder.String(), nil
}

func (c *Conn) cwd(args []string) {
	newDir := args[0]

	newRelPath := filepath.Join(c.workDir, newDir)
	absPath := filepath.Join(c.rootDir, newRelPath)
	cleanPath := filepath.Clean(absPath)

	if !strings.HasPrefix(cleanPath, c.rootDir) {
		c.respond(fmt.Sprintf(status550msg, "Permission denied. Path traversal detected."))
		return
	}

	info, err := os.Stat(cleanPath)
	if os.IsNotExist(err) || !info.IsDir() {
		c.respond(fmt.Sprintf(status550msg, "Directory not found."))
		return
	}

	newRelPath, err = filepath.Rel(c.rootDir, cleanPath)
	if err != nil {
		c.respond(fmt.Sprintf(status550msg, "Failed to get relative path."))
		return
	}

	c.workDir = newRelPath
	c.respond(status250)
}

func (c *Conn) pwd() {
	msg := fmt.Sprintf("257 \"%s\" is the current directory", c.workDir)
	c.respond(msg)
}

func (c *Conn) retr(args []string) {
	c.respond(status150)

	// c.setDataType(binar)
	// defer c.setDataType(ascii)

	filename := filepath.Join(c.rootDir, c.workDir, args[0])

	f, err := os.Open(filename)
	if err != nil {
		c.respond(fmt.Sprintf(status550msg, "failed to open file"))
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		line := sc.Text()
		c.sendData(line)
	}
	// _, err = io.Copy(c.activeConn, f)
	// if err != nil {
	// 	c.respond(fmt.Sprintf(status550msg, "failed to copy file"))
	// 	return
	// }

	c.respond(status226)
}

func (c *Conn) size(args []string) {
	filename := args[0]

	filename = filepath.Join(c.rootDir, c.workDir, filename)

	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			c.respond("550 File not found.")
		} else {
			c.respond("450 Requested file action not taken.")
		}
		return
	}

	if info.IsDir() {
		c.respond("550 Not a plain file.")
		return
	}

	c.respond(fmt.Sprintf("213 %d", info.Size()))
}

func (c *Conn) port(args []string) {
	args = strings.Split(args[0], ",")
	hostArgs, portArgs := args[:4], args[4:]
	hostname := strings.Join(hostArgs, ".")

	highByte, _ := strconv.Atoi(portArgs[0])
	lowByte, _ := strconv.Atoi(portArgs[1])

	port := strconv.Itoa(highByte*256 + lowByte)

	activeConn, err := net.Dial("tcp", net.JoinHostPort(hostname, port))
	if err != nil {
		log.Printf("Failed to connect to client: %v", err)
		c.respond(status425)
		return
	}

	c.activeConn = activeConn

	c.respond(status200)
}

func (c *Conn) user(args []string) {
	c.userName = args[0]
	c.respond(status230)
}

func (c *Conn) quit() {
	c.respond(status221)
	c.conn.Close()
}
