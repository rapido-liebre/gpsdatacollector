package internal

type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type GpsData struct {
	Id       string `json:"id,omitempty"`
	DeviceId string `json:"device_id"`
	Point    Coordinates
}
