package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var trie *TrieNode
var db *sql.DB

type PostRow struct {
	PostId       int64
	PosterId     string
	CommId       string
	ParentPostId string
	TextContent  string
	MediaLinks   string
	EventId      string
	PostDate     string
}
type UserRow struct {
	Username   string
	PosterId   string
	MediaLinks string
}
type ResultRows struct {
	PostRows []PostRow
	UserRows []UserRow
}

func search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	word := r.URL.Query().Get("search")
	addQuery(word, trie)
	rows, err := db.Query(`SELECT * FROM forum WHERE textContent LIKE '%' || $1 || '%'`, word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	var rowsData []PostRow
	for rows.Next() {
		var (
			PostId       int64
			PosterId     string
			CommId       string
			ParentPostId string
			MediaLinks   string
			TextContent  string
			EventId      string
			PostDate     string
		)
		if err := rows.Scan(&PostId, &PosterId, &PostDate, &CommId, &ParentPostId, &TextContent, &MediaLinks, &EventId); err != nil {
			log.Fatal(err)
		}
		rowsData = append(rowsData, PostRow{PostId: PostId, PosterId: PosterId, PostDate: PostDate, CommId: CommId, ParentPostId: ParentPostId, TextContent: TextContent, MediaLinks: MediaLinks, EventId: EventId})
	}
	rows, err = db.Query(`SELECT * FROM users WHERE username LIKE '%' || $1 || '%'`, word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	var userData []UserRow
	for rows.Next() {
		var (
			PosterId   string
			JoinDate   string
			Username   string
			Password   string
			Email      string
			MediaLinks string
		)
		if err := rows.Scan(&PosterId, &JoinDate, &Username, &Password, &Email, &MediaLinks); err != nil {
			log.Fatal(err)
		}
		userData = append(userData, UserRow{PosterId: PosterId, Username: Username, MediaLinks: MediaLinks})
	}
	structResult := ResultRows{PostRows: rowsData, UserRows: userData}
	result, error := json.Marshal(structResult)
	if error != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(result)
}
func auto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	word := r.URL.Query().Get("search")
	result, err := json.Marshal(getTop(word, trie))
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	w.Write(result)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/auto", auto)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":1026", nil)
}
func main() {
	addQuery("1up", trie)
	addQuery("1upagain", trie)
	addQuery("b@sman", trie)
	addQuery("battery", trie)
	addQuery("batting", trie)
	addQuery("batmobile", trie)
	addQuery("bats", trie)
	addQuery("batch", trie)
	addQuery("battle", trie)
	addQuery("baton", trie)
	addQuery("bathtub", trie)
	addQuery("batik", trie)
	addQuery("batter", trie)
	addQuery("batwoman", trie)
	addQuery("batty", trie)
	addQuery("bathe", trie)
	addQuery("baton", trie)
	addQuery("batsman", trie)
	addQuery("battalion", trie)
	addQuery("batfish", trie)
	addQuery("batmobile", trie)
	addQuery("batik", trie)
	addQuery("batter", trie)
	addQuery("batwoman", trie)
	addQuery("batty", trie)
	addQuery("bathe", trie)
	addQuery("baton", trie)
	addQuery("batsman", trie)
	addQuery("battalion", trie)
	addQuery("batfish", trie)
	addQuery("batmobile", trie)
	addQuery("batik", trie)
	addQuery("batter", trie)
	addQuery("batwoman", trie)
	addQuery("batty", trie)
	addQuery("bathe", trie)
	addQuery("baton", trie)
	addQuery("batsman", trie)
	addQuery("battalion", trie)
	addQuery("batfish", trie)
	addQuery("batmobile", trie)
	addQuery("batik", trie)
	addQuery("batter", trie)
	addQuery("batwoman", trie)
	addQuery("batty", trie)
	addQuery("bathe", trie)
	addQuery("baton", trie)
	addQuery("batsman", trie)
	addQuery("battalion", trie)
	addQuery("batfish", trie)
	addQuery("bat", trie)
	addQuery("battery", trie)
	addQuery("batman", trie)
	addQuery("batting", trie)
	addQuery("batter", trie)
	addQuery("batmobile", trie)
	addQuery("1", trie)
	handleRequests()
}
func init() {
	trie = &TrieNode{
		count:   0,
		nodes:   make(map[rune]*TrieNode),
		wordEnd: false,
	}
	godotenv.Load()
	host, hostError := os.LookupEnv("DB_HOST")
	if !hostError {
		panic("Couldn't get DB host")
	}
	user, userError := os.LookupEnv("DB_USER")
	if !userError {
		panic("Couldn't get DB username")
	}
	password, passwordError := os.LookupEnv("DB_PASSWORD")
	if !passwordError {
		panic("Couldn't get DB password")
	}
	port, portError := strconv.Atoi(os.Getenv("DB_PORT"))
	if portError != nil {
		panic(portError)
	}
	name, nameError := os.LookupEnv("DB_NAME")
	if !nameError {
		panic("Couldn't get DB name")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
