package ftp

type dataType int

const (
	ascii dataType = iota
	binar
)

func (c *Conn) setDataType(dataType dataType) {
	c.dataType = dataType
}