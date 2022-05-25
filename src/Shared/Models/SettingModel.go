package Models

import "time"

type Setting struct {
	SecreyKey string
	Server    struct {
		Port int
	}
	Database struct {
		Host         string
		Port         int
		Username     string
		Password     string
		DatabaseName string
	}
	HttpClient struct {
		Timeout time.Duration
	}
	URL struct {
		Templates string
	}
}
