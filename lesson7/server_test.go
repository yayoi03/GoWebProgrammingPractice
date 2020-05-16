package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

// func TestHandleGet(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest)

// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/post/1", nil)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	if post.Id != 1 {
// 		t.Error("Cannot retrieve JSON post")
// 	}
// }

// func TestHandlePut(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest)

// 	writer := httptest.NewRecorder()
// 	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
// 	request, _ := http.NewRequest("PUT", "/post/1", json)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}

// }

// var mux *http.ServeMux
// var writer *httptest.ResponseRecorder

// func TestMain(m *testing.M) {
// 	setUp()
// 	code := m.Run()
// 	os.Exit(code)
// }

// func setUp() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest)
// 	writer := httptest.NewRecorder()
// }

// type PostTestSuite struct{}

// func init() {
// 	Suite(&PostTestSuite{})
// }

// func Test(t *testing.T) { TestingT(t) }

// func (s *PostTestSuite) TestHandleGet(c *C) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
// 	writer := httptest.NewRecorder()

// 	request, _ := http.NewRequest("GET", "/post/1", nil)
// 	mux.ServeHTTP(writer, request)

// 	c.Check(writer.Code, Equals, 200)

// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	c.Check(post.Id, Equals, 1)
// }

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)

	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(`{"content":"Upload post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Content, Equals, "Upload post")
}

// func TestHandleGet(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
// 	writer := httptest.NewRecorder()

// 	request, _ := http.NewRequest("GET", "/post/1", nil)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	if post.Id != 1 {
// 		t.Error("Cannot retrieve JSON post")
// 	}
// }

// func TestHandlePut(t *testing.T) {
// 	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
// 	request, _ := http.NewRequest("PUT", "/post/1", json)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}

// }
