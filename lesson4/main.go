package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/headers", headers)
	// server.ListenAndServe()

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/body", body)
	// server.ListenAndServe()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http.HandleFunc("/process", process)
	// http.HandleFunc("/write", writeExample)
	// http.HandleFunc("/writeheader", writeHeaderExample)
	// http.HandleFunc("/redirect", headerExample)
	// http.HandleFunc("/json", jsonExample)

	// http.HandleFunc("/set_cookie", setCookie)
	// http.HandleFunc("/get_cookie", getCookie)

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}

// func headers(w http.ResponseWriter, r *http.Request) {
// 	// h := r.Header["Accept-Encoding"]
// 	h := r.Header.Get("Accept-Encoding")
// 	fmt.Fprintln(w, h)
// }

// func body(w http.ResponseWriter, r *http.Request) {
// 	len := r.ContentLength
// 	body := make([]byte, len)
// 	r.Body.Read(body)
// 	fmt.Fprintln(w, string(body))
// }

// func process(w http.ResponseWriter, r *http.Request) {
// 	// r.ParseForm()
// 	// // fmt.Fprintln(w, r.PostForm)
// 	// fmt.Fprintln(w, r.FormValue("hello"))
// 	// fmt.Fprintln(w, r.Form)

// 	r.ParseMultipartForm(1024)
// 	fileHeader := r.MultipartForm.File["uploaded"][0]
// 	file, err := fileHeader.Open()
// 	if err == nil {
// 		data, err := ioutil.ReadAll(file)
// 		if err == nil {
// 			fmt.Fprintln(w, string(data))
// 		}
// 	}
// }

// func writeExample(w http.ResponseWriter, r *http.Request) {
// 	str := `<html>
// 	<head><title>Go Web App</title></head>
// 	<body><h1>Hello world</h1></body>
// 	</html>`
// 	w.Write([]byte(str))
// }

// func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(501)
// 	fmt.Fprintln(w, "そのようなサービスはございません。他を当たってください")
// }

// func headerExample(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Location", "http://google.com")
// 	w.WriteHeader(302)
// }

// type Post struct {
// 	User    string
// 	Threads []string
// }

// func jsonExample(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	post := &Post{
// 		User:    "yayoi",
// 		Threads: []string{"first", "second", "third"},
// 	}
// 	json, _ := json.Marshal(post)
// 	w.Write(json)
// }

// func setCookie(w http.ResponseWriter, r *http.Request) {
// 	c1 := http.Cookie{
// 		Name:     "first_cookie",
// 		Value:    "Web App",
// 		HttpOnly: true,
// 	}
// 	c2 := http.Cookie{
// 		Name:     "second_cookie",
// 		Value:    "Manning",
// 		HttpOnly: true,
// 	}

// 	// w.Header().Set("Set-Cookie", c1.String())
// 	// w.Header().Add("Set-Cookie", c2.String())

// 	http.SetCookie(w, &c1)
// 	http.SetCookie(w, &c2)
// }

// func getCookie(w http.ResponseWriter, r *http.Request) {
// 	c1, err := r.Cookie("first_cookie")
// 	if err != nil {
// 		fmt.Fprintln(w, "Cannot get")
// 	}
// 	cs := r.Cookies()
// 	fmt.Fprintln(w, c1)
// 	fmt.Fprintln(w, cs)
// }

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "メッセージがありません")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
