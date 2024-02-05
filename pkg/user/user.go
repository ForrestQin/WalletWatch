package user

import (
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserId      primitive.ObjectID `json:"user_id,omitempty" bson:"_id,omitempty"`
	UserName    string             `json:"user_name"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	CreatedDate time.Time          `json:"created_date"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		CreatedDate string `json:"created_date"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.CreatedDate = time.Now()
	return nil
}
