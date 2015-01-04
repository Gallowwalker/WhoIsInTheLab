package main

import (
	"encoding/json"
)

type jsonEncoder struct {
}

func (e jsonEncoder) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}
