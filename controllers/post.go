package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct{}

func (e *PostController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
