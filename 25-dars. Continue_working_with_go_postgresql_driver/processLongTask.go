package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	userId = 0
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")
	ctx := context.WithValue(r.Context(), userId, id)
	result := ProcessLongTask(ctx)
	_, err := w.Write([]byte(result))
	if err != nil {
		fmt.Println("write error:", err)
	}
}

func ProcessLongTask(ctx context.Context) string {
	id := ctx.Value(userId)
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("time out %s", id)
	case <-ctx.Done():
		return fmt.Sprintf("DOne")

	}
}
