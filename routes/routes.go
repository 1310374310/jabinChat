package routes

import (
	"net/http"

	"github.com/jabin/Chatplatm/handlers"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		Name:        "home",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"newThread", // 访问功能请求
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread", // 提交创建表单
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
