// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// // type Post struct {
// // 	XMLName  xml.Name  `xml:"post"`
// // 	Id       string    `xml:"id,attr"`
// // 	Content  string    `xml:"content"`
// // 	Author   Author    `xml:"author"`
// // 	Xml      string    `xml:",innerxml"`
// // 	Comments []Comment `xml:"comments>comment"`
// // }

// // type Comment struct {
// // 	Id      string `xml:"id,attr"`
// // 	Content string `xml:"content"`
// // 	Author  Author `xml:"author"`
// // }

// // type Author struct {
// // 	Id   string `xml:"id,attr"`
// // 	Name string `xml:",chardata"`
// // }

// // func main() {
// // 	xmlFile, err := os.Open("post.xml")
// // 	if err != nil {
// // 		fmt.Println("Error opening XML file:", err)
// // 		return
// // 	}
// // 	defer xmlFile.Close()

// // 	decoder := xml.NewDecoder(xmlFile)
// // 	for {
// // 		t, err := decoder.Token()
// // 		if err == io.EOF {
// // 			break
// // 		}
// // 		if err != nil {
// // 			fmt.Println("Error decoding XML into tokens:", err)
// // 			return
// // 		}

// // 		switch se := t.(type) {
// // 		case xml.StartElement:
// // 			if se.Name.Local == "comment" {
// // 				var comment xml.Comment
// // 				decoder.DecodeElement(&comment, &se)
// // 				//fmt.Println(comment)
// // 			}
// // 		}
// // 	}

// // 	// xmlData, err := ioutil.ReadAll(xmlFile)
// // 	// if err != nil {
// // 	// 	fmt.Println("Error reading XML data:", err)
// // 	// 	return
// // 	// }

// // 	// var post Post
// // 	// xml.Unmarshal(xmlData, &post)
// // 	// fmt.Println(post)
// // }

// // type Post struct {
// // 	XMLName xml.Name `xml:"post"`
// // 	Id      string   `xml:"id,attr"`
// // 	Content string   `xml:"content"`
// // 	Author  Author   `xml:"author"`
// // }

// // type Author struct {
// // 	Id   string `xml:"id,attr"`
// // 	Name string `xml:",chardata"`
// // }

// // func main() {
// // 	post := Post{
// // 		Id:      "1",
// // 		Content: "Hello World!!",
// // 		Author: Author{
// // 			Id:   "2",
// // 			Name: "Sau Sheong",
// // 		},
// // 	}

// // 	xmlFile, err := os.Create("post.xml")
// // 	if err != nil {
// // 		fmt.Println("Error creating XML file:", err)
// // 		return
// // 	}
// // 	encoder := xml.NewEncoder(xmlFile)
// // 	encoder.Indent("", "\t")
// // 	err = encoder.Encode(&post)
// // 	if err != nil {
// // 		fmt.Println("Error encoding XML to file:", err)
// // 		return
// // 	}

// // 	// output, err := xml.Marshal(&post)
// // 	// if err != nil {
// // 	// 	fmt.Println("Error marshalling to XML:", err)
// // 	// 	return
// // 	// }
// // 	// err = ioutil.WriteFile("post.xml", output, 0644)
// // 	// if err != nil {
// // 	// 	fmt.Println("Error writing XML to file:", err)
// // 	// 	return
// // 	// }

// // }

// type Post struct {
// 	Id       int       `json:"id"`
// 	Content  string    `json:"content"`
// 	Author   Author    `json:"author"`
// 	Comments []Comment `json:"comments"`
// }

// type Comment struct {
// 	Id      int    `json:"id"`
// 	Content string `json:"content"`
// 	Author  string `json:"author"`
// }

// type Author struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func main() {
// 	// jsonFile, err := os.Open("post.json")
// 	// if err != nil {
// 	// 	fmt.Println("Error opening JSON file:", err)
// 	// 	return
// 	// }
// 	// defer jsonFile.Close()

// 	// decoder := json.NewDecoder(jsonFile)
// 	// for {
// 	// 	var post Post
// 	// 	err := decoder.Decode(&post)
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println("Error decoding JSON:", err)
// 	// 		return
// 	// 	}
// 	// 	fmt.Println(post)
// 	// }

// 	// jsonData, err := ioutil.ReadAll(jsonFile)
// 	// if err != nil {
// 	// 	fmt.Println("Error reading json data:", err)
// 	// 	return
// 	// }

// 	// var post Post
// 	// json.Unmarshal(jsonData, &post)
// 	// fmt.Println(post)

// 	post := Post{
// 		Id:      1,
// 		Content: "Hello World!!!!",
// 		Author: Author{
// 			Id:   2,
// 			Name: "Sau Sheong",
// 		},
// 		Comments: []Comment{
// 			Comment{
// 				Id:      3,
// 				Content: "Hey",
// 				Author:  "Sau",
// 			},
// 			Comment{
// 				Id:      4,
// 				Content: "Hey2",
// 				Author:  "Sau2",
// 			},
// 		},
// 	}

// 	// output, err := json.MarshalIndent(&post, "", "\t\t")
// 	// if err != nil {
// 	// 	fmt.Println("Error marshalling to JSON:", err)
// 	// 	return
// 	// }
// 	// err = ioutil.WriteFile("post.json", output, 0644)
// 	// if err != nil {
// 	// 	fmt.Println("Error writing JSON to file:", err)
// 	// 	return
// 	// }

// 	jsonFile, err := os.Create("post.json")
// 	if err != nil {
// 		fmt.Println("Error creating JSON file:", err)
// 		return
// 	}
// 	encoder := json.NewEncoder(jsonFile)
// 	err = encoder.Encode(&post)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON to file:", err)
// 		return
// 	}

// }
