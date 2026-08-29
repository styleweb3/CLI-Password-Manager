package models

import (
	"time"
)

type PasswordEntry struct {
	ServiceName string
	Folder		string
	LogIn 		string
	Password	string
	CreatedAt 	time.Time
}