package timecamp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	types "github.com/Leapforce-nl/go_types"
)

// type
//
type Timecamp struct {
	token            string
	apiURL           string
	startDateEntries *time.Time
}

// New //
//
func New(apiURL string, token string, startDateEntries *time.Time) (*Timecamp, error) {
	i := new(Timecamp)

	if apiURL == "" {
		return nil, &types.ErrorString{"Timecamp ApiUrl not provided"}
	}
	if token == "" {
		return nil, &types.ErrorString{"Timecamp Token not provided"}
	}

	i.apiURL = apiURL
	i.token = token
	i.startDateEntries = startDateEntries

	if !strings.HasSuffix(i.apiURL, "/") {
		i.apiURL = i.apiURL + "/"
	}

	return i, nil
}

// Get //
//
func (i *Timecamp) Get(url string, model interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json")
	//req.Header.Set("authorization", "Basic "+i.token)

	// Send out the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	errr := json.Unmarshal(b, &model)
	if errr != nil {
		return err
	}

	return nil
}
