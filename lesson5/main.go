package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//Addr: "127.0.0.1:8080",
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}
	http.HandleFunc("/process", process)
	//http.HandleFunc("/form", form)
	server.ListenAndServe()
}

// func process(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("tmpl.html")
// 	// t, _ := template.ParseGlob("*.html")
// 	t.Execute(w, "Hello World!")

// 	// t, _ := template.ParseFiles("t1.html","t2.html")
// 	// t.ExecuteTemplate(w,"t2.html","Hello World!")
// }

// func formatDate(t time.Time) string {
// 	layout := "2006-01-02"
// 	return t.Format(layout)
// }

func process(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("tmpl.html")
	// rand.Seed(time.Now().Unix())
	// t.Execute(w, rand.Intn(10) > 5)

	// daysOfWeek := []string{"a", "b", "c", "d", "e"}
	// t.Execute(w, daysOfWeek)

	// t, _ := template.ParseFiles("t1.html", "t2.html")

	// t.Execute(w, "Hello")

	// funcMap := template.FuncMap{"fdate": formatDate}
	// t := template.New("tmpl.html").Funcs(funcMap)
	// t, _ = t.ParseFiles("tmpl.html")
	// t.Execute(w, time.Now())

	// t, _ := template.ParseFiles("tmpl.html")
	// content := `I asked: <i>"What's up"<i>?`
	// t.Execute(w, content)

	// t, _ := template.ParseFiles("tmpl.html")
	// // t.Execute(w, r.FormValue("comment"))
	// t.Execute(w, template.HTML(r.FormValue("comment")))

	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		// t, _ = template.ParseFiles("layout.html", "blue_hello.html")
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

// func form(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("attack.html")
// 	t.Execute(w, nil)
// }
