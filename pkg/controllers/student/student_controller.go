package student

import (
	"encoding/json"
	"github.com/poncorobbin/go-simple-rest/pkg/db"
	. "github.com/poncorobbin/go-simple-rest/pkg/models"
	"net/http"
)

var studentRepo db.Repo[Student] = &Student{}

func ActionStudent(mux *http.ServeMux) {
	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGet(w, r)
		} else if r.Method == "POST" {
			handlePOST(w, r)
		} else if r.Method == "PUT" {
			handlePUT(w, r)
		}
	})
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	var (
		res []byte
		err error
	)

	if id := r.URL.Query().Get("id"); id != "" {
		res, err = json.Marshal(studentRepo.FindOne(id))
	} else {
		res, err = json.Marshal(studentRepo.Find())
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := studentRepo.Save(student)

	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func handlePUT(w http.ResponseWriter, r *http.Request) {
	var student, payload Student

	id := r.URL.Query().Get("id")
	student = studentRepo.FindOne(id)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	student.Name = payload.Name
	student.Age = payload.Age

	res, err := json.Marshal(studentRepo.Save(student))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
