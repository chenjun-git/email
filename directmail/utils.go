package directmail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//control over the request lifecycle
var httpClient = &http.Client{
	Timeout: 2000 * time.Millisecond,
}

func SetClientTimeout(duration time.Duration) {
	httpClient.Timeout = duration
}

func httpReqWithParams(method, url string, v interface{}) ([]byte, error) {
	parmas, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(parmas)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
