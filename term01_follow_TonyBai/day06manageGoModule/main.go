package main

import (
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

func main() {
    logrus.Println("hello, gopath mode")
		logrus.Println(uuid.NewString())
}