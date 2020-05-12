package playground

import (
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(playgroundHtml)
}

