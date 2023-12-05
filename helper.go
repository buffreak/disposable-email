package disposable

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func Request(method, link string, body io.Reader, header http.Header, opts ...interface{}) (*http.Response, error) {
	defer HandlePanic("helper.Request")
	client := http.Client{}
	request, err := http.NewRequest(method, link, body)
	if err != nil {
		return nil, errors.New("failed initialize request")
	}
	if header != nil {
		request.Header = header
	}
	return client.Do(request)
}

func HandlePanic(name string) {
	if r := recover(); r != nil {
		log.Printf("Recover from %s\n", name)
	}
}
