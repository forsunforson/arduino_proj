package serial

import (
	"github.com/tarm/serial"
)

const (
	// EOF end signal
	EOF = byte(4)
)

// Port use port to read & write byte
type Port struct {
	p      *serial.Port
	c      chan byte
	buffer []byte
}

// NewPort input serial address and baud rate, return a Port pointer
func NewPort(name string, baud int) (*Port, error) {
	c := &serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	port := &Port{
		p:      s,
		c:      make(chan byte, 128),
		buffer: make([]byte, 128),
	}
	return port, nil
}

func (p *Port) Read(buf []byte) (int, error) {
	return p.p.Read(buf)
}

// OpenChan open a channel to read serial buffer
func (p *Port) OpenChan() {
	go func() {
		for {
			n, err := p.p.Read(p.buffer)
			if err != nil {
				p.c <- EOF
				return
			}
			for i := 0; i < n; i++ {
				p.c <- p.buffer[i]
			}
		}
	}()
}

// CloseChan close serial buffer
func (p *Port) CloseChan() {
	p.c <- EOF
}

// GetChan return a channel object to read
func (p *Port) GetChan() chan byte {
	return p.c
}

// Write write byte array
func (p *Port) Write(buff []byte) (int, error) {
	return p.p.Write(buff)
}
