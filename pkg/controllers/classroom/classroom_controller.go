package classroom

import "net/http"

func ActionClassroom(mux *http.ServeMux) {
	mux.HandleFunc("/classroom", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGet(w, r)
		}
	})
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is classroom"))
}
