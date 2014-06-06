package main

type DataStore interface {
	GetAllUsers() ([]SimpleUser, error)
	GetUser(id int) (User, error)
	AddUser(u User) (int64, error)
	GetDevicesByUserId(userId int) ([]Device, error)
}
