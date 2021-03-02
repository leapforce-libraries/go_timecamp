package timecamp

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Entry struct {
	ID               int    `json:"id"`
	Duration         string `json:"duration"`
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	TaskID           string `json:"task_id"`
	LastModify       string `json:"last_modify"`
	Date             string `json:"date"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	Locked           string `json:"locked"`
	Name             string `json:"name"`
	AddonsExternalID string `json:"addons_external_id"`
	Billable         int    `json:"billable"`
	InvoiceID        string `json:"invoiceId"`
	Color            string `json:"color"`
	Description      string `json:"description"`
	LastModify2      *time.Time
	Date2            *civil.Date
	StartTime2       *civil.Time
	EndTime2         *civil.Time
	Locked2          *bool
	Billable2        *bool
}

// GetEntriesByUserID returns all entries
//
func (service *Service) GetEntriesByUserID(userID string) (*[]Entry, *errortools.Error) {
	if service == nil {
		return nil, nil
	}

	startDateString := service.startDateEntries.Format("2006-01-02")
	endDateString := time.Now().Format("2006-01-02")

	entries := []Entry{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("entries/format/json/api_token/%s/user_ids/%s/from/%s/to/%s", service.token, userID, startDateString, endDateString)),
		ResponseModel: &entries,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	for index := range entries {
		entries[index] = entries[index].ParseDates().ParseBooleans()
	}

	return &entries, nil
}

// ParseDates //
//
func (e Entry) ParseDates() Entry {
	// parse LastModify to time.Time
	if e.LastModify != "" {
		_e, err := time.Parse("2006-01-02 15:04:05", e.LastModify)
		if err != nil {
			log.Println(err)
		}
		e.LastModify2 = &_e
	}
	// parse LastModify to time.Time
	if e.Date != "" {
		_e, err := time.Parse("2006-01-02", e.Date)
		if err != nil {
			log.Println(err)
		}
		e.Date2 = &civil.Date{_e.Year(), _e.Month(), _e.Day()}
	}
	// parse StartTime to time.Time
	if e.Date != "" && e.StartTime != "" {
		_e, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", e.Date, e.StartTime))
		if err != nil {
			log.Println(err)
		}
		e.StartTime2 = &civil.Time{_e.Hour(), _e.Minute(), _e.Second(), 0}

	}
	// parse EndTime to time.Time
	if e.Date != "" && e.EndTime != "" {
		_e, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", e.Date, e.EndTime))
		if err != nil {
			log.Println(err)
		}
		e.EndTime2 = &civil.Time{_e.Hour(), _e.Minute(), _e.Second(), 0}
	}

	return e
}

// ParseBooleans //
//
func (e Entry) ParseBooleans() Entry {
	// parse Locked to *bool
	switch e.Locked {
	case "0":
		b := false
		e.Locked2 = &b
		break
	case "1":
		b := true
		e.Locked2 = &b
		break
	default:
		e.Locked2 = nil
		break
	}
	// parse Billable to *bool
	switch e.Billable {
	case 0:
		b := false
		e.Billable2 = &b
		break
	case 1:
		b := true
		e.Billable2 = &b
		break
	default:
		e.Billable2 = nil
		break
	}

	return e
}
