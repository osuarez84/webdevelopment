package main

import (
	"fmt"
	"net/http"
)

const faqPage = `
<h1>FAQ page</h1>
<p>Q: Is there a free version?</p>
<p>A: Yes! We offer a free trial for 30 days on any paid plans</p>
<p>Q: What are your support hours?</p>
<p>A: WE have support staff answering emails 24/7 though response times may be a bit slower on weekends</p>
`

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my big site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me "+
		"at <a href=\"mailto:omsulo@gmail.com\">omsulo@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, faqPage)
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//	}
//}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", router)
}
