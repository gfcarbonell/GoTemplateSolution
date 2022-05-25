package Utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpClient(method string, url string, body interface{}, headers map[string]string) ([]byte, int, error) {
	fmt.Println("Calling API...")
	data, err := json.Marshal(body)

	if err != nil {
		fmt.Print(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	bodyByte := bytes.NewBuffer(data)
	req, err := http.NewRequest(method, url, bodyByte)

	if err != nil {
		fmt.Print(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{
		Timeout: SETTING.HttpClient.Timeout * time.Second,
	}

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	return bodyBytes, response.StatusCode, nil
}
