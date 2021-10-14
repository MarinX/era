package rpc

import "io"

type Conn struct {
	in  io.ReadCloser
	out io.Writer
}

func (c *Conn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *Conn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *Conn) Close() error {
	return c.in.Close()
}
