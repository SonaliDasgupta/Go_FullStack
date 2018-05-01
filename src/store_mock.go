package main

import (
	"github.com/stretchr/testify/mock"

)

type MockStore struct{
	mock.Mock
}

func (m *MockStore) CreateBird( bird *Bird) error{
	rets := m.Called(bird)
	return rets.Error(0)
	
}

func (m *MockStore) GetBirds() ([]*Bird, error){
	rets := m.Called()
	return rets.Get(0).([]*Bird), rets.Error(1)
}

func InitMockStore() *MockStore{
	s := new(MockStore)
	store =s
	return s
}



