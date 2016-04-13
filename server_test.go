package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/segmentio/direct-integration-go"
	"github.com/segmentio/direct-integration-go/mocks"
	"github.com/stretchr/testify/mock"
)

func post(url, body string) (*http.Response, error) {
	return http.Post(url, "application/json", strings.NewReader(body))
}

func assertResponseOK(t *testing.T, resp *http.Response, err error) {
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func mustParseTime(test *testing.T, value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	assert.Equal(test, nil, err)
	return t
}

func TestTrack(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandleTrack", integration.TrackMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "track",
		},
		Event: "Viewed Product",
		Properties: map[string]interface{}{
			"revenue": json.Number("4.99"),
		},
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "track",
    "event": "Viewed Product",
    "properties": {
      "revenue": 4.99
    }
  }`)
	assertResponseOK(t, resp, err)
}

func TestPage(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandlePage", integration.PageMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "page",
		},
		Name:       "Home",
		Properties: map[string]interface{}{},
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "page",
    "name": "Home",
    "properties": {}
  }`)
	assertResponseOK(t, resp, err)
}

func TestScreen(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandleScreen", integration.ScreenMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "screen",
		},
		Name:       "Home",
		Properties: map[string]interface{}{},
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "screen",
    "name": "Home",
    "properties": {}
  }`)
	assertResponseOK(t, resp, err)
}

func TestIdentify(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandleIdentify", integration.IdentifyMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "identify",
		},
		Traits: map[string]interface{}{
			"name":   "Peter Gibbons",
			"email":  "peter@initech.com",
			"plan":   "premium",
			"logins": json.Number("5"),
		},
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "identify",
    "traits": {
	    "name": "Peter Gibbons",
	    "email": "peter@initech.com",
	    "plan": "premium",
	    "logins": 5
	  }
  }`)
	assertResponseOK(t, resp, err)
}

func TestAlias(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandleAlias", integration.AliasMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "alias",
		},
		PreviousID: "97980cfea0066",
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "alias",
		"previous_id": "97980cfea0066"
  }`)
	assertResponseOK(t, resp, err)
}

func TestGroup(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	handler.On("HandleGroup", integration.GroupMessage{
		Message: integration.Message{
			AnonymousID: "507f191e810c19729de860ea",
			Context: map[string]interface{}{
				"ip":        "8.8.8.8",
				"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36",
			},
			MessageID:  "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
			ReceivedAt: mustParseTime(t, "2015-02-23T22:28:55.387Z"),
			SentAt:     mustParseTime(t, "2015-02-23T22:28:55.111Z"),
			UserID:     "97980cfea0067",
			Version:    "1.1",
			Type:       "group",
		},
		GroupID: "97980cfea0065",
		Traits: map[string]interface{}{
			"size": json.Number("50"),
			"name": "Segment",
		},
	}, mock.Anything, mock.Anything).Return(nil)

	resp, err := post(ts.URL, `{
    "anonymous_id": "507f191e810c19729de860ea",
    "context": {
      "ip": "8.8.8.8",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36"
    },
    "message_id": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "received_at": "2015-02-23T22:28:55.387Z",
    "sent_at": "2015-02-23T22:28:55.111Z",
    "user_id": "97980cfea0067",
    "version": "1.1",
    "type": "group",
		"group_id": "97980cfea0065",
    "traits": {
	    "size": 50,
	    "name": "Segment"
	  }
  }`)
	assertResponseOK(t, resp, err)
}

func Test400OnInvalidJSON(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	resp, err := post(ts.URL, `{}`)
	assert.Equal(t, nil, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func Test405OnInvalidMethod(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.Equal(t, nil, err)
	assert.Equal(t, 405, resp.StatusCode)
}

func Test400OnInvalidType(t *testing.T) {
	handler := &mocks.Handler{}
	ts := httptest.NewServer(integration.NewServer(handler))
	defer ts.Close()

	resp, err := post(ts.URL, `{"type": "tracks"}`)
	assert.Equal(t, nil, err)
	assert.Equal(t, 400, resp.StatusCode)
}
