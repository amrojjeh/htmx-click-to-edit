package main

import (
	"net/http"
	"html/template"
)

type userInfo struct {
	FirstName string
	LastName string
	Email string
}

func indexHandler(t *template.Template, info *userInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		err := t.Execute(w, info)
		if err != nil {
			panic(err)
		}
	}
}

func editHandler(t *template.Template, info *userInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		err := t.ExecuteTemplate(w, "contact/1/edit", info)
		if err != nil {
			panic(err)
		}
	}
}

func infoViewHandler(t *template.Template, info *userInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			
			if !(r.Form.Has("firstName") &&
				r.Form.Has("lastName") &&
				r.Form.Has("email")) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			info.FirstName = r.Form.Get("firstName")
			info.LastName = r.Form.Get("lastName")
			info.Email = r.Form.Get("email")
		}

		err := t.ExecuteTemplate(w, "contact/1", info)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	info := &userInfo{"Joe", "Blow", "jow@blow.com"}
	t := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", indexHandler(t, info))
	http.HandleFunc("/contact/1/edit", editHandler(t, info))
	http.HandleFunc("/contact/1", infoViewHandler(t, info))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
