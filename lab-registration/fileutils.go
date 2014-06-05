package main

import(
	"io/ioutil"
	"log"
)

func ReadFile(filename string) (string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file %v", filename)
	}
	return string(buf)
}
