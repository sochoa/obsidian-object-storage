package object

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sochoa/obsidian/crud/config"
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

	// Injecting dependency with Currying
	wrapHandlerWithConfig := func(hdlr configuredMuxRequestHandler) muxRequestHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			hdlr(cfg, w, r)
		}
	}

	handlerTypes := map[string]configuredMuxRequestHandler{
		"POST":   PostObject,
		"GET":    GetObject,
		"DELETE": DeleteObject,
		//"PUT":    PutObject,
	}

	for key, val := range handlerTypes {
		log.Printf("Configuring request hanlderfor %s requests using URI pattern: \"%s\"", key, uriPattern)
		requestRouter.HandleFunc(uriPattern, wrapHandlerWithConfig(val)).Methods(key)
	}
}
