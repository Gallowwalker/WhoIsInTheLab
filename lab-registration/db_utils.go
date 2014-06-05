package main

import (
	"os/exec"
)

func ExecuteSql(sqlFile string, dbcfg Config) (string, error) {
	mysqlOut, err := exec.Command("sh" , "-c", "mysql" +  " -u " + dbcfg.Username + " -p" + dbcfg.Password + " " + dbcfg.Database + " < " + sqlFile).Output()
	if err != nil {
		return  "", err
	}
	return string(mysqlOut[:]), nil
}
