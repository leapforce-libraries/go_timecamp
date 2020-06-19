package timecamp

import (
	"fmt"
	"log"
	"time"
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
func (t *Timecamp) GetTasks() (map[string]Task, error) {

	urlStr := "%stasks/format/json/api_token/%s"

	url := fmt.Sprintf(urlStr, t.apiURL, t.token)
	//fmt.Printf(url)

	tasks := make(map[string]Task)

	err := t.Get(url, &tasks)
	if err != nil {
		return nil, err
	}

	for index := range tasks {
		tasks[index] = tasks[index].ParseDates().ParseBooleans()
	}

	return tasks, nil
}

// ParseDates //
//
func (t Task) ParseDates() Task {
	// parse AddDate to *bool
	if t.AddDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.AddDate)
		if err != nil {
			log.Println(err)
		}
		t.AddDate2 = &_t
	}
	// parse CheckedDate to time.Time
	if t.CheckedDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.CheckedDate)
		if err != nil {
			log.Println(err)
		}
		t.CheckedDate2 = &_t
	}
	// parse DueDate to time.Time
	if t.DueDate != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.DueDate)
		if err != nil {
			log.Println(err)
		}
		t.DueDate2 = &_t
	}
	// parse ModifyTime to time.Time
	if t.ModifyTime != "" {
		_t, err := time.Parse("2006-01-02 15:04:05", t.ModifyTime)
		if err != nil {
			log.Println(err)
		}
		t.ModifyTime2 = &_t
	}

	return t
}

// ParseBooleans //
//
func (t Task) ParseBooleans() Task {
	// parse Archived to *bool
	switch t.Archived {
	case "0":
		b := false
		t.Archived2 = &b
		break
	case "1":
		b := true
		t.Archived2 = &b
		break
	default:
		t.Archived2 = nil
		break
	}
	// parse Budgeted to *bool
	switch t.Budgeted {
	case "0":
		b := false
		t.Budgeted2 = &b
		break
	case "1":
		b := true
		t.Budgeted2 = &b
		break
	default:
		t.Budgeted2 = nil
		break
	}
	// parse Billable to *bool
	switch t.Billable {
	case "0":
		b := false
		t.Billable2 = &b
		break
	case "1":
		b := true
		t.Billable2 = &b
		break
	default:
		t.Billable2 = nil
		break
	}

	return t
}
