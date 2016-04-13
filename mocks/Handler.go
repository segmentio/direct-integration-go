package mocks

import "github.com/segmentio/direct-integration-go"
import "github.com/stretchr/testify/mock"

type Handler struct {
	mock.Mock
}

// HandleTrack provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandleTrack(_a0 integration.TrackMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}

// HandlePage provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandlePage(_a0 integration.PageMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}

// HandleScreen provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandleScreen(_a0 integration.ScreenMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}

// HandleIdentify provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandleIdentify(_a0 integration.IdentifyMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}

// HandleAlias provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandleAlias(_a0 integration.AliasMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}

// HandleGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *Handler) HandleGroup(_a0 integration.GroupMessage, _a1 *integration.ResponseWriter, _a2 *integration.Request) {
	_m.Called(_a0, _a1, _a2)
}
