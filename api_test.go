package main_test

import (
	"fmt"
	"github.com/iamxuwenjin/blog/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	baseUrl     = "http://127.0.0.1:8000/api/v1/tags"
	contentType = "application/x-www-form-urlencoded"
	w           = httptest.NewRecorder()
	router      = routers.InitRouter()
)

func readContent(rsp *httptest.ResponseRecorder) (content []byte) {
	content, _ = ioutil.ReadAll(rsp.Body)
	return
}

func TestGetTags(t *testing.T) {
	req, _ := http.NewRequest("GET", baseUrl, nil)
	router.ServeHTTP(w, req)
	fmt.Println(string(readContent(w)))
}

func TestAddTags(t *testing.T) {
	//p := make(map[string]string)
	//p["name"] = "python"
	//p["createdBy"] = "Bob"
	//p["state"] = "1"
	p := "?name=python&state=1&created_by=Bob"
	//jsonStr, _ := json.Marshal(p)
	//rsp, err := c.Post(baseUrl+"/api/v1/tags", contentType, bytes.NewBuffer([]byte(p)))
	req, _ := http.NewRequest("POST", baseUrl+p, nil)
	req.Header.Add("Content-Type", contentType)
	router.ServeHTTP(w, req)
	fmt.Println(string(readContent(w)))
}

func TestEditTags(t *testing.T) {
	p := "/4?name=python&state=1&modified_by=jack"
	req, _ := http.NewRequest("PUT", baseUrl+p, nil)
	req.Header.Add("Content-Type", contentType)
	router.ServeHTTP(w, req)
	fmt.Println(string(readContent(w)))
}
