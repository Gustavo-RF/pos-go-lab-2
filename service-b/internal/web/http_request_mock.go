package web

import "github.com/stretchr/testify/mock"

type MockRequestFunc struct {
	mock.Mock
}

func (m *MockRequestFunc) Request(url, method string) ([]byte, error) {
	args := m.Called(url, method)
	data, _ := args.Get(0).([]byte)
	return data, args.Error(1)
}
