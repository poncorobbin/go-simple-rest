package student

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var data = []Student{
	{Id: "1", Name: "ponco", Age: 24},
}

func ActionStudent() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
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
		res, err = json.Marshal(selectStudent(id))
	} else {
		res, err = json.Marshal(data)
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
	data = append(data, student)

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
	student = selectStudent(id)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	student.Name = payload.Name
	student.Age = payload.Age

	res, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func selectStudent(id string) Student {
	for _, each := range data {
		if each.Id == id {
			return each
		}
	}
	return Student{}
}
