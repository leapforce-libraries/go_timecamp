package timecamp

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	GroupID     string `json:"group_id"`
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	LoginCount  string `json:"login_count"`
	LoginTime   string `json:"login_time"`
	DisplayName string `json:"display_name"`
	SynchTime   string `json:"synch_time"`
	LoginTime2  *time.Time
	SynchTime2  *time.Time
}

// GetUsers returns all users
//
func (t *Timecamp) GetUsers() ([]User, error) {

	urlStr := "%susers/format/json/api_token/%s"

	url := fmt.Sprintf(urlStr, t.apiURL, t.token)
	//fmt.Printf(url)

	users := []User{}

	err := t.Get(url, &users)
	if err != nil {
		return nil, err
	}

	for index := range users {
		users[index] = users[index].ParseDates()
	}

	return users, nil
}

// ParseDates //
//
func (t User) ParseDates() User {
	// parse LoginTime to *bool
	if t.LoginTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.LoginTime)
		if err != nil {
			log.Println(err)
		}
		t.LoginTime2 = &_t
	}
	// parse SynchTime to time.Time
	if t.SynchTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.SynchTime)
		if err != nil {
			log.Println(err)
		}
		t.SynchTime2 = &_t
	}

	return t
}
