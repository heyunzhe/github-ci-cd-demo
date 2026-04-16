package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Song struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

var songs = []Song{
	{1, "稻香", "周杰伦"},
	{2, "演员", "薛之谦"},
	{3, "晴天", "周杰伦"},
	{4, "夜曲", "周杰伦"},
}

var current Song

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/song", getSong)
	http.HandleFunc("/guess", guessSong)
	http.HandleFunc("/health", healthCheck)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 获取题目
func getSong(w http.ResponseWriter, r *http.Request) {
	current = songs[rand.Intn(len(songs))]

	resp := map[string]interface{}{
		"id": current.ID,
		// 不返回答案！
	}

	json.NewEncoder(w).Encode(resp)
}

// 提交答案
func guessSong(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	answer := req["answer"]

	correct := answer == current.Name

	resp := map[string]interface{}{
		"correct": correct,
		"answer":  current.Name,
	}

	json.NewEncoder(w).Encode(resp)
}

// 健康检查（运维重点）
func healthCheck(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(resp)
}
