package router

import (
	"github.com/cherryReptile/GoSeeder/internal/jobs"
	"github.com/cherryReptile/GoSeeder/internal/queue"
	"github.com/cherryReptile/GoSeeder/internal/responses"
	"net/http"
)

type Router struct {
	Worker *queue.JobWorker
}

func (router *Router) Index(w http.ResponseWriter, r *http.Request) {
	responseJson(w, responses.VersionResponse{Version: "1.0", Name: "Starter Kit v1.0"})
}

func (router *Router) Test(w http.ResponseWriter, r *http.Request) {
	var j jobs.TestJob
	var i interface{}

	j.Init(i)

	router.Worker.Add(&j)
}
