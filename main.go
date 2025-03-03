package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func readFile(filePath string) (error, string) {
	f := "files/" + filePath
	data, err := os.ReadFile(f)
	if err != nil {
		return err, ""
	}
	return nil, string(data)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		errMsg := fmt.Sprintf("Method not allowed %s need %s", r.Method, http.MethodGet)
		response := map[string]string{
			"message": errMsg,
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		slog.Error(
			"Method not allowed",
			"method", r.Method,
			"url", r.URL.String(),
			"client", r.RemoteAddr,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "OK",
	}
	json.NewEncoder(w).Encode(response)
	slog.Info(
		"Handled request",
		"method", r.Method,
		"url", r.URL.String(),
		"client", r.RemoteAddr,
	)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		errMsg := fmt.Sprintf("Method not allowed %s need %s", r.Method, http.MethodGet)
		response := map[string]string{
			"message": errMsg,
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		slog.Error(
			"Method not allowed",
			"method", r.Method,
			"url", r.URL.String(),
			"client", r.RemoteAddr,
		)
		return
	}

	resrouce := strings.TrimPrefix(r.URL.Path, "/files/")
	err, payload := readFile(resrouce)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to fetch resrouce", http.StatusInternalServerError)
		slog.Error(
			"Failed to fetch resource",
			"method", r.Method,
			"url", r.URL.String(),
			"client", r.RemoteAddr,
			"error", err.Error(),
			"resource", resrouce,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": payload,
	}
	json.NewEncoder(w).Encode(response)
	slog.Info(
		"Handled request",
		"method", r.Method,
		"url", r.URL.String(),
		"client", r.RemoteAddr,
		"resource", resrouce,
	)
}

func main() {
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/files/", fileHandler)
	slog.Info("Starting on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
