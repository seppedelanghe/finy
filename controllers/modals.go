package controllers

import (
	"net/http"

	"finy.be/api/rendering"
)


func GetEmptyModal(w http.ResponseWriter, r *http.Request) {
	modalId := r.PathValue("id")
	rendering.Renderer.HTML(w, http.StatusOK, "modals/empty", H { "Id": modalId})
}

