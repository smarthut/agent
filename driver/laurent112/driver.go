package laurent112

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/smarthut/agent/helper/boolconv"
)

const (
	socketsNum  = 12
	cmdTemplate = "REL,%d,%d"
)

// Errors
var (
	ErrOutOfBounds    = "laurent112: device does not contain socket %d"
	ErrBadResponse    = "laurent112: request REL,%d,%d was not successful, response: %s"
	ErrBadNewStatus   = errors.New("laurent112: status should be 0 (OFF) or 1 (ON)")
	ErrUnableToChange = "laurent112: changing REL %d to %d was not successful, the value is still %d"
)

// Statuses
const (
	StatusPending = "pending" // connecting for a first time
	StatusOffline = "offline" // device is not responsing
	StatusOnline  = "online"  // device is responsing
)

// Laurent112 holds data for MegaD328 device
type Laurent112 struct {
	sync.Mutex // ???
	address    url.URL
	client     http.Client
	Status     string        `json:"status"`
	UpdatedAt  time.Time     `json:"updated_at"`
	Sockets    []interface{} `json:"sockets"` // TODO: replace with bool
	Error      error         `json:"error,omitempty"`
}

// New creates a new Laurent112 device
func New(host string) *Laurent112 {
	return &Laurent112{
		address: url.URL{
			Scheme: "http",
			Host:   host,
		},
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		Status:  StatusPending,
		Sockets: make([]interface{}, socketsNum),
	}
}

// Ping checks if the device is alive
func (d *Laurent112) Ping() (bool, error) {
	resp, err := d.client.Get(d.address.String())
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// Read reads value for a socket by it's ID
func (d *Laurent112) Read(id int) (interface{}, error) {
	if id < 0 && id > len(d.Sockets) {
		return nil, fmt.Errorf(ErrOutOfBounds, id)
	}

	return d.Sockets[id], nil
}

func (d *Laurent112) Write(id int, status interface{}) error {
	newStatus, err := boolconv.GetBool(status)
	if err != nil {
		return err
	}

	if id < 0 || id > len(d.Sockets) {
		return fmt.Errorf(ErrOutOfBounds, id)
	}

	if (d.Sockets[id]).(bool) == newStatus {
		d.Fetch() // force update, skip if status already in desired state
		return nil
	}

	u := d.address
	u.Path = "cmd.cgi"
	v := url.Values{}
	cmd := fmt.Sprintf(cmdTemplate, id+1, boolconv.Btoi(newStatus)) // add 1 to id, Laurent112 starting count from 1
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
		return fmt.Errorf(ErrBadResponse, id+1, newStatus, string(body))
	}

	d.Fetch() // force update status, check if new status is same sa requested
	if (d.Sockets[id]).(bool) != newStatus {
		return fmt.Errorf(ErrUnableToChange, id+1, newStatus, d.Sockets[id])
	}

	return nil
}

// Response of Laurent112
type response struct {
	XMLName xml.Name `xml:"response"`
	SysTime int      `xml:"systime0"`
	Relay   string   `xml:"rele_table0"`
}

// Fetch fetches data from the remote device
func (d *Laurent112) Fetch() error {
	if d.Error != nil {
		// Clean error on Fetch
		d.Error = nil
	}

	u := d.address
	u.Path = "status.xml"

	resp, err := d.client.Get(u.String())
	if err != nil {
		d.Status = StatusOffline
		d.Error = err
		return err
	}
	defer resp.Body.Close()

	d.UpdatedAt = time.Now()

	var r response
	if err := xml.NewDecoder(resp.Body).Decode(&r); err != nil {
		d.Status = StatusOffline
		d.Error = err
		return err
	}

	for i := range d.Sockets {
		d.Sockets[i] = parseValue(string(r.Relay[i]))
	}

	// Update status if it is not "online"
	if d.Status != StatusOnline {
		d.Status = StatusOnline
	}

	return nil
}

func parseValue(value string) interface{} {
	num, err := strconv.Atoi(value)
	// Return the value if it parsed successfully
	if err == nil {
		return num
	}

	return boolconv.Itob(num)
}
