package timecamp

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timecamp/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type User struct {
	GroupID     go_types.Int64String   `json:"group_id"`
	UserID      go_types.Int64String   `json:"user_id"`
	Email       string                 `json:"email"`
	LoginCount  go_types.Int64String   `json:"login_count"`
	LoginTime   t_types.DateTimeString `json:"login_time"`
	DisplayName string                 `json:"display_name"`
	SynchTime   t_types.DateTimeString `json:"synch_time"`
}

// GetUsers returns all users
//
func (service *Service) GetUsers() (*[]User, *errortools.Error) {
	users := []User{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("users"),
		ResponseModel: &users,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &users, nil
}
