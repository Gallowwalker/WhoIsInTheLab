package main

type encoder interface {
	Encode(obj interface{}) ([]byte, error)
}

func must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
