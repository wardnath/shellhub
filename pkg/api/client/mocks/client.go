// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/shellhub-io/shellhub/pkg/models"
	mock "github.com/stretchr/testify/mock"

	revdial "github.com/shellhub-io/shellhub/pkg/revdial"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AuthDevice provides a mock function with given fields: req
func (_m *Client) AuthDevice(req *models.DeviceAuthRequest) (*models.DeviceAuthResponse, error) {
	ret := _m.Called(req)

	var r0 *models.DeviceAuthResponse
	if rf, ok := ret.Get(0).(func(*models.DeviceAuthRequest) *models.DeviceAuthResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DeviceAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.DeviceAuthRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthPublicKey provides a mock function with given fields: req, token
func (_m *Client) AuthPublicKey(req *models.PublicKeyAuthRequest, token string) (*models.PublicKeyAuthResponse, error) {
	ret := _m.Called(req, token)

	var r0 *models.PublicKeyAuthResponse
	if rf, ok := ret.Get(0).(func(*models.PublicKeyAuthRequest, string) *models.PublicKeyAuthResponse); ok {
		r0 = rf(req, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKeyAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.PublicKeyAuthRequest, string) error); ok {
		r1 = rf(req, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Endpoints provides a mock function with given fields:
func (_m *Client) Endpoints() (*models.Endpoints, error) {
	ret := _m.Called()

	var r0 *models.Endpoints
	if rf, ok := ret.Get(0).(func() *models.Endpoints); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Endpoints)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: uid
func (_m *Client) GetDevice(uid string) (*models.Device, error) {
	ret := _m.Called(uid)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(string) *models.Device); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfo provides a mock function with given fields: agentVersion
func (_m *Client) GetInfo(agentVersion string) (*models.Info, error) {
	ret := _m.Called(agentVersion)

	var r0 *models.Info
	if rf, ok := ret.Get(0).(func(string) *models.Info); ok {
		r0 = rf(agentVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(agentVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDevices provides a mock function with given fields:
func (_m *Client) ListDevices() ([]models.Device, error) {
	ret := _m.Called()

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReverseListener provides a mock function with given fields: token
func (_m *Client) NewReverseListener(token string) (*revdial.Listener, error) {
	ret := _m.Called(token)

	var r0 *revdial.Listener
	if rf, ok := ret.Get(0).(func(string) *revdial.Listener); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*revdial.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}