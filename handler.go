package integration

// Handler responds to a Segment message defined by the spec.
// It's a thin wrapper around http.Handler.ServeHTTP, so all the sames rule
// apply. The only difference is that the request body is already read before
// calling the handler.
//
// https://segment.com/docs/spec/
type Handler interface {
	HandleTrack(TrackMessage, *ResponseWriter, *Request)
	HandlePage(PageMessage, *ResponseWriter, *Request)
	HandleScreen(ScreenMessage, *ResponseWriter, *Request)
	HandleIdentify(IdentifyMessage, *ResponseWriter, *Request)
	HandleAlias(AliasMessage, *ResponseWriter, *Request)
	HandleGroup(GroupMessage, *ResponseWriter, *Request)
}
