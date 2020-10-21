package db

// TODO Make this structure with sync.Mutex
// Type Users struct {
// mu *sync.Mutex
// users map[]...
// }
//И для этой структуры методы new, get, store

//'In-memory' storage
var Users = map[string]*User{}

type User struct {
	Token   string
	Devices map[string]*Device
}

type Device struct {
	Token string
	Data  []DataMeasure
}

//Adding new measures to Device
func (data *Device) AddData(appendData []DataMeasure) {
	data.Data = append(data.Data, appendData...)
}

// Struct for data load when User adds
// a device to account
type DevicePayload struct {
	UserEmail string `json:"userEmail"`
	Name      string `json:"deviceName"`
}

// Struct for User account load
type UserPayload struct {
	UserEmail string `json:"userEmail"`
}

//Structure for storing measure from device
//Sample
//{
//  "Time": "2019-07-24T09:35:06.654Z", // time measure
//   "T": 35.0 							// temperature
// }
type DataMeasure struct {
	Time string  `json:"Time"`
	T    float32 `json:"T"`
}
