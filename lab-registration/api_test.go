package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"
	"flag"

	"github.com/go-martini/martini"
	. "github.com/smartystreets/goconvey/convey"
)

var testDbConfigFile string

const defaultTestConfig = "./test-data/db.cfg"

func init() {
	flag.StringVar(&testDbConfigFile, "testconfig", defaultTestConfig, "Test database config file")
}
func MockApi() *martini.Martini {
	m := SetupMartini()
	m.Map("test-data/arp-data")
	return m
}

func DbApi() *martini.Martini {
	m := MockApi()
	conf := ReadConfig(testDbConfigFile)
	var dataStore DataStore = CreateTestMysqlDataStoreFromConfig(conf)

	m.MapTo(dataStore, (*DataStore)(nil))
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
func TestAddUser(t *testing.T) {
	Convey("Given that add user endpoint is called with valid data", t, func() {
		userJson := ReadFile("./test-data/test_user.json")
		r, _ := http.NewRequest("POST", "/users", strings.NewReader(userJson))
		w := httptest.NewRecorder()
		DbApi().ServeHTTP(w, r)

		var result map[string]interface{}
		Convey("it should respond with success and id of newly created user", func() {
			json.Unmarshal(w.Body.Bytes(), &result)
			So(w.Code, ShouldEqual, 200)
			So(result["success"], ShouldEqual, true)
		})
	})

}
