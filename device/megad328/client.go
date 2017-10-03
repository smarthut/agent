package megad328

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const socketsNum = 14

// Errors
var (
	ErrGetStatus    = errors.New("laurent112: unable to get status")
	ErrOutOfBounds  = errors.New("laurent112: socket out of bounds")
	ErrBadNewStatus = errors.New("laurent112: status should be 0 (OFF) or 1 (ON)")
)

// MegaD328 contains data for MegaD328 device
type MegaD328 struct {
	url       url.URL
	client    http.Client
	UpdatedAt time.Time          `json:"updated_at"`
	Sockets   [socketsNum]socket `json:"sockets"`
}

type socket interface{}

// New created a new MegaD328 device
func New(host, password string) *MegaD328 {
	return &MegaD328{
		url: url.URL{
			Scheme:   "http",
			Host:     host,
			Path:     password + "/",
			RawQuery: "cmd=all",
		},
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// UpdateSockets fetch sockets from MegaD328 and update time
func (d *MegaD328) UpdateSockets() error {
	resp, err := d.client.Get(d.url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	vals := strings.Split(string(body), ";")

	d.UpdatedAt = time.Now()
	for i := range d.Sockets {
		// we don't care about the error, not all values are parsable
		v, _ := strconv.ParseFloat(vals[i], 64)
		d.Sockets[i] = v
	}

	return nil
}

// Get returns socket#id value
func (d *MegaD328) Get(id int) (interface{}, error) {
	if id < 0 || id > len(d.Sockets) {
		return nil, ErrOutOfBounds
	}

	return d.Sockets[id], nil
}

// Set ...
func (d *MegaD328) Set(id, value int) error {
	return nil
}

// Len ...
func (d *MegaD328) Len() int {
	return len(d.Sockets)
}
