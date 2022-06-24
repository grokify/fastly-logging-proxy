package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const fastlyChallengeURLPath = "/.well-known/fastly/logging/challenge"

func main() {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if len(port) == 0 {
		port = "8080"
	}

	fastlyChallengeHandler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, challengeResponseBody())
		if err != nil {
			httpError(w, http.StatusInternalServerError)
		}
	}

	fastlyLogHandler := func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			httpError(w, http.StatusInternalServerError)
		}

		_, err = http.Post(os.Getenv("PROXY_URL"), req.Header.Get("Content-Type"), bytes.NewBuffer(body))
		if err != nil {
			httpError(w, http.StatusInternalServerError)
		}
	}

	http.HandleFunc(fastlyChallengeURLPath, fastlyChallengeHandler)
	http.HandleFunc("/", fastlyLogHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func challengeResponseBody() string {
	ids := strings.Split(os.Getenv("FASTLY_SERVICE_IDS"), ",")
	responses := []string{}
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
