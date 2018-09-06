package megad328

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const socketsNum = 14

// Errors
var (
	ErrOutOfBounds = "megad328: device does not contain socket %d"
)

// Statuses
const (
	StatusPending = "pending" // connecting for a first time
	StatusOffline = "offline" // device is not responsing
	StatusOnline  = "online"  // device is responsing
)

// MegaD328 holds data for MegaD328 device
type MegaD328 struct {
	sync.Mutex // ???
	address    string
	client     http.Client
	Status     string        `json:"status"`
	UpdatedAt  time.Time     `json:"updated_at"`
	Sockets    []interface{} `json:"sockets"`
	Error      error         `json:"error,omitempty"`
}

// New creates a new MegaD328 device
func New(host, password string) *MegaD328 {
	u := url.URL{
		Scheme:   "http",
		Host:     host,
		Path:     password + "/",
		RawQuery: "cmd=all",
	}
	return &MegaD328{
		address: u.String(),
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		Status:  StatusPending,
		Sockets: make([]interface{}, socketsNum),
	}
}

// Ping checks if the device is alive
func (d *MegaD328) Ping() (bool, error) {
	resp, err := d.client.Get(d.address)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// Read reads value for a socket by it's ID
func (d *MegaD328) Read(id int) (interface{}, error) {
	if id < 0 && id > len(d.Sockets) {
		return nil, fmt.Errorf(ErrOutOfBounds, id)
	}

	return d.Sockets[id], nil
}

func (d *MegaD328) Write(id int, status interface{}) error {
	return nil // not implemented
}

// Fetch fetches data from the remote device
func (d *MegaD328) Fetch() error {
	if d.Error != nil {
		// Clean error on Fetch
		d.Error = nil
	}

	resp, err := d.client.Get(d.address)
	if err != nil {
		d.Status = StatusOffline
		d.Error = err
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		d.Status = StatusOffline
		d.Error = err
		return err
	}

	d.UpdatedAt = time.Now()

	values := strings.Split(string(body), ";")
	for i := range d.Sockets {
		d.Sockets[i] = parseValue(values[i])
	}

	// Update status if it is not "online"
	if d.Status != StatusOnline {
		d.Status = StatusOnline
	}

	return nil
}

func parseValue(value string) interface{} {
	num, err := strconv.ParseFloat(value, 64)
	// Return the value if it parsed successfully
	if err == nil {
		return num
	}

	// Replace all ON/x and OFF/x with true and false
	if strings.Contains(value, "OFF") {
		return false
	}
	if strings.Contains(value, "ON") {
		return true
	}

	// Just return the value if every parsers failed
	return value
}
