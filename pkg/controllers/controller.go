package controllers

import (
	"net/http"

	"github.com/poncorobbin/go-simple-rest/pkg/controllers/classroom"
	"github.com/poncorobbin/go-simple-rest/pkg/controllers/student"
)

func New(mux *http.ServeMux) {
	student.ActionStudent(mux)
	classroom.ActionClassroom(mux)
}
