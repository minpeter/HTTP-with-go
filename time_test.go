package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	//바디를 소비하는데 발생하는 오버해드 방지
	resp, err := http.Head("https://www.time.gov")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("no Date header received from time.gov")
	}
	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}
