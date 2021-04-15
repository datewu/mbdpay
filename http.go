package mbdpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getJSON(resp *http.Response, container interface{}) error {
	if resp.StatusCode > http.StatusPermanentRedirect {
		return fmt.Errorf("bad http response status: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(container)
}

// p should be a pointer type
func postJSON(url string, payload interface{}, p interface{}) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)

	ct := "application/json; charset=UTF-8"
	resp, err := http.Post(url, ct, r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return getJSON(resp, p)
}
