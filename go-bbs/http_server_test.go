package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PichuChen/go-bbs/api"
)

func Test_Login(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		path   string
		params interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{path: "/login", params: &api.LoginParams{UserID: "SYSOP", Passwd: "123123"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router, _ := initGin()

			w := httptest.NewRecorder()
			jsonStr, _ := json.Marshal(tt.args.params)
			req, _ := http.NewRequest("POST", tt.args.path, bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			bodyStr := w.Body.String()
			if !strings.Contains(bodyStr, "Jwt") {
				t.Errorf("initGin(): %v", bodyStr)
			}
		})
	}
}

func Test_Ping(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		path   string
		params interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:    args{path: "/ping", params: &LoginRequiredParams{UserID: "SYSOP", Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmUiOjE2MDU0OTYzMzQsIlVzZXJJRCI6IlNZU09QIn0.33UbL2z85_w9Z84HWyAKnYWG9omWPyMPNJwIHnV6Aa0", Data: nil}},
			wantErr: false,
		},

		{
			args:    args{path: "/ping", params: &LoginRequiredParams{Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmUiOjE2MDU0OTYzMzQsIlVzZXJJRCI6IlNZU09QIn0.33UbL2z85_w9Z84HWyAKnYWG9omWPyMPNJwIHnV6Aa0", Data: nil}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router, _ := initGin()

			w := httptest.NewRecorder()
			jsonStr, _ := json.Marshal(tt.args.params)
			req, _ := http.NewRequest("POST", tt.args.path, bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			if (w.Code != http.StatusOK) != tt.wantErr {
				t.Errorf("code: %v wantErr: %v", w.Code, tt.wantErr)
			}
		})
	}
}
