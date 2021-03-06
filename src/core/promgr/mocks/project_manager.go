// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	metamgr "github.com/goharbor/harbor/src/core/promgr/metamgr"
	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/common/models"
)

// ProjectManager is an autogenerated mock type for the ProjectManager type
type ProjectManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *ProjectManager) Create(_a0 *models.Project) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*models.Project) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Project) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: projectIDOrName
func (_m *ProjectManager) Delete(projectIDOrName interface{}) error {
	ret := _m.Called(projectIDOrName)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(projectIDOrName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: projectIDOrName
func (_m *ProjectManager) Exists(projectIDOrName interface{}) (bool, error) {
	ret := _m.Called(projectIDOrName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(projectIDOrName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(projectIDOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: projectIDOrName
func (_m *ProjectManager) Get(projectIDOrName interface{}) (*models.Project, error) {
	ret := _m.Called(projectIDOrName)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(interface{}) *models.Project); ok {
		r0 = rf(projectIDOrName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(projectIDOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadataManager provides a mock function with given fields:
func (_m *ProjectManager) GetMetadataManager() metamgr.ProjectMetadataManager {
	ret := _m.Called()

	var r0 metamgr.ProjectMetadataManager
	if rf, ok := ret.Get(0).(func() metamgr.ProjectMetadataManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metamgr.ProjectMetadataManager)
		}
	}

	return r0
}

// GetPublic provides a mock function with given fields:
func (_m *ProjectManager) GetPublic() ([]*models.Project, error) {
	ret := _m.Called()

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func() []*models.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
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

// IsPublic provides a mock function with given fields: projectIDOrName
func (_m *ProjectManager) IsPublic(projectIDOrName interface{}) (bool, error) {
	ret := _m.Called(projectIDOrName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(projectIDOrName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(projectIDOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: query
func (_m *ProjectManager) List(query *models.ProjectQueryParam) (*models.ProjectQueryResult, error) {
	ret := _m.Called(query)

	var r0 *models.ProjectQueryResult
	if rf, ok := ret.Get(0).(func(*models.ProjectQueryParam) *models.ProjectQueryResult); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProjectQueryResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ProjectQueryParam) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: projectIDOrName, project
func (_m *ProjectManager) Update(projectIDOrName interface{}, project *models.Project) error {
	ret := _m.Called(projectIDOrName, project)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *models.Project) error); ok {
		r0 = rf(projectIDOrName, project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
