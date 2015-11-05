package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/emicklei/go-restful"
)

func main() {
	c := restful.NewContainer()
	ws := new(restful.WebService)
	ws.
		Path("/").
		Doc("API Gateway")

	ws.Route(ws.GET("/").To(gateway).
		Doc("gateway function").
		Operation("gateway"))

	ws.Route(ws.GET("/{asdf}").To(gateway).
		Operation("gateway"))

	ws.Route(ws.GET("/{asdf}/{asdf}").To(gateway).
		Operation("gateway"))

	c.Add(ws)
	server := &http.Server{Addr: "localhost:8080", Handler: c}

	log.Printf("start listening on localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func gateway(request *restful.Request, response *restful.Response) {
	log.Printf(request.HeaderParameter("API-Version"))

	if request.HeaderParameter("API-Version") == "1.0" {
		target, err := url.Parse("http://localhost:8008")
		if err != nil {
			log.Print(err)
			return
		}
		c := httputil.NewSingleHostReverseProxy(target)
		c.ServeHTTP(response.ResponseWriter, request.Request)
	}

	if request.HeaderParameter("API-Version") == "1.1" {
		target, err := url.Parse("http://localhost:8008")
		if err != nil {
			log.Print(err)
			return
		}
		c := httputil.NewSingleHostReverseProxy(target)
		c.ServeHTTP(response.ResponseWriter, request.Request)
	}

	response.WriteErrorString(http.StatusNotFound, "404: Version not found!.")
}
