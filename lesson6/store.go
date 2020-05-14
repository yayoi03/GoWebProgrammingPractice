package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

// type Post struct {
// 	Id       int
// 	Content  string
// 	Author   string
// 	Comments []Comment
// }

// type Comment struct {
// 	Id      int
// 	Content string
// 	Author  string
// 	Post    *Post
// }

// type Post struct {
// 	Id         int
// 	Content    string
// 	AuthorName string `db: author`
// }

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

// var PostById map[int]*Post
// var PostByAuthor map[string][]*Post

// func store(post Post) {
// 	PostById[post.Id] = &post
// 	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
// }

// func store(data interface{}, filename string) {
// 	buffer := new(bytes.Buffer)
// 	encoder := gob.NewEncoder(buffer)
// 	err := encoder.Encode(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func load(data interface{}, filename string) {
// 	raw, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	buffer := bytes.NewBuffer(raw)
// 	dec := gob.NewDecoder(buffer)
// 	err = dec.Decode(data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

//var Db *sql.DB
//var Db *sqlx.DB
var Db *gorm.DB

func init() {
	// var err error
	// Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=Svolme0906 sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }

	// var err error
	// Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=Svolme0906 sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }

	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=Svolme0906 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

// func Posts(limit int) (posts []Post, err error) { //???
// 	rows, err := Db.Query("select id,content,author from posts limit $1 ", limit)
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		post := Post{}
// 		err = rows.Scan(&post.Id, &post.Content, &post.Author)
// 		if err != nil {
// 			return
// 		}
// 		posts = append(posts, post)
// 	}
// 	rows.Close()
// 	return
// }

// func GetPost(id int) (post Post, err error) {
// 	post = Post{}
// 	err = Db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
// 	return
// }

// func (post *Post) Create() (err error) {
// 	statement := "insert into posts (content, author) values ($1,$2) returning id"
// 	stmt, err := Db.Prepare(statement)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()
// 	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
// 	return
// }

// func (post *Post) Update() (err error) {
// 	_, err = Db.Exec("update posts set content=$2,author=$3 where id=$1", post.Id, post.Content, post.Author)
// 	return
// }

// func (post *Post) Delete() (err error) {
// 	_, err = Db.Exec("delete from posts where id=$1", post.Id)
// 	return
// }

// func GetPost(id int) (post Post, err error) {
// 	post = Post{}
// 	post.Comments = []Comment{}
// 	err = Db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)

// 	rows, err := Db.Query("select id,content,author from comments where post_id=$1", id)
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		comment := Comment{Post: &post}
// 		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
// 		if err != nil {

// 			return
// 		}
// 	}
// 	rows.Close()

// 	return
// }

// func (comment *Comment) Create() (err error) {
// 	if comment.Post == nil {
// 		err = errors.New("投稿が見つかりません")
// 		return
// 	}
// 	err = Db.QueryRow("insert into comments (content,author,post_id) values ($1, $2,$3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
// 	return
// }

// func (post *Post) Create() (err error) {
// 	err = Db.QueryRow("insert into posts (content,author) values ($1,$2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
// 	return
// }

// func GetPost(id int) (post Post, err error) {
// 	post = Post{}
// 	err = Db.QueryRowx("select id,content,author from posts where id=$1", id).StructScan(&post)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func main() {
	// PostById = make(map[int]*Post)
	// PostByAuthor = make(map[string][]*Post)

	// post1 := Post{Id: 1, Content: "Hello1", Author: "Sau"}
	// post2 := Post{Id: 2, Content: "Hello2", Author: "pa"}
	// post3 := Post{Id: 3, Content: "Hello3", Author: "pi"}
	// post4 := Post{Id: 4, Content: "Hello4", Author: "Sau"}

	// store(post1)
	// store(post2)
	// store(post3)
	// store(post4)

	// fmt.Println(PostById[1])
	// fmt.Println(PostById[2])

	// for _, post := range PostByAuthor["Sau"] {
	// 	fmt.Println(post)
	// }

	// for _, post := range PostByAuthor["pi"] {
	// 	fmt.Println(post)
	// }

	// data := []byte("Hello World\n")
	// err := ioutil.WriteFile("data1", data, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// read1, _ := ioutil.ReadFile("data1")
	// fmt.Print(string(read1))

	// file1, _ := os.Create("data2")
	// defer file1.Close()

	// bytes, _ := file1.Write(data)
	// fmt.Printf("Wrote %d bytes to file \n", bytes)

	// file2, _ := os.Open("data2")
	// defer file2.Close()

	// read2 := make([]byte, len(data))
	// bytes, _ = file2.Read(read2)
	// fmt.Printf("Read %d bytes from file\n", bytes)
	// fmt.Println(string(read2))

	// csvFile, err := os.Create("posts.csv")
	// if err != nil {
	// 	panic(nil)
	// }
	// defer csvFile.Close()

	// allPosts := []Post{
	// 	Post{Id: 1, Content: "Hello1", Author: "Sau"},
	// 	Post{Id: 2, Content: "Hello2", Author: "pa"},
	// 	Post{Id: 3, Content: "Hello3", Author: "pi"},
	// 	Post{Id: 4, Content: "Hello4", Author: "Sau"},
	// }

	// writer := csv.NewWriter(csvFile)

	// for _, post := range allPosts {
	// 	line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
	// 	err := writer.Write(line)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// writer.Flush()

	// file, err := os.Open("posts.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1
	// record, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// var posts []Post
	// for _, item := range record {
	// 	id, _ := strconv.ParseInt(item[0], 0, 0)
	// 	post := Post{Id: int(id), Content: item[1], Author: item[2]}
	// 	posts = append(posts, post)
	// }

	// for _, c := range posts {
	// 	fmt.Println(c.Id)
	// 	fmt.Println(c.Content)
	// 	fmt.Println(c.Author)
	// }

	// post := Post{Id: 1, Content: "Hello1", Author: "Sau"}
	// store(post, "post1")
	// var postRead Post
	// load(&postRead, "post1")
	// fmt.Println(postRead)

	// post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	// fmt.Println(post)
	// post.Create()
	// fmt.Println(post)

	// readPost, _ := GetPost(post.Id)
	// fmt.Println(readPost)

	// readPost.Content = "Bonjour Monde!"
	// readPost.Author = "Pierre"
	// readPost.Update()

	// posts, _ := Posts(5)
	// //fmt.Println(posts[1])
	// for _, c := range posts {
	// 	fmt.Println(c)
	// }
	// //fmt.Println(len(posts))

	// readPost.Delete()

	// post := Post{Content: "Hello World", AuthorName: "Sau Sheong"}
	// post.Create()

	// fmt.Println(post)
	// readPost := Post{}
	// readPost, _ = GetPost(post.Id)
	// fmt.Println(readPost)

	// comment := Comment{Content: "いい投稿じゃん", Author: "Joe", Post: &post}
	// comment.Create()
	//readPost, _ := GetPost(post.Id)

	// fmt.Println(readPost)
	// fmt.Println(readPost.Comments)
	//fmt.Println(readPost.Comments[0].Post)

	post := Post{Content: "Hello World", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "いい投稿じゃん", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author=$1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])

}
