package models

import (
	"fmt"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// 检查session是否有效
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	if err != nil {
		valid = false
		return
	}

	if session.Id != 0 {
		valid = true
	}
	return
}

// 从数据库中删除一个session
func (session Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 通过session 获取用户信息
func (session Session) User() (user User, err error) {
	user = User{}
	// TODO 后续可以使用ORM模块实现
	err = Db.QueryRow("select id,uuid,name, email, created_at from users where id=?", session.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return

}

// 删库跑路
func SessionDeleteAll() (err error) {
	statment := "delete from sessions"
	_, err = Db.Exec(statment)
	return
}
