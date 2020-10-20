package devices

import (
	"encoding/json"
	"github.com/andrewGIS/apiserver/db"
	"github.com/andrewGIS/apiserver/utils"
	"io/ioutil"
	"net/http"
	"time"
)

// Function for adding new device
// POST method allowed. simple function.
// Expected db.DevicePayload structure in input
// JSON. Device with same name is not allowed
// Set token for auth for this device
func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":

		//Check input JSON
		err := utils.IsJsonValid(w, r)
		if err != nil {
			http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
			return
		}

		// Parse data to JSON dv.DevicePayload Type
		devicePayload := db.DevicePayload{}
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(b, &devicePayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Check that user in base
		user, found := db.Users[devicePayload.UserEmail]
		if !found {
			http.Error(w, "User with this email is not exists", http.StatusBadRequest)
			return
		}

		//Check that name of device is new
		_, found = user.Devices[devicePayload.Name]
		if found {
			http.Error(w, "Device with this name already exists", http.StatusBadRequest)
			return
		}

		//Create auth token for device and set cookies
		deviceToken := utils.String(50)
		cookie := http.Cookie{
			Name:     "deviceToken",
			Value:    deviceToken,
			Expires:  time.Time{},
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
			SameSite: 0,
		}
		http.SetCookie(w, &cookie)

		cookie = http.Cookie{
			Name:     "deviceName",
			Value:    devicePayload.Name,
			Expires:  time.Time{},
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
			SameSite: 0,
		}
		http.SetCookie(w, &cookie)

		cookie = http.Cookie{
			Name:     "userEmail",
			Value:    devicePayload.UserEmail,
			Expires:  time.Time{},
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
			SameSite: 0,
		}
		http.SetCookie(w, &cookie)

		device := &db.Device{}
		device.Token = deviceToken
		device.Data = []db.DataMeasure{}

		db.Users[devicePayload.UserEmail].Devices[devicePayload.Name] = device
		_, err = w.Write([]byte(deviceToken))
		if err != nil {
			http.Error(w, "Error writing token", http.StatusInternalServerError)
			return
		}

	}
}

// Function for adding Data Measure to specified device
// POST method allowed.
// Expected list of object in input
// [
//		{
//  		"Time": "2019-07-24T09:35:06.654Z",
//   		"T": 35.0
// 		},
//		{
//  		"Time": "2019-07-24T09:35:06.654Z",
//   		"T": 36.0
// 		}
//]
// JSON. Device with same name is not allowed
func DataWriter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		userEmail, err := r.Cookie("userEmail")
		if err != nil {
			http.Error(w, "No registered user", http.StatusUnprocessableEntity)
			return
		}

		deviceName, err := r.Cookie("deviceName")
		if err != nil {
			http.Error(w, "Device is not found", http.StatusUnprocessableEntity)
			return
		}

		device := db.Users[userEmail.Value].Devices[deviceName.Value]

		var dataPayload []db.DataMeasure

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(b, &dataPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		device.AddData(dataPayload)
		w.Write([]byte("Data loaded!"))
	}
}ggo

// Function for get statistic by device
// GET method allowed.
// Parameter "deviceName" for selecting user's device
// Expected db.DevicePayload structure in input
// JSON. Device with same name is not allowed
// Set token for auth for this device
func DeviceDataMeasureInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		userEmail, err := r.Cookie("userEmail")
		if err != nil {
			http.Error(w, "No user found", http.StatusUnprocessableEntity)
			return
		}

		deviceNames, ok := r.URL.Query()["deviceName"]

		if !ok || len(deviceNames) > 1 {
			http.Error(w, "No device specified or request more than one device", http.StatusBadRequest)
			return
		}

		_, found := db.Users[userEmail.Value].Devices[deviceNames[0]]
		if !found {
			http.Error(w, "No such device", http.StatusBadRequest)
			return
		}

		var deviceStat = map[string]map[string]db.DataMeasure{}
		d := db.Users[userEmail.Value].Devices[deviceNames[0]]

		// Error if data is empty
		if len(d.Data) == 0 {
			http.Error(w, "No data are loaded to device", http.StatusBadRequest)
			return
		}

		var stats = map[string]db.DataMeasure{}
		stats["Max"] = utils.MaxValueInArray(d.Data)
		stats["Min"] = utils.MinValueInArray(d.Data)
		deviceStat[deviceNames[0]] = stats

		deviceStatData, err := json.Marshal(&deviceStat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(deviceStatData)
		if err != nil {
			http.Error(w, "Error writing device statistic", http.StatusUnprocessableEntity)
		}

	}
}
