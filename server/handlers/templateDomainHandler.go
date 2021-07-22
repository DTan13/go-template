package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DTan13/template/service"
)

type DomainHandler struct {
	Service service.TemplateDomainService
}

func (h *DomainHandler) GetAllDomains(w http.ResponseWriter, r *http.Request) {
	domains, err := h.Service.GetAllDomains()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
	}
	writeResponse(w, http.StatusOK, domains)
}

func writeResponse(w http.ResponseWriter, code int, data ...interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data[0])
	if err != nil {
		_, _ = fmt.Fprintln(w, "error")
		return
	}
}
