package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", headerHandle)
	http.HandleFunc("/healthz", healthzHandle)
	log.Println("Start HTTP Server 127.0.0.1:8088")
	log.Println("/")
	log.Println("/healthz")
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func headerHandle(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	fmt.Fprintf(w, "code: %d\n", code)
	headers := r.Header
	for header, value := range headers {
		fmt.Fprintf(w, "%s: %s\n", header, value)
	}
	version, err := getVersion()
	if err == nil && len(version) > 0 {
		fmt.Fprintf(w, "Version: %s\n", version)
	} else {
		fmt.Fprintf(w, "Version: %s\n", "")
	}
	log.Printf("[%d] %s %s\n", code, r.RequestURI, r.RemoteAddr)
}

func healthzHandle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, http.StatusOK)
}

func getVersion() (version string, err error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}