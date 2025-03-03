# Go Server

This is a project where I try to outline most boilerplate HTTP server code in golang. Since I always seem to forget how to encode JSON, setup resrouce handers and other basic HTTP serverside logic I keep some examples in this repository. The goal is to utilize golangs standard library (in this case net/http)

## Setting up handlers

Exmaple of a handler method that relays HTTP headers and JSON encoded payloads

```golang
import (
	"net/http"
)

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

func main() {
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/files/", fileHandler)
	slog.Info("Starting on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```
