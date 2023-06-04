package main

import (
	_ "github.com/go-redis/redis/v9" // "_"为空导入
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

func main() {
    logrus.Println("hello, gopath mode")
		logrus.Println(uuid.NewString())
}