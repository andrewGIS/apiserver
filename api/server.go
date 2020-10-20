package api

import (
	"github.com/andrewGIS/apiserver/api/controllers/devices"
	"github.com/andrewGIS/apiserver/api/controllers/users"
	"github.com/andrewGIS/apiserver/api/middlewears"
	"log"
	"net/http"
)

//Base server struct on port 8080
type Server struct {
}

//Create and run server
func (*Server) Start() {
	//http.HandleFunc("/testCookie", testReadCookie)
	//http.HandleFunc("/testParams", testGetParams)

	mux := http.NewServeMux()

	//Handlers for adding new Users and get devices
	// for specific user
	userHandler := http.HandlerFunc(users.UserHandler)
	userDevicesHandler := http.HandlerFunc(users.UserDevices)

	mux.Handle("/users", userHandler)
	mux.Handle("/users/devices", middlewears.DBCheck(userDevicesHandler))

	// Handlers for adding new device, write data to specific device
	// and get statistic by device
	deviceHandler := http.HandlerFunc(devices.DeviceHandler)
	deviceDataHandler := http.HandlerFunc(devices.DataWriter)
	deviceDataInfoHandler := http.HandlerFunc(devices.DeviceDataMeasureInfo)

	mux.Handle("/devices", middlewears.DBCheck(deviceHandler))
	mux.Handle("/devices/data",
		middlewears.DBCheck(middlewears.DeviceAuth(deviceDataHandler)))
	mux.Handle("/devices/data/info",
		middlewears.DBCheck(middlewears.UserAuth(deviceDataInfoHandler)))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
