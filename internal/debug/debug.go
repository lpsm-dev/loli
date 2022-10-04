package debug

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	log "github.com/ci-monk/loli/internal/log"
)

var (
	// Verbose determines if debugging output is displayed to the user
	Verbose bool
	output  io.Writer = os.Stderr
)

// Println conditionally outputs a message to Stderr
func Println(args ...interface{}) {
	if Verbose {
		fmt.Fprintln(output, args...)
	}
}

// DumpRequest dumps out the provided http.Request
func DumpRequest(req *http.Request) {
	if !Verbose {
		return
	}

	var bodyCopy bytes.Buffer
	body := io.TeeReader(req.Body, &bodyCopy)
	req.Body = io.NopCloser(body)

	dump, err := httputil.DumpRequest(req, req.ContentLength > 0)
	if err != nil {
		log.Error(err)
	}

	Println("\n========================= BEGIN DumpRequest =========================")
	Println(string(dump))
	Println("========================= END DumpRequest =========================")
	Println("")

	req.Body = io.NopCloser(&bodyCopy)
}

// DumpResponse dumps out the provided http.Response
func DumpResponse(res *http.Response) {
	if !Verbose {
		return
	}

	var bodyCopy bytes.Buffer
	body := io.TeeReader(res.Body, &bodyCopy)
	res.Body = io.NopCloser(body)

	dump, err := httputil.DumpResponse(res, res.ContentLength > 0)
	if err != nil {
		log.Error(err)
	}

	Println("\n========================= BEGIN DumpResponse =========================")
	Println(string(dump))
	Println("========================= END DumpResponse =========================")
	Println("")

	res.Body = io.NopCloser(body)
}
