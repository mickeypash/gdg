// Code generated by mockery v2.34.0. DO NOT EDIT.

package mocks

import (
	filters "github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"

	models "github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// TeamsApi is an autogenerated mock type for the TeamsApi type
type TeamsApi struct {
	mock.Mock
}

// DeleteTeam provides a mock function with given fields: filter
func (_m *TeamsApi) DeleteTeam(filter filters.Filter) ([]*models.TeamDTO, error) {
	ret := _m.Called(filter)

	var r0 []*models.TeamDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(filters.Filter) ([]*models.TeamDTO, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(filters.Filter) []*models.TeamDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TeamDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(filters.Filter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadTeams provides a mock function with given fields: filter
func (_m *TeamsApi) DownloadTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// ListTeams provides a mock function with given fields: filter
func (_m *TeamsApi) ListTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// UploadTeams provides a mock function with given fields: filter
func (_m *TeamsApi) UploadTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// NewTeamsApi creates a new instance of TeamsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTeamsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *TeamsApi {
	mock := &TeamsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}