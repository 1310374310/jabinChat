package handlers

// 定义首页处理器

import (
	"log"
	"net/http"

	"github.com/jabin/Chatplatm/models"
)

// 保持与HandlerFunc 的参数一致
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err = session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	} else {
		log.Fatal(err)
	}
}

func Err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
