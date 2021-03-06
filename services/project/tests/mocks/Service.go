// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	models "in-backend/services/project/models"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateCandidateProject provides a mock function with given fields: ctx, m
func (_m *Service) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	ret := _m.Called(ctx, m)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.CandidateProject) error); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProject provides a mock function with given fields: ctx, m, cid
func (_m *Service) CreateProject(ctx context.Context, m *models.Project, cid uint64) (*models.Project, error) {
	ret := _m.Called(ctx, m, cid)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project, uint64) *models.Project); ok {
		r0 = rf(ctx, m, cid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Project, uint64) error); ok {
		r1 = rf(ctx, m, cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRating provides a mock function with given fields: ctx, m
func (_m *Service) CreateRating(ctx context.Context, m *models.Rating) error {
	ret := _m.Called(ctx, m)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Rating) error); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCandidateProject provides a mock function with given fields: ctx, id
func (_m *Service) DeleteCandidateProject(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProject provides a mock function with given fields: ctx, id
func (_m *Service) DeleteProject(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRating provides a mock function with given fields: ctx, id
func (_m *Service) DeleteRating(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProjects provides a mock function with given fields: ctx, f
func (_m *Service) GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func(context.Context, models.ProjectFilters) []*models.Project); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ProjectFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByID provides a mock function with given fields: ctx, id
func (_m *Service) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Project); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanProject provides a mock function with given fields: ctx, id
func (_m *Service) ScanProject(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: ctx, m
func (_m *Service) UpdateProject(ctx context.Context, m *models.Project) (*models.Project, error) {
	ret := _m.Called(ctx, m)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) *models.Project); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Project) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
