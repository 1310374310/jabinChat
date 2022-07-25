package models

import (
	"fmt"
	"time"
)

/*
	用户模型
*/

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// 为用户创建一个新的session
func (user *User) CreateSession() (session Session, err error) {
	statment := "insert into sessions (uuid,email,user_id,created_at) values (?,?,?,?)"
	stmtin, err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	_, err = stmtin.Exec(uuid, user.Email, user.Id, time.Now())
	if err != nil {
		return
	}

	stmtout, err := Db.Prepare("select id,uuid,email,user_id,created_at from sessions where uuid=?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return
	}

	return
}

// 获取用户的session
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from users where user_id=?", user.Id).Scan(&session.Id, &session.Uuid, &session.Email, &session.CreatedAt)
	if err != nil {
		return
	}
	return
}

// 创建一个新用户， 保存用户信息
func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password,created_at) values (?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	_, err = stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), user.CreatedAt)

	stmtout, err := Db.Prepare("select id,uuid,created_at from users where uuid=?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// 从数据库中删除特定用户
func (user *User) Delete() (err error) {
	statement := "delete from users where id=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uuid)
	return
}

// 更新用户信息
func (user *User) Update() (err error) {
	statement := "update users set name = ?, email=? where id=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return

}

// Delete all users from database
func UserDeleteAll() (err error) {
	statment := "delete from users"
	_, err = Db.Exec(statment)
	return
}

// Get all users in the database and returns them
func Users() (users []User, err error) {
	rows, err := Db.Query("select id,uuid,name,email,password,created_at from users")
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	rows.Close()
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name,email,password,created_at from users where email=?", email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name,email,password,created_at from users where uuid=?", uuid).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (t Thread, err error) {
	statement := "insert into threads (uuid,topic,user_id,created_at) values (?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id,uuid,topic,user_id,created_at from threads where uuid=?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// Create a new post to a thread
func (user *User) CreatePost(t Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid,body,user_id,thread_id,created_at) values (?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	_, err = stmtin.Exec(uuid, body, user.Id, t.Id, time.Now())

	stmtout, err := Db.Prepare("select id,uuid,body,user_id,created_at from posts where uuid=?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.CreatedAt)
	return
}
