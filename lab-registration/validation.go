package main

import (
	"net/http"

	"github.com/martini-contrib/binding"
)

func (u User) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	if len(u.FirstName) < 4 {
		errors = append(errors, binding.Error{
			FieldNames: []string{"firstname"},
			Classification: "Incorrect data",
			Message: "first name must have minimum length of 4 characters",
		})
	}
	return errors
}
