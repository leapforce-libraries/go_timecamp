package timecamp

import (
	"fmt"
	"log"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
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
func (service *Service) GetUsers() (*[]User, *errortools.Error) {
	users := []User{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("users/format/json/api_token/%s", service.token)),
		ResponseModel: &users,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	for index := range users {
		users[index] = users[index].ParseDates()
	}

	return &users, nil
}

// ParseDates //
//
func (service User) ParseDates() User {
	// parse LoginTime to *bool
	if service.LoginTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.LoginTime)
		if err != nil {
			log.Println(err)
		}
		service.LoginTime2 = &_t
	}
	// parse SynchTime to time.Time
	if service.SynchTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.SynchTime)
		if err != nil {
			log.Println(err)
		}
		service.SynchTime2 = &_t
	}

	return service
}
