package model

// Sensor common sensor data struct
type Sensor struct {
	Location   string `json:"loc"`
	Timestamp  int64  `json:"timestamp"`
	Raw        int16  `json:"raw"`
	Value      int32  `json:"value"`
	Unit       string `json:"unit"`
	SensorType string `json:"sensortype"`
	ID         int64  `json:"id"`
}
