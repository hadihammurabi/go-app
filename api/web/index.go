package web

import (
	"net/http"

	"github.com/gowok/gowok"
	"github.com/ngamux/ngamux"
)

func Index() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		token := "adacd8a852a0813c9bf8e7690f4461d56930c867e241c55eac0afa5d7dd9ac87"
		gowok.HttpOk(w, ngamux.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
			"token":   token,
		})
	}
}
