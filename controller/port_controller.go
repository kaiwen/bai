package controller

import (
	"log"
	"net/http"
	"strconv"

	"bagent.com/bagent/service"
)

func Init() {
	http.HandleFunc("/port", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if err := r.ParseForm(); err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}

			pass := r.Form.Get("pass")
			port, err := service.GetPort(pass)
			if err != nil {
				log.Println(err)
				w.WriteHeader(403)
				return
			}

			w.Write([]byte(strconv.Itoa(int(port))))
			return
		} else if r.Method == "POST" {
			portStr := r.FormValue("port")
			pass := r.FormValue("pass")
			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Println(err)
				w.WriteHeader(400)
				return
			}

			if err := service.SetPort(uint16(port), pass); err != nil {
				log.Println(err)
				w.WriteHeader(403)
				return
			}

			w.Write([]byte("success"))
			return
		}
	})

	http.HandleFunc("/pass", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			oldPass := r.FormValue("old_pass")
			newPass := r.FormValue("new_pass")
			if err := service.SetPass(newPass, oldPass); err != nil {
				log.Println(err)
				w.WriteHeader(403)
				return
			}

			w.Write([]byte("success"))
			return
		}
	})

	http.HandleFunc("/port_pass", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			oldPass := r.FormValue("old_pass")
			newPass := r.FormValue("new_pass")
			portStr := r.FormValue("port")
			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Println(err)
				w.WriteHeader(400)
				return
			}

			if err := service.SetPortAndPass(uint16(port), newPass, oldPass); err != nil {
				log.Println(err)
				w.WriteHeader(403)
				return
			}

			w.Write([]byte("success"))
			return
		}
	})
}
