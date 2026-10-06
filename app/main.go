// A placeholder web app for a Krateo event-site blueprint.
//
// It serves one page at "/" and nothing else. That is deliberate: it has no
// health endpoint, so a liveness probe pointed at /healthz gets a 404 and the
// pod never becomes Ready — which is the condition chapter 5 diagnoses and
// fixes by correcting the blueprint's probePath.
package main

import (
	"fmt"
	"net/http"
	"os"
)

const page = `<!doctype html><html lang="en"><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1"><title>%s</title>
<style>:root{color-scheme:light dark}body{margin:0;min-height:100vh;display:grid;place-items:center;
font:16px/1.5 ui-sans-serif,system-ui,sans-serif;background:#0b1620;color:#e8eef4}
.c{text-align:center;padding:32px}h1{margin:0 0 8px;font-size:28px;letter-spacing:-.02em}
p{margin:0;color:#8fa6b8}b{color:#11b2e2}</style></head><body><div class="c">
<h1>%s</h1><p>Served by <b>%s</b> — a Krateo composition.</p></div></body></html>`

func main() {
	site := os.Getenv("SITE_NAME")
	if site == "" {
		site = "Event site"
	}
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r) // no /healthz here, on purpose
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, page, site, site, "krateo")
	})
	fmt.Fprintf(os.Stderr, "listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
