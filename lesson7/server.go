package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=Svolme0906 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// func retrieve(id int) (post Post, err error) {
// 	post = Post{}
// 	err = Db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
// 	return
// }

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content=$2,author=$3 where id=$1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id=$1", post.Id)
	return
}

// type Post struct {
// 	Id      int    `json:"id"`
// 	Content string `json:"content"`
// 	Author  string `json:"author"`
// }

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	switch r.Method {
// 	case "GET":
// 		err = handleGet(w, r)
// 	case "POST":
// 		err = handlePost(w, r)
// 	case "PUT":
// 		err = handlePut(w, r)
// 	case "DELETE":
// 		err = handleDelete(w, r)
// 	}
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
// 	id, err := strconv.Atoi(path.Base(r.URL.Path))
// 	if err != nil {
// 		return
// 	}
// 	post, err := retrieve(id)
// 	if err != nil {
// 		return
// 	}
// 	output, err := json.MarshalIndent(&post, "", "\t\t")
// 	if err != nil {
// 		return
// 	}
// 	w.Header().Set("Content-Type", "applicatioon/json")
// 	w.Write(output)
// 	return
// }

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "applicatioon/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
