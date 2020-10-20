package middlewears

import (
	"github.com/andrewGIS/apiserver/db"
	"net/http"
	"strings"
)

//Check authorization for user. Token extracts from cookie
func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 401 if not user token founded
		accessToken, err := r.Cookie("userToken")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//Check that token is valid
		userEmail, err := r.Cookie("userEmail")
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if len(db.Users) == 0 {
			http.Error(w, "No user in base", http.StatusUnauthorized)
			return
		}

		token := db.Users[userEmail.Value].Token
		value := strings.Compare(
			token,
			accessToken.Value,
		)
		if value != 0 {
			http.Error(w, "Token not valid", http.StatusUnauthorized)
			return
		}

		// Continue process
		next.ServeHTTP(w, r)

	})
}

//Check authorization for current device. Token extracts from cookie
//if request for registering device will be send from same machine
// name of the device sets as last one
func DeviceAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 401 if device token not founded
		deviceToken, err := r.Cookie("deviceToken")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		deviceName, err := r.Cookie("deviceName")
		if err != nil {
			http.Error(w, "Unauthorized device", http.StatusUnprocessableEntity)
			return
		}

		userEmail, err := r.Cookie("userEmail")
		if err != nil {
			http.Error(w, "No active user", http.StatusUnprocessableEntity)
			return
		}

		if len(db.Users) == 0 {
			http.Error(w, "No users in base", http.StatusUnauthorized)
			return
		}

		//Check that token is valid
		token := db.Users[userEmail.Value].Devices[deviceName.Value].Token
		value := strings.Compare(
			token,
			deviceToken.Value,
		)
		if value != 0 {
			http.Error(w, "Unauthorized device", http.StatusUnauthorized)
			return
		}

		// Continue process
		next.ServeHTTP(w, r)

	})
}
