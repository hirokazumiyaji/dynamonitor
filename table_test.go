package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func init() {
	client = NewClient(
		&Config{
			Aws: AwsConfig{
				Region:          "ap-northeast-1",
				AccessKeyId:     "test",
				SecretAccessKey: "test",
				EndpointUrl:     "http://127.0.0.1:8000",
			},
		},
	)
}

func TestTableList(t *testing.T) {
	m := web.New()
	m.Get("/api/tables", tableList)

	ts := httptest.NewServer(m)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/tables")
	if err != nil {
		t.Error("response error", err.Error())
	}
	fmt.Println(res)
}

func TestTableDetail(t *testing.T) {
	m := web.New()
	m.Get("/api/tables/:tablename", tableDetail)

	ts := httptest.NewServer(m)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/tables/example")
	if err != nil {
		t.Error("response error", err.Error())
	}
	fmt.Println(res)
}
