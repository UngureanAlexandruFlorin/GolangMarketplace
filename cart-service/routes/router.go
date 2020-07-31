package routes

import (
	"net/http"
	"strings"
)

type Router struct {
	Exists     bool
	Path       string
	MountRoute func(string)
}

func (r *Router) HandleFunc(path string, handler http.Handler) {
	// req.URL.Path = formatPath(req.URL.Path)
	http.Handle(path, handler)
}

func (r *Router) Use(path string, router *Router) string {
	if router != nil && router.Exists == true {
		router.MountRoute(path)
	} else {
		return path
	}
	return ""
}

func (r *Router) ListenAndServe(port string, handler http.Handler) {
	http.ListenAndServe(port, handler)
}

func formatPath(path string) string {
	var urlTokens []string = strings.Split(path, "/")
	var temp []string = make([]string, 0)

	for index := 2; index < len(urlTokens); index++ {
		temp = append(temp, urlTokens[index])
	}

	return "/" + strings.Join(temp, "/")
}
