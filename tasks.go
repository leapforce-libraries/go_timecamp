package timecamp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timecamp/types"
	go_types "github.com/leapforce-libraries/go_types"
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
	Billable       go_types.BoolInt        `json:"billable"`
	Note           *string                 `json:"note"`
	PublicHash     *string                 `json:"public_hash"`
	AddDate        t_types.DateTimeString  `json:"add_date"`
	ModifyTime     *t_types.DateTimeString `json:"modify_time"`
	Color          string                  `json:"color"`
	UserAccessType int64                   `json:"user_access_type"`
}

type TaskPermission string

const (
	TaskPermissionCreateSubtask    TaskPermission = "create_subtask"
	TaskPermissionEditTaskSettings TaskPermission = "edit_task_settings"
	TaskPermissionTrackTime        TaskPermission = "track_time"
	TaskPermissionViewDetailedData TaskPermission = "view_detailed_data"
)

type TaskStatus string

const (
	TaskStatusActive   TaskStatus = "active"
	TaskStatusArchived TaskStatus = "archived"
	TaskStatusAll      TaskStatus = "all"
)

type GetTasksConfig struct {
	ExternalTaskID *string
	Minimal        *bool
	Permissions    *[]TaskPermission
	Status         *TaskStatus
	TaskID         *int64
}

// GetTasks returns all tasks
//
func (service *Service) GetTasks(config *GetTasksConfig) (*map[string]Task, *errortools.Error) {
	params := url.Values{}

	if config != nil {
		if config.ExternalTaskID != nil {
			params.Add("external_task_id", *config.ExternalTaskID)
		}
		if config.Minimal != nil {
			if *config.Minimal {
				params.Add("minimal", "1")
			} else {
				params.Add("minimal", "0")
			}
		}
		if config.Permissions != nil {
			if len(*config.Permissions) > 0 {
				_permissions := []string{}
				for _, permission := range *config.Permissions {
					_permissions = append(_permissions, string(permission))
				}
				params.Add("perms", strings.Join(_permissions, ","))
			}
		}
		if config.Status != nil {
			params.Add("status", string(*config.Status))
		}
		if config.TaskID != nil {
			params.Add("task_id", fmt.Sprintf("%v", *config.TaskID))
		}
	}

	tasks := make(map[string]Task)

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("tasks?%s", params.Encode())),
		ResponseModel: &tasks,
	}
	_, _, e := service.httpService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &tasks, nil
}
