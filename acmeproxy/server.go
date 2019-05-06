package acmeproxy

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/xenolf/lego/challenge"
)

type Server struct {
	HttpServer     *http.Server
	Provider       challenge.Provider
	HtpasswdFile   string
	AllowedDomains []string
	AccesslogFile  string
}

func NewServer(config *Config) (*Server, error) {
	return &Server{
		HttpServer:     config.HttpServer,
		Provider:       config.Provider,
		HtpasswdFile:   config.HtpasswdFile,
		AllowedDomains: config.AllowedDomains,
		AccesslogFile:  config.AccesslogFile,
	}, nil
}

func RunServer(config *Config) {
	if config.HttpServer.TLSConfig != nil {
		log.WithFields(log.Fields{
			"addr":          "https://" + config.HttpServer.Addr,
			"htpasswdfile":  config.HtpasswdFile,
			"accesslogfile": config.AccesslogFile,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServeTLS("", ""))
	} else {
		log.WithFields(log.Fields{
			"addr":          "http://" + config.HttpServer.Addr,
			"htpasswdfile":  config.HtpasswdFile,
			"accesslogfile": config.AccesslogFile,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServe())
	}
}