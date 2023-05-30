package router

import (
	"fmt"
	"net/http"
)

type muxRouter struct{}

var muxServer = http.NewServeMux()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxServer.HandleFunc(uri, f)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxServer.HandleFunc(uri, f)
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v \n", port)
	http.ListenAndServe(port, muxServer)
}
