package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	envVarFastlyServiceIDs = "FASTLY_SERVICE_IDS"
	envVarPort             = "PORT"
	envVarProxyURL         = "PROXY_URL"
	fastlyChallengeURLPath = "/.well-known/fastly/logging/challenge"
)

func main() {
	port := strings.TrimSpace(os.Getenv(envVarPort))
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc(fastlyChallengeURLPath, func(w http.ResponseWriter, req *http.Request) {
		if _, err := io.WriteString(w, challengeResponseBody(os.Getenv(envVarFastlyServiceIDs))); err != nil {
			httpError(w, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if _, err := http.Post(os.Getenv(envVarProxyURL), req.Header.Get("Content-Type"), req.Body); err != nil {
			httpError(w, http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func challengeResponseBody(serviceIDs ...string) string {
	responses := []string{}
	for _, idStr := range serviceIDs {
		ids := strings.Split(idStr, ",")
		for _, id := range ids {
			id = strings.TrimSpace(id)
			if len(id) == 0 {
				continue
			} else if id == "*" {
				responses = append(responses, "*")
			} else {
				responses = append(responses, sum256String(id))
			}
		}
	}
	if len(responses) == 0 {
		responses = append(responses, "*")
	}
	return strings.Join(responses, "\n")
}

func sum256String(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash[:])
}

func httpError(w http.ResponseWriter, httpStatus int) {
	http.Error(w, http.StatusText(httpStatus), httpStatus)
}
