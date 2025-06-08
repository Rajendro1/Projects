package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/proto"

	pb "go_react/backend/proto"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

var db *sql.DB

func main() {
	var err error
	// db, err = sql.Open("mysql", "user:pass@tcp(localhost:3306)/testdb")
	dsn := os.Getenv("DB_DSN")
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/rest", jsonHandler)
	http.HandleFunc("/resp", respHandler)
	http.HandleFunc("/proto", protoHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func getUsers() []User {
	rows, _ := db.Query("SELECT id, name, email, age FROM users LIMIT 1000")
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		users = append(users, u)
	}
	return users
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	users := getUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func respHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	users := getUsers()
	var b strings.Builder
	for _, u := range users {
		b.WriteString(fmt.Sprintf("*4\r\n$%d\r\n%d\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n$%d\r\n%d\r\n",
			len(strconv.Itoa(u.ID)), u.ID,
			len(u.Name), u.Name,
			len(u.Email), u.Email,
			len(strconv.Itoa(u.Age)), u.Age))
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(b.String()))
}

func protoHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	users := getUsers()
	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id: int32(u.ID), Name: u.Name, Email: u.Email, Age: int32(u.Age),
		})
	}
	data, _ := proto.Marshal(&pb.UserList{Users: pbUsers})
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // You can restrict this to "http://localhost:3000"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
