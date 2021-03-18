package timecamp

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timecamp/types"
)

type Task struct {
	TaskID           int64   `json:"task_id"`
	ParentID         *int64  `json:"parent_id"`
	AssignedBy       *int64  `json:"assigned_by"`
	Name             string  `json:"name"`
	ExternalTaskID   *string `json:"external_task_id"`
	ExternalParentID *string `json:"external_parent_id"`
	Level            int64   `json:"level"`
	//Archived         t_types.BoolInt `json:"archived"`
	Archived       int                     `json:"archived"`
	Tags           *string                 `json:"tags"`
	Budgeted       int64                   `json:"budgeted"`
	BudgetUnit     string                  `json:"budget_unit"`
	RootGroupID    *int64                  `json:"root_group_id"`
	Billable       t_types.BoolInt         `json:"billable"`
	Note           *string                 `json:"note"`
	PublicHash     *string                 `json:"public_hash"`
	AddDate        t_types.DateTimeString  `json:"add_date"`
	ModifyTime     *t_types.DateTimeString `json:"modify_time"`
	Color          string                  `json:"color"`
	UserAccessType int64                   `json:"user_access_type"`
}

// GetTasks returns all tasks
//
func (service *Service) GetTasks() (*map[string]Task, *errortools.Error) {
	tasks := make(map[string]Task)

	requestConfig := go_http.RequestConfig{
		URL:           service.url("tasks"),
		ResponseModel: &tasks,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &tasks, nil
}
