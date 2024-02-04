package user

import (
    "time"
)

type User struct {
    userId      int
    userName    string
    email       string
    password    string
    createdDate time.Time
}
