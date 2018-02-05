package db

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var randKeys chan string

func randKey() string {
	return <-randKeys
}

func randKeyBackend() {
	for {
		b := make([]byte, 16)
		rand.Read(b)
		randKeys <- fmt.Sprintf("%x", b)
	}
}

func initRandKey() {
	randKeys = make(chan string, 10)
	go randKeyBackend()
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
