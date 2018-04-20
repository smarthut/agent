// Package laurent112 provides Go API for MP716 Laurent 112 device
// more information: http://kernelchip.ru/Laurent-112.php
package laurent112

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	socketsNum  = 12
	cmdTemplate = "REL,%d,%d"
)

// Errors
var (
	ErrGetStatus    = errors.New("laurent112: unable to get status")
	ErrOutOfBounds  = "laurent112: socket %d is out of bounds"
	ErrBadResponse  = "laurent112: request REL,%d,%d was not successful, response: %s"
	ErrBadNewStatus = errors.New("laurent112: status should be 0 (OFF) or 1 (ON)")
)

// Laurent112 contains data for Megad328 device
type Laurent112 struct {
	url       url.URL
	client    http.Client     // HTTP connection
	UpdatedAt time.Time       `json:"updated_at"`
	Sockets   [socketsNum]int `json:"sockets"`
}

type socket int

// New created a new Laurent112 device
func New(host string) *Laurent112 {
	return &Laurent112{
		url: url.URL{
			Scheme: "http",
			Host:   host,
		},
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Response of Laurent112
type response struct {
	XMLName xml.Name `xml:"response"`
	SysTime int      `xml:"systime0"`
	Relay   string   `xml:"rele_table0"`
}

// UpdateSockets fetch sockets from Laurent112 and update time
func (d *Laurent112) UpdateSockets() error {
	u := d.url
	u.Path = "status.xml"

	resp, err := d.client.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r response
	xml.NewDecoder(resp.Body).Decode(&r)

	vals := strings.Split(r.Relay, "")

	d.UpdatedAt = time.Now()
	for i := range d.Sockets {
		v, err := strconv.Atoi(vals[i])
		if err != nil {
			return err
		}
		d.Sockets[i] = v
	}

	return nil
}

// Get returns socket#id value
func (d *Laurent112) Get(id int) (interface{}, error) {
	if id < 0 || id > len(d.Sockets) {
		return nil, fmt.Errorf(ErrOutOfBounds, id)
	}

	return d.Sockets[id], nil
}

// Set status to socket
func (d *Laurent112) Set(id, status int) error {
	if id < 0 || id > len(d.Sockets) {
		return fmt.Errorf(ErrOutOfBounds, id)
	}

	if status != 0 && status != 1 {
		return fmt.Errorf(ErrOutOfBounds, id)
	}

	if int(d.Sockets[id]) == status {
		d.UpdateSockets() // force update, skip if status already in desired state
		return nil
	}

	u := d.url
	u.Path = "cmd.cgi"
	v := url.Values{}
	cmd := fmt.Sprintf(cmdTemplate, id+1, status) // add 1 to id, Laurent112 starting count from 1
	v.Add("cmd", cmd)
	u.RawQuery = v.Encode()

	resp, err := d.client.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) != "DONE" {
		return fmt.Errorf(ErrBadResponse, id+1, status, string(body))
	}

	d.UpdateSockets() // force update status, check if new status is wanted
	if int(d.Sockets[id]) != status {
		return fmt.Errorf("laurent112: changing REL %d to %d was not successful, the value is still %d", id+1, status, int(d.Sockets[id]))
	}

	return nil
}
