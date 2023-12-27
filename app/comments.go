package app

import (
	"net/http"
	"strings"
)

func (a *App) ShowComments(w http.ResponseWriter, r *http.Request) error {
	return a.ServeAPI(w, r, map[string]any{
		"comments": a.comments,
	})
}

func (a *App) AddComment(w http.ResponseWriter, r *http.Request) error {
	comment := strings.Clone(r.FormValue("comment"))
	a.comments = append(a.comments, comment)

	return a.ShowComments(w, r)
}
