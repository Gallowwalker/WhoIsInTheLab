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
	err := d.db.Get(&user, "SELECT * FROM who_users WHERE user_id=?", id)
	if err != nil {
		return user, fmt.Errorf("User with id %d not found", id)
	}

	return user, nil
}

func (d MySqlDatastore) AddUser(u User) (int64, error) {
	res, err := d.db.NamedExec(`INSERT INTO who_users (user_name1, user_name2, user_twitter, user_facebook, user_tel, user_website, user_fstoken, user_google_plus, user_fscheckin) VALUES (:user_name1, :user_name2, :user_twitter, :user_facebook, :user_tel, :user_website, :user_fstoken, :user_google_plus, :user_fscheckin)`, &u)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err

}

