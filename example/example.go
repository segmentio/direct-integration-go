package main

import (
	"flag"
	"log"

	"net/http"

	"github.com/segmentio/direct-integration-go"
)

var port = flag.String("port", "8080", "http port")

func main() {
	server := integration.NewServer(&handler{})
	if err := http.ListenAndServe(":"+*port, server); err != nil {
		log.Fatal(err)
	}
}

type handler struct{}

func (h *handler) HandleTrack(track integration.TrackMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("track:", track)
}

func (h *handler) HandlePage(page integration.PageMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("page:", page)
}

func (h *handler) HandleScreen(screen integration.ScreenMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("screen:", screen)
}

func (h *handler) HandleIdentify(identify integration.IdentifyMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("identify:", identify)
}

func (h *handler) HandleAlias(alias integration.AliasMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("alias:", alias)
}

func (h *handler) HandleGroup(group integration.GroupMessage, w *integration.ResponseWriter, r *integration.Request) {
	log.Println("group:", group)
}
