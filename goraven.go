// A simple library for interacting with the Rainforest Automation RAVEn
package goraven

import (
	"bufio"
	"encoding/xml"
	"github.com/schleibinger/sio"
	"syscall"
)

// Types of events
const (
	TIME      = "time"
	PRICE     = "price"
	DEMAND    = "demand"
	SUMMATION = "summation"
	MESSAGE   = "message"
)

type Raven struct {
	p      *sio.Port
	reader *bufio.Reader
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
	r := bufio.NewReader(p)
	return &Raven{p, r}, nil
}

// Disconnect closes the RAVEn's port safely
func (r *Raven) Disconnect() error {
	return r.p.Close()
}

// simpleCommand sends a simple command
func (r *Raven) simpleCommand(command string) error {
	v := &smplCommand{Name: command}
	return r.sendCommand(v)
}

// sendCommand sends a generic command
func (r *Raven) sendCommand(v interface{}) error {
	enc := xml.NewEncoder(r.p)
	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}

func (r *Raven) Read() (line []byte, err error) {
	return r.reader.ReadBytes(13) // Read until CR (line break)
}
