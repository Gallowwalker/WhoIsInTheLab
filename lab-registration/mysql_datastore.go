package main

import (
	 "log"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
)


const defaultPort = ":3306"
const testDatabaseName = "whoIsInTheLabTest"

type MySqlDatastore struct {
	db *sqlx.DB
}

func CreateMySqlDatastore(user, pass, host, dbName string) (DataStore) {
	if len(strings.Split(host, ":")) == 1 {
		host = host + defaultPort
	}

	connString := user + ":" + pass + "@tcp(" + host +")/"
	if len(strings.TrimSpace(dbName)) > 0 {
		connString +=  dbName
	}
	db, err := sqlx.Connect("mysql", connString)
	if (err != nil) {
		log.Fatal(err)
	}
	var dataStore DataStore = &MySqlDatastore{db}
	return dataStore
}

func CreateTestMysqlDataStoreFromConfig(config Config) (DataStore) {
	dataStore := CreateMySqlDatastore(config.Username, config.Password, config.Host, config.Database)

	config.Database = testDatabaseName

	mysql := dataStore.(*MySqlDatastore)
	mysql.db.Exec("create database if not exists " + testDatabaseName + ";")
	mysql.db.Exec("use " + testDatabaseName + ";")
	_, err := ExecuteSql("./test-data/schema.sql", config)
	checkError(err)
	return dataStore
}

func CreateMysqlDataStoreFromConfig(config Config) (DataStore) {
	return CreateMySqlDatastore(config.Username, config.Password, config.Host, config.Database)
}


func (d MySqlDatastore) GetAllUsers() ([]SimpleUser, error) {
	users := []SimpleUser{}

	dbError := d.db.Select(&users, "SELECT user_id, user_name1, user_name2 from who_users")
	if dbError != nil {
		return nil, dbError
	}
	return users, dbError
}

func (d MySqlDatastore) GetUser(id int) (User, error) {
	user := User{}
	err := d.db.Get(&user, `SELECT  user_name1, 
					user_name2, 
					user_id,
					user_twitter, 
					user_email, 
					user_facebook, 
					user_tel, 
					user_website,  
					user_google_plus, 
					user_fscheckin 
					FROM who_users WHERE user_id=?`, id)
	if err != nil {
		return user, fmt.Errorf("User with id %d not found", id)
	}

	return user, nil
}

func (d MySqlDatastore) AddUser(u User) (int64, error) {
	res, err := d.db.NamedExec(`INSERT INTO who_users (user_name1, user_name2, user_twitter, user_email, user_facebook, user_tel, user_website, user_fstoken, user_google_plus, user_fscheckin) VALUES (:user_name1, :user_name2, :user_twitter, :user_email, :user_facebook, :user_tel, :user_website, :user_fstoken, :user_google_plus, :user_fscheckin)`, &u)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err

}
func (d MySqlDatastore) UpdateUser(userId int, u User) (error) {
	u.Id = int32(userId)
	_, err := d.db.NamedExec(`UPDATE who_users SET user_name1=:user_name1, user_name2=:user_name2, user_twitter=:user_twitter, user_email=:user_email, user_facebook=:user_facebook, user_tel=:user_tel, user_website=:user_website, user_google_plus=:user_google_plus WHERE user_id = :user_id `, &u)
	if err != nil {
		log.Println(err)
		return  err
	}
	return nil
}

func (d MySqlDatastore) GetDevicesByUserId(id int) ([]Device, error) {
	devices := []Device{}
	err := d.db.Select(&devices, "SELECT * FROM who_devices WHERE device_uid=?", id)
	if err != nil {
		return nil, fmt.Errorf("User with id %d not found", id)
	}

	return devices, nil
}

func (d MySqlDatastore) AddDevice(userId int, device Device) (int64, error) {
	_, notFound := d.GetUser(userId)
	if notFound != nil {
		return 0, notFound
	}
	device.UserId = userId
	res, err := d.db.NamedExec(`INSERT INTO who_devices (device_MAC, device_comment, device_uid) VALUES (:device_MAC, :device_comment, :device_uid)`, &device)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (d MySqlDatastore) GetDeviceByMac(mac string) (*Device, error) {
	device := Device{}
	err := d.db.Get(&device, "SELECT * FROM who_devices WHERE device_MAC=?", mac)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if device.Id == 0 {
		return nil, fmt.Errorf("Device with mac %s dosent exist", mac)
	}
	return &device, nil
}
