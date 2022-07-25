package models

import (
	"fmt"
	"log"
	"time"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (post *Post) CreatedAtDate() string {
	local, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
		return "error time"
	}
	return post.CreatedAt.In(local).Format("2006-01-02 15:04:05")
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	err := Db.QueryRow("select id,uuid,name,email,created_at from users where id=?", post.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
