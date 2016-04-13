package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NewServer returns an http.Handler that parses HTTP requests into Segment
// messages and calls methods on the given integration handler.
func NewServer(h Handler) http.Handler {
	return &server{h}
}

type server struct {
	h Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var msg struct {
		Type string `json:"type"`
	}
	var buf bytes.Buffer
	defer req.Body.Close()
	dec := json.NewDecoder(io.TeeReader(req.Body, &buf))
	if err := dec.Decode(&msg); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	dec = json.NewDecoder(&buf)
	dec.UseNumber()

	switch msg.Type {
	case "track":
		var track TrackMessage
		if err := dec.Decode(&track); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandleTrack(track, &ResponseWriter{w}, &Request{req})
	case "page":
		var page PageMessage
		if err := dec.Decode(&page); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandlePage(page, &ResponseWriter{w}, &Request{req})
	case "screen":
		var screen ScreenMessage
		if err := dec.Decode(&screen); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandleScreen(screen, &ResponseWriter{w}, &Request{req})
	case "identify":
		var identify IdentifyMessage
		if err := dec.Decode(&identify); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandleIdentify(identify, &ResponseWriter{w}, &Request{req})
	case "alias":
		var alias AliasMessage
		if err := dec.Decode(&alias); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandleAlias(alias, &ResponseWriter{w}, &Request{req})
	case "group":
		var group GroupMessage
		if err := dec.Decode(&group); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		s.h.HandleGroup(group, &ResponseWriter{w}, &Request{req})
	default:
		http.Error(w, fmt.Sprintf("%s: unknown message type (%s)", http.StatusText(http.StatusBadRequest), msg.Type), http.StatusBadRequest)
		return
	}
}
