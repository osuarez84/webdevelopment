package controllers

import (
	"github.com/osuarez84/webdevelopment/views"
	"html/template"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7.",
		},
		{
			Question: "How to I contact support?",
			Answer:   `Email us - <a href="mailto:lenslocked.com">support@lenslocked.com`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
