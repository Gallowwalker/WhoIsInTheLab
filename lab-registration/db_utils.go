package main

import (
	"os/exec"
)

func ExecuteSql(sqlFile string, dbcfg Config) (string, error) {
	password := " -p" + dbcfg.Password
	if len(dbcfg.Password) == 0 {
		password = ""
	}
	mysqlOut, err := exec.Command("sh" , "-c", "mysql" +  " -u " + dbcfg.Username + password + " " + dbcfg.Database + " < " + sqlFile).Output()
	if err != nil {
		return  "", err
	}
	return string(mysqlOut[:]), nil
}
