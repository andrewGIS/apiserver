package middlewears

import (
	"github.com/andrewGIS/apiserver/db"
	"net/http"
)

//Middleware for check that db is not empty
// It is necessary for requests where data
// is loaded from db.Users
func DBCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(db.Users) == 0 {
			http.Error(w, "No users in base", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)

	})
}
