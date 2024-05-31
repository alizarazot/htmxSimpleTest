package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
	<meta charset=utf-8>
	 <script src="https://unpkg.com/htmx.org@1.9.12"></script>
</head>
<body>
	<h1> Current hour is: <span data-hx-get="/hour" data-hx-trigger="every 1s"></span> </h1>
	<h1
		data-hx-get="/mouse-enter" 
		data-hx-swap="outerHTML"
		data-hx-trigger=mouseenter
		data-hx-on:htmx:config-request="event.detail.parameters.changed = 'false'"
		> Hello! </h1>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/mouse-enter", func(w http.ResponseWriter, r *http.Request) {
		text := "Hello!"
		ch := "false"

		if r.URL.Query()["changed"][0] == "false" {
			text = "Heey!"
			ch = "true"
		}

		fmt.Fprintf(w, `
		<h1
			data-hx-get="/mouse-enter" 
			data-hx-swap="outerHTML"
			data-hx-trigger=mouseenter
			data-hx-on:htmx:config-request="event.detail.parameters.changed = '%s'"
		> %s </h1>
		`, ch, text)
	})

	http.HandleFunc("/hour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, time.Now().Format("15:04:05"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexHTML)
}
