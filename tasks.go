package timecamp

import (
	"fmt"
	"log"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Task struct {
	TaskID           string `json:"task_id"`
	ParentID         string `json:"parent_id"`
	Name             string `json:"name"`
	ExternalTaskID   string `json:"external_task_id"`
	ExternalParentID string `json:"external_parent_id"`
	Level            string `json:"level"`
	AddDate          string `json:"add_date"`
	Archived         string `json:"archived"`
	Color            string `json:"color"`
	Tags             string `json:"tags"`
	Budgeted         string `json:"budgeted"`
	CheckedDate      string `json:"checked_date"`
	RootGroupID      string `json:"root_group_id"`
	AssignedTo       string `json:"assigned_to"`
	AssignedBy       string `json:"assigned_by"`
	DueDate          string `json:"due_date"`
	Note             string `json:"note"`
	Context          string `json:"context"`
	Folder           string `json:"folder"`
	Repeat           string `json:"repeat"`
	Billable         string `json:"billable"`
	BudgetUnit       string `json:"budget_unit"`
	PublicHash       string `json:"public_hash"`
	ModifyTime       string `json:"modify_time"`
	UserAccessType   int    `json:"user_access_type"`
	AddDate2         *time.Time
	Archived2        *bool
	Budgeted2        *bool
	CheckedDate2     *time.Time
	DueDate2         *time.Time
	Billable2        *bool
	ModifyTime2      *time.Time
}

// GetTasks returns all tasks
//
func (service *Service) GetTasks() (*map[string]Task, *errortools.Error) {
	tasks := make(map[string]Task)

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("tasks/format/json/api_token/%s", service.token)),
		ResponseModel: &tasks,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	for index := range tasks {
		tasks[index] = tasks[index].ParseDates().ParseBooleans()
	}

	return &tasks, nil
}

// ParseDates //
//
func (service Task) ParseDates() Task {
	// parse AddDate to *bool
	if service.AddDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.AddDate)
		if err != nil {
			log.Println(err)
		}
		service.AddDate2 = &_t
	}
	// parse CheckedDate to time.Time
	if service.CheckedDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.CheckedDate)
		if err != nil {
			log.Println(err)
		}
		service.CheckedDate2 = &_t
	}
	// parse DueDate to time.Time
	if service.DueDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.DueDate)
		if err != nil {
			log.Println(err)
		}
		service.DueDate2 = &_t
	}
	// parse ModifyTime to time.Time
	if service.ModifyTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", service.ModifyTime)
		if err != nil {
			log.Println(err)
		}
		service.ModifyTime2 = &_t
	}

	return service
}

// ParseBooleans //
//
func (service Task) ParseBooleans() Task {
	// parse Archived to *bool
	switch service.Archived {
	case "0":
		b := false
		service.Archived2 = &b
		break
	case "1":
		b := true
		service.Archived2 = &b
		break
	default:
		service.Archived2 = nil
		break
	}
	// parse Budgeted to *bool
	switch service.Budgeted {
	case "0":
		b := false
		service.Budgeted2 = &b
		break
	case "1":
		b := true
		service.Budgeted2 = &b
		break
	default:
		service.Budgeted2 = nil
		break
	}
	// parse Billable to *bool
	switch service.Billable {
	case "0":
		b := false
		service.Billable2 = &b
		break
	case "1":
		b := true
		service.Billable2 = &b
		break
	default:
		service.Billable2 = nil
		break
	}

	return service
}
