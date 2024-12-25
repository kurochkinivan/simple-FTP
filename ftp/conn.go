package ftp

import (
	"net"
	"path/filepath"
)

type Conn struct {
	conn       net.Conn
	activeConn net.Conn
	dataType   dataType
	workDir    string
	rootDir    string
	userName   string
}

func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:     conn,
		dataType: ascii,
		workDir:  ".",
		rootDir:  rootDir,
	}
}

func (c *Conn) getCurPath() string {
	return filepath.Join(c.rootDir, c.workDir)
}
