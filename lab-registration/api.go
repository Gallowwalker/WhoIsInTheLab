package main

import (
	"net/http"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/martini-contrib/binding"
	"github.com/go-martini/martini"
)


func GetMac(res http.ResponseWriter, req *http.Request, enc encoder, dataStore DataStore, arpFile string) (int, []byte) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	mac, err := GetMacAddress(arpFile, ip)
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrMacNotFound, err.Error())))
	}
	device, err := dataStore.GetDeviceByMac(mac)
	responseMap := map[string]interface{} {"mac":mac, "registered": device != nil}
	if device != nil {
		responseMap["device"] = *device
	}
	return http.StatusOK, must(enc.Encode(responseMap))
}

func GetUsers(dataStore DataStore, enc encoder) (int, []byte) {
	users, err := dataStore.GetAllUsers()
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Can't get users try again later")))
	}
	return http.StatusOK, must(enc.Encode(users))
}

func GetUser(dataStore DataStore, enc encoder, params martini.Params) (int, []byte) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Invalid id")))
	}

	user, dbErr := dataStore.GetUser(id)
	if dbErr != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrUserNotFound, dbErr.Error())))
	}
	return http.StatusOK, must(enc.Encode(user))
}

func AddUser(user User, enc encoder, dataStore DataStore, bindErrors binding.Errors) (int, []byte) {
	if len(bindErrors) > 0 {
		errors, _ := json.Marshal(bindErrors)
		return http.StatusBadRequest, errors
	}
	id, err := dataStore.AddUser(user)
	if err != nil {
		return http.StatusBadRequest, must(enc.Encode(NewError(ErrInternal, "Cannot create user")))
	}

	response := AddResponse{Success: true, Id: id}
	return http.StatusOK, must(enc.Encode(response))
}
func UpdateUser(user User, enc encoder, dataStore DataStore, params martini.Params, bindErrors binding.Errors) (int, []byte) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Invalid id")))
	}
	if len(bindErrors) > 0 {
		errors, _ := json.Marshal(bindErrors)
		return http.StatusBadRequest, errors
	}
	err = dataStore.UpdateUser(id, user)
	if err != nil {
		return http.StatusBadRequest, must(enc.Encode(NewError(ErrInternal, "Cannot update user")))
	}

	response := AddResponse{Success: true, Id: int64(id)}
	return http.StatusOK, must(enc.Encode(response))
}

func GetDevicesByUser(dataStore DataStore, enc encoder, params martini.Params) (int, []byte) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Invalid id")))
	}
	devices, err := dataStore.GetDevicesByUserId(id)
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Can't get user's devices try again later")))
	}
	return http.StatusOK, must(enc.Encode(devices))
}
func AddDevice(device Device, enc encoder, dataStore DataStore, bindErrors binding.Errors, params martini.Params) (int, []byte) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusNotFound, must(enc.Encode(NewError(ErrInternal, "Invalid user id")))
	}
	if len(bindErrors) > 0 {
		errors, _ := json.Marshal(bindErrors)
		return http.StatusBadRequest, errors
	}
	deviceId, err := dataStore.AddDevice(id, device)
	if err != nil {
		return http.StatusBadRequest, must(enc.Encode(NewError(ErrInternal, "Cannot add device")))
	}

	return http.StatusOK, must(enc.Encode(map[string]interface{} {"success": true, "id": deviceId}))
}
