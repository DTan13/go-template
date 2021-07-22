package cmd

import (
	"fmt"
	"os"

	"github.com/DTan13/template/cmd/handlers"
	"github.com/DTan13/template/domain"
	"github.com/DTan13/template/service"
)

func Start() {
	// Get a better argument parsing library like https://github.com/jessevdk/go-flags
	args := os.Args[1:]

	domainHandler := handlers.DomainHandler{Service: service.NewTemplateDomain(domain.NewTemplateDomainStud())}
	if len(args) > 0 { // TODO: Is this that efficient?
		switch args[0] {
		case "domains":
			domainHandler.GetAllDomains()
		case "help":
			help()
		default:
			help()
		}
	} else {
		help()
	}
}

func help() {
	fmt.Println("This is help for this template!\ndomains\t\tGet the list of domains.\nhelp\t\tPrint this help.")
}
