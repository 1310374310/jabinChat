package test

import (
	"testing"

	"github.com/jabin/Chatplatm/models"
)

func TestCreate(t *testing.T) {
	user := models.User{
		Name:     "jabin",
		Password: "123456",
		Email:    "1310374310@qq.com",
	}

	err := user.Create()
	if err != nil {
		t.Fatalf("error:%v", err)
	}
}
