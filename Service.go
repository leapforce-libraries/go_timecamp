package timecamp

import (
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName string = "TimeCamp"
	apiUrl  string = "https://app.timecamp.com/third_party/api"
)

// type
//
type Service struct {
	token            string
	startDateEntries *time.Time
	httpService      *go_http.Service
}

type ServiceConfig struct {
	Token            string
	StartDateEntries *time.Time
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	if config.Token == "" {
		return nil, errortools.ErrorMessage("Token not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		token:            config.Token,
		startDateEntries: config.StartDateEntries,
		httpService:      httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Bearer %s", service.token))
	header.Set("Accept", "application/json")
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if errorResponse.Message != "" {
		e.SetMessage(errorResponse.Message)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.token
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
