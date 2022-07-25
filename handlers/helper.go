package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/jabin/Chatplatm/config"
	"github.com/jabin/Chatplatm/models"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var logger *log.Logger
var config *Configuration
var localizer *i18n.Localizer

func init() {

	config = LoadConfig()
	localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)

	file, err := os.OpenFile("logs/jabinchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARINING ")
	logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func err_measage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

// 通过cookie 判断用户是否已经登录
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

// 解析HTML模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成相应的HTML
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)

}

// 返回版本号
func Version() string {
	return "0.1"
}
