package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/yt-passwords-video/code/web/templates"
)

type pageController struct {
}

func PageController(db *sqlx.DB) pageController {
	return pageController{}
}

func (p *pageController) Index(w http.ResponseWriter, r *http.Request) {
	templates.Pages.Index.Execute(w, nil)
}
