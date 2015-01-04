package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/encoder"
)

var configFile string

const defaultConfig = "./config/db.cfg"
const arpTableFile = "/proc/net/arp"

func init() {
	flag.StringVar(&configFile, "config", defaultConfig, "Database config file")
}

func JsonContent(c martini.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func SetupMartini() *martini.Martini {
	m := martini.New()
	m.Use(martini.Recovery())
	m.MapTo(jsonEncoder{}, (*encoder.Encoder)(nil))

	r := martini.NewRouter()
	r.Get("/mac", JsonContent, GetMac)
	r.Get("/users/:id", JsonContent, GetUser)
	r.Get("/users", JsonContent, GetUsers)
	r.Get("/users/:id/devices", JsonContent, GetDevicesByUser)
	r.Post("/users", JsonContent, binding.Json(User{}), AddUser)
	r.Put("/users/:id", JsonContent, binding.Json(User{}), UpdateUser)
	r.Post("/users/:id/devices", JsonContent, binding.Json(Device{}), AddDevice)
	m.Action(r.Handle)
	return m
}

func main() {
	flag.Parse()
	conf := ReadConfig(configFile)

	m := SetupMartini()
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))
	var dataStore DataStore = CreateMysqlDataStoreFromConfig(conf)
	m.MapTo(dataStore, (*DataStore)(nil))
	m.Map(arpTableFile)

	log.Fatal(http.ListenAndServe(":8080", m))
}
