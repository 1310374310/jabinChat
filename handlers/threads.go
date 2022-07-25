package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jabin/Chatplatm/models"
)

// GET /threads/new
// 创建群组页面
func NewThread(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /threads/new")
	_, err := session(w, r) // 判断是否登录
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "auth.navbar", "new.thread") // 跳转到创建组群页面
	}
}

// POST   /thread/create
// 执行组群创建逻辑
func CreateThread(w http.ResponseWriter, r *http.Request) {
	log.Println("POST   /thread/create")
	sess, err := session(w, r) // 判断是否登录
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			fmt.Printf("Cannot parse form, err:%v\n", err)
		}

		user, err := sess.User()
		if err != nil {
			fmt.Printf("Cannot get user from session,err:%v\n", err)
		}
		topic := r.PostFormValue("topic")
		if _, err = user.CreateThread(topic); err != nil {
			fmt.Printf("Cannot create thread,err:%v\n", err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

// GET /thread/read
// 通过ID渲染指定群组页面(群组详情)
func ReadThread(w http.ResponseWriter, r *http.Request) {
	info("GET /thread/read")
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		err_measage(w, r, "Cannnot read thread")
	} else {
		_, err := session(w, r)
		if err != nil {
			fmt.Println(err)
			generateHTML(w, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
