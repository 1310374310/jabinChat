package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jabin/Chatplatm/models"
)

// GET /login
// 登录页面
func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(w, nil)
}

// GET /signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	log.Println("GET /signup")
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	info("POST /signup")
	err := r.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}

	user := models.User{
		Name:      r.PostFormValue("name"),
		Email:     r.PostFormValue("email"),
		Password:  r.PostFormValue("password"),
		CreatedAt: time.Now(),
	}

	if err = user.Create(); err != nil {
		msg := fmt.Sprintf("Cannot create user, err:%v", err)
		danger(msg)
	}

	http.Redirect(w, r, "/login", 302) //临时重定向
}

// POST /authenticate
// 通过邮箱和密码字段堆用户进行认证
func Authenticate(w http.ResponseWriter, r *http.Request) {
	info("POST /authenticate")
	err := r.ParseForm()
	if err != nil {
		danger("Cannot parse form")
	}

	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger("Cannot find user")
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger("Cannot create session")
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// 登录成功跳转首页
		http.Redirect(w, r, "/", 302)
	} else {

		// 登录失败跳转回登录页码
		http.Redirect(w, r, "/login", 302)
	}
}

// GET /logout
// 用户退出
func Logout(w http.ResponseWriter, r *http.Request) {
	info("GET /logout")
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	} else {
		danger("Failed to get cookie")
	}

	http.Redirect(w, r, "/", 302)

}
