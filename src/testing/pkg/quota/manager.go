// Code generated by mockery v1.0.0. DO NOT EDIT.

package quota

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/quota/models"
	mock "github.com/stretchr/testify/mock"

	types "github.com/goharbor/harbor/src/pkg/types"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, reference, referenceID, hardLimits, usages
func (_m *Manager) Create(ctx context.Context, reference string, referenceID string, hardLimits types.ResourceList, usages ...types.ResourceList) (int64, error) {
	_va := make([]interface{}, len(usages))
	for _i := range usages {
		_va[_i] = usages[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reference, referenceID, hardLimits)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, types.ResourceList, ...types.ResourceList) int64); ok {
		r0 = rf(ctx, reference, referenceID, hardLimits, usages...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, types.ResourceList, ...types.ResourceList) error); ok {
		r1 = rf(ctx, reference, referenceID, hardLimits, usages...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Manager) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Manager) Get(ctx context.Context, id int64) (*models.Quota, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Quota
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Quota); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRef provides a mock function with given fields: ctx, reference, referenceID
func (_m *Manager) GetByRef(ctx context.Context, reference string, referenceID string) (*models.Quota, error) {
	ret := _m.Called(ctx, reference, referenceID)

	var r0 *models.Quota
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Quota); ok {
		r0 = rf(ctx, reference, referenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, reference, referenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRefForUpdate provides a mock function with given fields: ctx, reference, referenceID
func (_m *Manager) GetByRefForUpdate(ctx context.Context, reference string, referenceID string) (*models.Quota, error) {
	ret := _m.Called(ctx, reference, referenceID)

	var r0 *models.Quota
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Quota); ok {
		r0 = rf(ctx, reference, referenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, reference, referenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Manager) Update(ctx context.Context, _a1 *models.Quota) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Quota) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
