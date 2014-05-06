package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/go-martini/martini"
	. "github.com/smartystreets/goconvey/convey"
)

func MockApi() *martini.Martini {
	m := SetupMartini()
	m.Map("test-data/arp-data")
	return m
}

func TestMacAPI(t *testing.T) {
	Convey("Given that user hits /mac endpoint with an ip address that is present in the arp table", t, func() {
		r, _ := http.NewRequest("GET", "/mac", nil)
		r.RemoteAddr = "192.168.50.1"
		w := httptest.NewRecorder()
		MockApi().ServeHTTP(w, r)

		var macResult map[string]interface{}
		Convey("it should return the corresponding mac address", func() {
			json.Unmarshal(w.Body.Bytes(), &macResult)
			So(w.Code, ShouldEqual, 200)
			So(macResult["mac"], ShouldEqual, "00:16:0a:13:96:7e")
		})
	})

	Convey("Given that user hits /mac endpoint with an ip that is not present in arp table", t, func() {
		r, _ := http.NewRequest("GET", "/mac", nil)
		r.RemoteAddr = "8.8.8.8"
		w := httptest.NewRecorder()
		MockApi().ServeHTTP(w, r)

		Convey("it should respond with error code 404", func() {
			So(w.Code, ShouldEqual, 404)
		})
	})

	Convey("Given that user hits /mac endpoint with an empty ip address", t, func() {
		r, _ := http.NewRequest("GET", "/mac", nil)
		w := httptest.NewRecorder()
		MockApi().ServeHTTP(w, r)

		Convey("it should respond with error code 404", func() {
			So(w.Code, ShouldEqual, 404)
		})
	})
}
