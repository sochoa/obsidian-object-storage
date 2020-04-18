package static

import (
	"fmt"
	"github.com/sochoa/obsidian/static/config"
	"log"
	"net/http"
)

func Serve(staticConfig config.StaticConfig) {
	fs := http.FileServer(http.Dir(staticConfig.Root))
	http.Handle(staticConfig.UriBase, fs)
	log.Println(fmt.Sprintf("Listening on %s:...", staticConfig.FormatEndpoint()))
	if err := http.ListenAndServe(staticConfig.FormatEndpoint(), nil); err != nil {
		log.Fatal(err)
	}
}
