package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func call(into interface{}, args map[string]string) error {
	uri := fmt.Sprintf("%s/sabnzbd/api?output=json&apikey=%s", *sabnzbdAddress, *sabnzbdAPIKey)
	for k, v := range args {
		uri = fmt.Sprintf("%s&%s=%s", uri, k, v)
	}
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bs, into)
}
