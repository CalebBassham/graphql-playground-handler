package playground

import (
	"net/http"
)

func PlaygroundMiddleware(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(playgroundHtml)
}

