package routes

import "net/http"

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func GetRoutes() []Route {
	return []Route{
		{
			Path:    "favicon.ico",
			Handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			Path:    "/",
			Handler: IndexRoute,
		},
		{
			Path:    "/consume",
			Handler: ConsumeRoute,
		},
	}
}
