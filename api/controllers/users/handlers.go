package users

import (
	"encoding/json"
	"github.com/andrewGIS/apiserver/db"
	"github.com/andrewGIS/apiserver/utils"
	"io/ioutil"
	"net/http"
	"time"
)

//Function for adding new user and list existing ones
// GET and POST method allowed. Email of user is check by
// simple function. Expected db.UserPayload structure in input
// JSON
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":

		usersJSON, err := json.Marshal(db.Users)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(usersJSON)

	case "POST":

		err := utils.IsJsonValid(w, r)
		if err != nil {
			return
		}

		userPayload := db.UserPayload{}
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(b, &userPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//validate email
		validEmail := utils.IsEmailValid(userPayload.UserEmail)
		if !validEmail {
			msg := "Email address is not valid"
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		accessToken := utils.String(50)

		cookie := http.Cookie{
			Name:     "userToken",
			Value:    accessToken,
			Expires:  time.Time{},
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
			SameSite: 0,
		}
		http.SetCookie(w, &cookie)

		cookie = http.Cookie{
			Name:     "userEmail",
			Value:    userPayload.UserEmail,
			Expires:  time.Time{},
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
			SameSite: 0,
		}
		http.SetCookie(w, &cookie)

		user := db.User{}
		user.Token = accessToken
		user.Devices = make(map[string]*db.Device)

		db.Users[userPayload.UserEmail] = &user

		w.Write([]byte(accessToken))
	}
}

// Function returning list user's device
// GET method allowed. Authorization is checked by
// token
func UserDevices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		//Get current user
		userEmail, err := r.Cookie("userEmail")
		if err != nil {
			http.Error(w, "No user found", http.StatusForbidden)
			return
		}

		if len(db.Users) == 0 {
			http.Error(w, "No user in db", http.StatusBadRequest)
			return
		}

		// Check that at least one devices is registered
		if len(db.Users[userEmail.Value].Devices) == 0 {
			http.Error(w, "No device register for user", http.StatusBadRequest)
			return
		}

		//Get devices
		var deviceMap = map[string]string{}
		for deviceName, device := range db.Users[userEmail.Value].Devices {
			lastMeasure := utils.GetLastTimeMeasure(device.Data)
			deviceMap[deviceName] = lastMeasure.Time
			//fmt.Println(lastMeasure.Time)
		}

		deviceInfo, err := json.Marshal(deviceMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(deviceInfo)

	}
}
