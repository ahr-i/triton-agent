package handler

import "net/http"

/* Ping Handler: Ping Check 용도 */
func (h *Handler) pingHandler(w http.ResponseWriter, r *http.Request) {
	rend.JSON(w, http.StatusBadRequest, nil)
}
