package controllers

import (
	"net/http"

	"github.com/poncorobbin/go-simple-rest/pkg/controllers/student"
)

func New(mux *http.ServeMux) {
	student.ActionStudent(mux)
}
