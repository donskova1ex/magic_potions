// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/donskova1ex/magic_potions/internal/processors (interfaces: IngredientsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/donskova1ex/magic_potions/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// IngredientsRepository is a mock of IngredientsRepository interface.
type IngredientsRepository struct {
	ctrl     *gomock.Controller
	recorder *IngredientsRepositoryMockRecorder
}

// IngredientsRepositoryMockRecorder is the mock recorder for IngredientsRepository.
type IngredientsRepositoryMockRecorder struct {
	mock *IngredientsRepository
}

// NewIngredientsRepository creates a new mock instance.
func NewIngredientsRepository(ctrl *gomock.Controller) *IngredientsRepository {
	mock := &IngredientsRepository{ctrl: ctrl}
	mock.recorder = &IngredientsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *IngredientsRepository) EXPECT() *IngredientsRepositoryMockRecorder {
	return m.recorder
}

// CreateIngredient mocks base method.
func (m *IngredientsRepository) CreateIngredient(arg0 context.Context, arg1 *domain.Ingredient) (*domain.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIngredient", arg0, arg1)
	ret0, _ := ret[0].(*domain.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIngredient indicates an expected call of CreateIngredient.
func (mr *IngredientsRepositoryMockRecorder) CreateIngredient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIngredient", reflect.TypeOf((*IngredientsRepository)(nil).CreateIngredient), arg0, arg1)
}

// DeleteIngredientByUUID mocks base method.
func (m *IngredientsRepository) DeleteIngredientByUUID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIngredientByUUID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIngredientByUUID indicates an expected call of DeleteIngredientByUUID.
func (mr *IngredientsRepositoryMockRecorder) DeleteIngredientByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIngredientByUUID", reflect.TypeOf((*IngredientsRepository)(nil).DeleteIngredientByUUID), arg0, arg1)
}

// IngredientByUUID mocks base method.
func (m *IngredientsRepository) IngredientByUUID(arg0 context.Context, arg1 string) (*domain.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngredientByUUID", arg0, arg1)
	ret0, _ := ret[0].(*domain.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngredientByUUID indicates an expected call of IngredientByUUID.
func (mr *IngredientsRepositoryMockRecorder) IngredientByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngredientByUUID", reflect.TypeOf((*IngredientsRepository)(nil).IngredientByUUID), arg0, arg1)
}

// IngredientsAll mocks base method.
func (m *IngredientsRepository) IngredientsAll(arg0 context.Context) ([]*domain.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngredientsAll", arg0)
	ret0, _ := ret[0].([]*domain.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngredientsAll indicates an expected call of IngredientsAll.
func (mr *IngredientsRepositoryMockRecorder) IngredientsAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngredientsAll", reflect.TypeOf((*IngredientsRepository)(nil).IngredientsAll), arg0)
}

// UpdateIngredientByUUID mocks base method.
func (m *IngredientsRepository) UpdateIngredientByUUID(arg0 context.Context, arg1 *domain.Ingredient) (*domain.Ingredient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIngredientByUUID", arg0, arg1)
	ret0, _ := ret[0].(*domain.Ingredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIngredientByUUID indicates an expected call of UpdateIngredientByUUID.
func (mr *IngredientsRepositoryMockRecorder) UpdateIngredientByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngredientByUUID", reflect.TypeOf((*IngredientsRepository)(nil).UpdateIngredientByUUID), arg0, arg1)
}
