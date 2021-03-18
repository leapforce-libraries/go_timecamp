package timecamp

import (
	"fmt"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_timecamp/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type Entry struct {
	ID               int64                  `json:"id"`
	Duration         go_types.Int64String   `json:"duration"`
	UserID           go_types.Int64String   `json:"user_id"`
	UserName         string                 `json:"user_name"`
	TaskID           go_types.Int64String   `json:"task_id"`
	LastModify       t_types.DateTimeString `json:"last_modify"`
	Date             t_types.DateString     `json:"date"`
	StartTime        t_types.TimeString     `json:"start_time"`
	EndTime          t_types.TimeString     `json:"end_time"`
	Locked           t_types.BoolString     `json:"locked"`
	Name             string                 `json:"name"`
	AddOnsExternalID string                 `json:"addons_external_id"`
	Billable         t_types.BoolInt        `json:"billable"`
	InvoiceID        go_types.Int64String   `json:"invoiceId"`
	Color            string                 `json:"color"`
	Description      string                 `json:"description"`
}

// GetEntriesByUserID returns all entries
//
func (service *Service) GetEntriesByUserID(userID int64) (*[]Entry, *errortools.Error) {
	if service == nil {
		return nil, nil
	}

	values := url.Values{}
	values.Set("user_ids", fmt.Sprintf("%v", userID))
	values.Set("from", service.startDateEntries.Format("2006-01-02"))
	values.Set("to", time.Now().Format("2006-01-02"))

	entries := []Entry{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("entries?%s", values.Encode())),
		ResponseModel: &entries,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &entries, nil
}
