package models

import (
	"time"
)

type PasswordEntry struct {
    ServiceName string    `json:"service_name"`
    Folder      string    `json:"folder"`
    LogIn       string    `json:"login"`
    Password    string    `json:"password"`
    CreatedAt   time.Time `json:"created_at"`
}