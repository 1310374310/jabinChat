package models

import (
	"fmt"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// formate the CreateAt data to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	local, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
		return "error time"
	}

	return thread.CreatedAt.In(local).Format("2006-01-02 15:04:05")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id=?", thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id,uuid,body,user_id,thread_id,created_at from posts where thread_id=?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get all threads in the database and return it
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id,uuid,topic,user_id,created_at from threads order by created_at")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		t := Thread{}
		err = rows.Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		threads = append(threads, t)
	}
	rows.Close()
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (t Thread, err error) {
	t = Thread{}
	err = Db.QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid=?", uuid).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	if err != nil {
		fmt.Printf("Cannot get thread, err: %v\n", err)
	}
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	err := Db.QueryRow("select id,uuid,name,email,created_at from users where id=?", thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		fmt.Printf("cannot get user, err: %v\n", err)
	}
	return
}
