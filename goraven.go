// A simple library for interacting with the Rainforest Automation RAVEn
package goraven

import (
	"encoding/xml"
	"github.com/schleibinger/sio"
	"syscall"
)

const (
	TIME      = "time"
	PRICE     = "price"
	DEMAND    = "demand"
	SUMMATION = "summation"
	MESSAGE   = "message"
)

type Raven struct {
	p *sio.Port
}

// The structure for a simple command with a single argument (Name)
type smplCommand struct {
	XMLName xml.Name `xml:"Command"`
	Name    string   `xml:"Name"`
}

// Connect opens a connection to a RAVEn, given the port name (/dev/ttyUSB0)
func Connect(dev string) (*Raven, error) {
	p, err := sio.Open(dev, syscall.B115200)
	if err != nil {
		return nil, err
	}
	return &Raven{p}, nil
}

// Disconnect closes the RAVEn's port safely
func (r *Raven) Disconnect() error {
	return r.p.Close()
}

// Send a simple command
func (r *Raven) simpleCommand(command string) error {
	sc := &smplCommand{Name: command}
	enc := xml.NewEncoder(r.p)
	if err := enc.Encode(sc); err != nil {
		return err
	}
	return nil
}
