package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
)

type Response map[string]interface{}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// ========== CASE 0: Hello ==========
func meHandler(w http.ResponseWriter, r *http.Request) {
	// Giả lập async như setTimeout 0 trong Node
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
	duration := time.Since(start)

	writeJSON(w, Response{
		"result":         "Hash completed (sync)",
		"hashedPassword": string(hashed),
		"duration":       duration.String(),
	})
}

// ========== CASE 2: Hash password with goroutine (simulate worker pool) ==========
func hashPasswordWorkerHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	password := "thisIsASecurePassword123"
	ch := make(chan string, 1)

	go func() {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		ch <- string(hashed)
	}()

	hashed := <-ch
	duration := time.Since(start)

	writeJSON(w, Response{
		"result":         "Hash completed (goroutine worker)",
		"hashedPassword": hashed,
		"duration":       duration.String(),
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
		"duration": time.Since(start).String(),
	})
}

// ========== CASE 4: Read medium Excel ==========
func excelMediumHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	f, _ := excelize.OpenFile("./data/medium.xlsx")
	rows, _ := f.GetRows("Sheet1")

	writeJSON(w, Response{
		"case":     "Medium Excel",
		"duration": time.Since(start).String(),
		"rows":     len(rows),
	})
}

// ========== CASE 5: Stream large Excel ==========
func excelLargeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	f, _ := excelize.OpenFile("./data/large.xlsx")
	rows, _ := f.Rows("Sheet1")

	count := 0
	for rows.Next() {
		count++
	}

	writeJSON(w, Response{
		"case":     "Large Excel Streaming",
		"duration": time.Since(start).String(),
		"rows":     count,
	})
}

// ========== CASE 6: Validate + Insert DB ==========
func excelValidateInsertHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost:5432/benchmark?sslmode=disable")
	defer db.Close()

	f, _ := excelize.OpenFile("./data/import.xlsx")
	rows, _ := f.Rows("Sheet1")

	inserted := 0
	for rows.Next() {
		row, _ := rows.Columns()
		if len(row) < 2 {
			continue
		}
		name := row[0]
		email := row[1]
		if !strings.Contains(email, "@") {
			continue
		}
		_, _ = db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
		inserted++
	}

	writeJSON(w, Response{
		"case":     "Validate + Insert",
		"inserted": inserted,
		"duration": time.Since(start).String(),
	})
}

func main() {
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/hash-password-sync", hashPasswordSyncHandler)
	http.HandleFunc("/hash-password", hashPasswordWorkerHandler)
	http.HandleFunc("/excel-small", excelSmallHandler)
	http.HandleFunc("/excel-medium", excelMediumHandler)
	http.HandleFunc("/excel-large", excelLargeHandler)
	http.HandleFunc("/excel-validate-insert", excelValidateInsertHandler)

	fmt.Println("Go server is running at http://localhost:8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
