package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DTan13/template/domain"
	"github.com/DTan13/template/server/handlers"
	"github.com/DTan13/template/service"
)

func Start() {
	domainHandler := handlers.DomainHandler{Service: service.NewTemplateDomain(domain.NewTemplateDomainStud())}

	http.HandleFunc("/domains", domainHandler.GetAllDomains)
	port := os.Getenv("PORT")
	fmt.Println("Starting Server at port", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("localhost:%v", port), nil))
}
