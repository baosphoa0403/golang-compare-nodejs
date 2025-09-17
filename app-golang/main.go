package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func meHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello gia bao"))
}

// Tạo một Goroutine để thực hiện hashing password
func hashPasswordAsync(wg *sync.WaitGroup, password string, ch chan<- string) {
	defer wg.Done()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ch <- "Error hashing password"
		return
	}
	ch <- string(hashedPassword)
}

func hashPasswordHandler(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	password := "thisIsASecurePassword123" // Password để hash
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go hashPasswordAsync(&wg, password, ch)

	// Đợi goroutine hoàn thành

	hashedPassword := <-ch
	duration := time.Since(start)

	w.Header().Set("Content-Type", "application/json")
	rs := fmt.Sprintf(`{"result": "Hash completed", "hashedPassword": "%s", "duration": "%s"}`, hashedPassword, duration)
	fmt.Println(rs)
	w.Write([]byte(rs))

	wg.Wait()

}

func main() {
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/hash-password", hashPasswordHandler)

	fmt.Println("Go server is running at http://localhost:8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
