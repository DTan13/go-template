package handlers

import (
	"fmt"
	"os"

	"github.com/DTan13/template/service"
)

type DomainHandler struct {
	Service service.TemplateDomainService
}

func (h *DomainHandler) GetAllDomains() {
	domains, err := h.Service.GetAllDomains()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stdout, "unexpected error occurred")
		return
	}

	for _, d := range domains {
		fmt.Println(d.Name)
	}

	return
}
