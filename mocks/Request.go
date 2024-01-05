// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	playwright "github.com/playwright-community/playwright-go"
	mock "github.com/stretchr/testify/mock"
)

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// AllHeaders provides a mock function with given fields:
func (_m *Request) AllHeaders() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Failure provides a mock function with given fields:
func (_m *Request) Failure() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Frame provides a mock function with given fields:
func (_m *Request) Frame() playwright.Frame {
	ret := _m.Called()

	var r0 playwright.Frame
	if rf, ok := ret.Get(0).(func() playwright.Frame); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(playwright.Frame)
		}
	}

	return r0
}

// HeaderValue provides a mock function with given fields: name
func (_m *Request) HeaderValue(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Headers provides a mock function with given fields:
func (_m *Request) Headers() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// HeadersArray provides a mock function with given fields:
func (_m *Request) HeadersArray() ([]playwright.NameValue, error) {
	ret := _m.Called()

	var r0 []playwright.NameValue
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]playwright.NameValue, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []playwright.NameValue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]playwright.NameValue)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNavigationRequest provides a mock function with given fields:
func (_m *Request) IsNavigationRequest() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *Request) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PostData provides a mock function with given fields:
func (_m *Request) PostData() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostDataBuffer provides a mock function with given fields:
func (_m *Request) PostDataBuffer() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostDataJSON provides a mock function with given fields: v
func (_m *Request) PostDataJSON(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedirectedFrom provides a mock function with given fields:
func (_m *Request) RedirectedFrom() playwright.Request {
	ret := _m.Called()

	var r0 playwright.Request
	if rf, ok := ret.Get(0).(func() playwright.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(playwright.Request)
		}
	}

	return r0
}

// RedirectedTo provides a mock function with given fields:
func (_m *Request) RedirectedTo() playwright.Request {
	ret := _m.Called()

	var r0 playwright.Request
	if rf, ok := ret.Get(0).(func() playwright.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(playwright.Request)
		}
	}

	return r0
}

// ResourceType provides a mock function with given fields:
func (_m *Request) ResourceType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Response provides a mock function with given fields:
func (_m *Request) Response() (playwright.Response, error) {
	ret := _m.Called()

	var r0 playwright.Response
	var r1 error
	if rf, ok := ret.Get(0).(func() (playwright.Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() playwright.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(playwright.Response)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sizes provides a mock function with given fields:
func (_m *Request) Sizes() (*playwright.RequestSizesResult, error) {
	ret := _m.Called()

	var r0 *playwright.RequestSizesResult
	var r1 error
	if rf, ok := ret.Get(0).(func() (*playwright.RequestSizesResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *playwright.RequestSizesResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*playwright.RequestSizesResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Timing provides a mock function with given fields:
func (_m *Request) Timing() *playwright.RequestTiming {
	ret := _m.Called()

	var r0 *playwright.RequestTiming
	if rf, ok := ret.Get(0).(func() *playwright.RequestTiming); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*playwright.RequestTiming)
		}
	}

	return r0
}

// URL provides a mock function with given fields:
func (_m *Request) URL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewRequest creates a new instance of Request. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *Request {
	mock := &Request{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
