package crud

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sochoa/obsidian/crud/config"
	"github.com/sochoa/obsidian/crud/object"
	"log"
	"net/http"
)

type muxRequestHandler func(http.ResponseWriter, *http.Request)

type configuredMuxRequestHandler func(config.ObjectStorageConfig, http.ResponseWriter, *http.Request)

func SetupObjectRoutes(requestRouter *mux.Router, cfg config.ObjectStorageConfig) {

	var (
		bucketRegex      string = "\\w[\\w\\d-_\\.]*"
		bucketUriPattern string = fmt.Sprintf("{bucket:%s}", bucketRegex)
		pathRegex        string = "\\w[\\w\\d-_\\.]*"
		pathUriPattern   string = fmt.Sprintf("{path:%s}", pathRegex)
		uriPattern       string = fmt.Sprintf("/%s/%s", bucketUriPattern, pathUriPattern)
	)

	wrapHandlerWithConfig := func(hdlr configuredMuxRequestHandler) muxRequestHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			hdlr(cfg, w, r)
		}
	}

	log.Printf("Configuring request hanlder:  \"%s\" for POST requests", uriPattern)
	requestRouter.HandleFunc(uriPattern, wrapHandlerWithConfig(object.PostObject)).Methods("POST")

	log.Printf("Configuring request hanlder:  \"%s\" for GET requests", uriPattern)
	requestRouter.HandleFunc(uriPattern, wrapHandlerWithConfig(object.GetObject)).Methods("GET")
}
