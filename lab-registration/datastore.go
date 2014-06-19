package main

type DataStore interface {
	GetAllUsers() ([]SimpleUser, error)
	GetUser(id int) (User, error)
	AddUser(u User) (int64, error)
	UpdateUser(userId int, user User) (error)
	GetDevicesByUserId(userId int) ([]Device, error)
	AddDevice(userId int, device Device) (int64, error)
	GetDeviceByMac(mac string) (*Device, error)
}
