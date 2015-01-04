package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"
	"flag"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	. "github.com/smartystreets/goconvey/convey"
)

var testDbConfigFile string

const defaultTestConfig = "./test-data/db.cfg"

func init() {
	flag.StringVar(&testDbConfigFile, "testconfig", defaultTestConfig, "Test database config file")
}
func MockApi() *martini.Martini {
	m := SetupMartini()
	m.Map("./test-data/arp-data")
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
	Convey("Given that user hits /mac endpoint with an ip address that is present in the arp table and mac is not reigstered", t, func() {
		r, _ := http.NewRequest("GET", "/mac", nil)
		r.RemoteAddr = "192.168.50.1"
		w := httptest.NewRecorder()
		DbApi().ServeHTTP(w, r)

		var macResult map[string]interface{}
		Convey("it should return the corresponding mac address", func() {
			json.Unmarshal(w.Body.Bytes(), &macResult)
			So(w.Code, ShouldEqual, 200)
			So(macResult["mac"], ShouldEqual, "00:16:0a:13:96:7e")
			So(macResult["registered"], ShouldEqual, false)
			So(macResult["device"], ShouldEqual, nil)
		})
	})

	Convey("Given that user hits /mac endpoint with an ip address that is present in the arp table and MAC is already registered", t, func() {
		dbApi := DbApi()
		userId := insertTestUser(dbApi).Id
		insert(ReadFile("./test-data/test_device.json"), fmt.Sprintf("/users/%d/devices", userId), dbApi)

		r, _ := http.NewRequest("GET", "/mac", nil)
		r.RemoteAddr = "192.168.50.1"
		w := httptest.NewRecorder()
		dbApi.ServeHTTP(w, r)

		var macResult map[string]interface{}
		Convey("it should return the corresponding mac address and a flag thath indicates that device is registered", func() {
			json.Unmarshal(w.Body.Bytes(), &macResult)
			So(w.Code, ShouldEqual, 200)
			So(macResult["mac"], ShouldEqual, "00:16:0a:13:96:7e")
			So(macResult["registered"], ShouldEqual, true)
			So(macResult["device"], ShouldNotEqual, nil)
		})
	})

	Convey("Given that user hits /mac endpoint with an empty ip address", t, func() {
		r, _ := http.NewRequest("GET", "/mac", nil)
		w := httptest.NewRecorder()
		DbApi().ServeHTTP(w, r)

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

		result := AddResponse{}
		Convey("it should respond with success and id of newly created user", func() {
			json.Unmarshal(w.Body.Bytes(), &result)
			So(w.Code, ShouldEqual, 200)
			So(result.Success, ShouldEqual, true)
			So(result.Id, ShouldBeGreaterThan, 0)
		})
	})
	Convey("Given that add user endpoint is called with a user with empty first name", t, func() {
		userJson := ReadFile("./test-data/invalid_user.json")
		r, _ := http.NewRequest("POST", "/users", strings.NewReader(userJson))
		w := httptest.NewRecorder()
		DbApi().ServeHTTP(w, r)

		var errors binding.Errors
		Convey("it shold respond with error message saying that firstname must not be empty", func() {
			json.Unmarshal(w.Body.Bytes(), &errors)
			So(w.Code, ShouldEqual, 400)
			So(errors.Has("Incorrect data"), ShouldBeTrue)
			So(errors[0].FieldNames[0], ShouldEqual, "firstname")
		})
	})
	Convey("Given that update user endpoint is called with valid data", t, func() {
		dbApi := DbApi()
		userResult := insertTestUser(dbApi)
		endpoint := fmt.Sprintf("/users/%d", userResult.Id)

		r, _ := http.NewRequest("PUT", endpoint, strings.NewReader(ReadFile("./test-data/test_user_modified.json")))
		w := httptest.NewRecorder()
		dbApi.ServeHTTP(w, r)

		result := AddResponse{}
		Convey("it should successfully update users data", func() {
			json.Unmarshal(w.Body.Bytes(), &result)
			So(w.Code, ShouldEqual, 200)
			So(result.Success, ShouldEqual, true)
			So(result.Id, ShouldBeGreaterThan, 0)
			user := getUser(dbApi, result.Id)
			So(user.Id, ShouldBeGreaterThan, 0)
			So(user.FirstName, ShouldEqual, "JohnDoe")
			So(user.Facebook, ShouldEqual, "john.the.doe")
			So(user.Phone, ShouldEqual, "911")
			So(user.Twitter, ShouldEqual, "johndoe.the.great")
			So(user.Email, ShouldEqual, "johndoe@example.com")
		})


	})

}
func TestAddDevice(t *testing.T) {
	Convey("Given that  add device for a given user endpoint is called with valid data", t, func() {
		deviceJson := ReadFile("./test-data/test_device.json")
		dbApi := DbApi()
		userResult := insertTestUser(dbApi)
		endpoint := fmt.Sprintf("/users/%d/devices", userResult.Id)

		r, _ := http.NewRequest("POST", endpoint, strings.NewReader(deviceJson))
		w := httptest.NewRecorder()
		dbApi.ServeHTTP(w, r)

		result := AddResponse{}
		Convey("it should respond with success and id of newly created device", func() {
			json.Unmarshal(w.Body.Bytes(), &result)
			So(w.Code, ShouldEqual, 200)
			So(result.Success, ShouldEqual, true)
			So(result.Id, ShouldBeGreaterThan, 0)
		})
	})
	Convey("Given that add device endpoint is called with invalid user id", t, func() {
		deviceJson := ReadFile("./test-data/test_device.json")
		dbApi := DbApi()
		endpoint := fmt.Sprintf("/users/%d/devices", 23)

		r, _ := http.NewRequest("POST", endpoint, strings.NewReader(deviceJson))
		w := httptest.NewRecorder()
		dbApi.ServeHTTP(w, r)

		result := Error{}
		Convey("it should respond with error saying that user doesen't exists", func() {
			json.Unmarshal(w.Body.Bytes(), &result)
			So(w.Code, ShouldEqual, 400)
			So(result.Message, ShouldEqual, "Cannot add device")
		})
	})
}

func insertTestUser(api *martini.Martini) (AddResponse) {
		return insert(ReadFile("./test-data/test_user.json"), "/users", api)
}

func insert(data, endpoint string, api *martini.Martini) AddResponse {
		r, _ := http.NewRequest("POST", endpoint, strings.NewReader(data))
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		result := AddResponse{}
		json.Unmarshal(w.Body.Bytes(), &result)
		return result
}

func getUser(api *martini.Martini, userId int64) (User) {
		r, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", userId), nil)
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		result := User{}
		json.Unmarshal(w.Body.Bytes(), &result)

		return result
}
