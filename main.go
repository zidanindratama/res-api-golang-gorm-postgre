package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// Article is a representation of a Article
type Article struct {
	ID    	int             `form:"id" json:"id";gorm:"primaryKey;autoIncrement"`
	Code  	string          `form:"code" json:"code"`
	Title  	string          `form:"title" json:"title"`
	Desc  	string          `form:"desc" json:"desc"`
	Content string 			`form:"content" json:"content"`
}

// Result is an array of Article
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}


var db *gorm.DB
var err error


func dbConnect() *gorm.DB {
	dsn := "host=localhost user=account password=KY8@GVL5\\q4e=S$R dbname=demo1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Auto Migrate
	_db.AutoMigrate(&Article{})

	if err != nil {
		log.Fatal(err)
	}

	return _db
}

func main() {
	db = dbConnect()

	handleRequest()
}

func handleRequest() {
	log.Println("Start the development server at http://127.0.0.1:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/articles", createArticle).Methods("POST")
	myRouter.HandleFunc("/api/articles", getArticles).Methods("GET")
	myRouter.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/api/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/api/articles/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(payloads, &article)

	db.Create(&article)

	res := Result{Code: 200, Data: article, Message: "Success create article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{}

	db.Find(&articles)

	res := Result{Code: 200, Data: articles, Message: "Success get articles"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article Article
	db.First(&article, articleID)

	res := Result{Code: 200, Data: article, Message: "Success get article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var articleUpdates Article
	json.Unmarshal(payloads, &articleUpdates)

	var article Article
	db.First(&article, articleID)
	db.Model(&article).Updates(articleUpdates)

	res := Result{Code: 200, Data: article, Message: "Success update article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article Article

	db.First(&article, articleID)
	db.Delete(&article)

	res := Result{Code: 200, Message: "Success delete article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}