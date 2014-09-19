// A simple library for interacting with the Rainforest Automation RAVEn
package goraven

import (
	"bufio"
	"encoding/xml"
	"errors"
	"github.com/schleibinger/sio"
	"syscall"
	"bytes"
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
	err = initReader(r)
	if err != nil {
		return nil, err
	}
	return &Raven{p, r}, nil
}

// Close closes the RAVEn's port safely
func (r *Raven) Close() error {
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

// initReader reads and discards the first message
func initReader(r *bufio.Reader) error {
	for {
		line, err := r.ReadBytes(10)
		if err != nil {
			return err
		}
		if isEndElement(line) {
			return nil
		}
	}
	panic("unreachable")
}

func nextStart(dec *xml.Decoder) (xml.StartElement, error) {
	// Find the next starting element
	for {
		t, err := dec.Token()
		if err != nil {
			return xml.StartElement{}, err
		}
		switch t := t.(type) {
		case xml.StartElement:
			return t, nil
		}
	}
}

// Receive grabs the next "Notify" message in the stream
func (r *Raven) Receive() (notify interface{}, err error) {
	dec := xml.NewDecoder(r.p)

	se, err := nextStart(dec)
	if err != nil {
		return nil, err
	}

	switch se.Name.Local {
	case "ConnectionStatus":
		notify = &ConnectionStatus{}
	case "DeviceInfo":
		notify = &DeviceInfo{}
	case "ScheduleInfo":
		notify = &ScheduleInfo{}
	case "MeterList":
		notify = &MeterList{}
	case "NetworkInfo":
		notify = &NetworkInfo{}
	case "MeterInfo":
		notify = &MeterInfo{}
	default:
		return nil, errors.New("Unrecognized Notify Message")
	}

	err = dec.DecodeElement(notify, &se)
	return notify, err
}

// Begins with '  <'
func isMidElement(line []byte) bool {
	return bytes.HasPrefix(line, []byte{32, 32, 60})
}

// Begins with '</'
func isEndElement(line []byte) bool {
	return bytes.HasPrefix(line, []byte{60, 47})
}
