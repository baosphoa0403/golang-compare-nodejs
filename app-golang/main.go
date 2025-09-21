package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
)

type Response map[string]interface{}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// Helper: convert duration → giây có 3 số thập phân
func durationInSeconds(start time.Time) string {
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.3fs", secs)
}

// ========== CASE 0: Hello ==========
func meHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		fmt.Println("set timeout 0s")
	}()
	writeJSON(w, Response{"result": "hello gia bao 123"})
}

// ========== CASE 1: Hash password sync ==========
func hashPasswordSyncHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	password := "thisIsASecurePassword123"

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	writeJSON(w, Response{
		"result":         "Hash completed (sync)",
		"hashedPassword": string(hashed),
		"duration":       durationInSeconds(start),
	})
}

// ========== CASE 2: Hash password with goroutine ==========
func hashPasswordWorkerHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	password := "thisIsASecurePassword123"
	ch := make(chan string, 1)

	go func() {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		ch <- string(hashed)
	}()

	hashed := <-ch

	writeJSON(w, Response{
		"result":         "Hash completed (goroutine worker)",
		"hashedPassword": hashed,
		"duration":       durationInSeconds(start),
	})
}

// ========== CASE 3: Read small Excel ==========
func excelSmallHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	f, _ := excelize.OpenFile("./data/small.xlsx")
	rows, _ := f.GetRows("Sheet1")

	writeJSON(w, Response{
		"case":     "Small Excel",
		"rows":     len(rows),
		"duration": durationInSeconds(start),
	})
}

// ========== CASE 4: Read medium Excel ==========
func excelMediumHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	f, _ := excelize.OpenFile("./data/medium.xlsx")
	rows, _ := f.GetRows("Sheet1")

	writeJSON(w, Response{
		"case":     "Medium Excel",
		"rows":     len(rows),
		"duration": durationInSeconds(start),
	})
}

// ========== CASE 5: Stream large Excel ==========
func excelLargeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	f, err := excelize.OpenFile("./data/large.xlsx")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	rows, err := f.Rows("Sheet1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
	}

	writeJSON(w, Response{
		"case":     "Large Excel Streaming",
		"rows":     count,
		"duration": durationInSeconds(start),
	})
}

func main() {
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/hash-password-sync", hashPasswordSyncHandler)
	http.HandleFunc("/hash-password", hashPasswordWorkerHandler)
	http.HandleFunc("/excel-small", excelSmallHandler)
	http.HandleFunc("/excel-medium", excelMediumHandler)
	http.HandleFunc("/excel-large", excelLargeHandler)

	fmt.Println("Go server is running at http://localhost:8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
