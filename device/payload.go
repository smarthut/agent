package device

// Payload holds payload for device
type Payload struct {
	ID     int         `json:"id"`
	Status interface{} `json:"status"`
}
