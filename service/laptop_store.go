package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"grpc_fundamental/proto"
	"sync"
)

var (
	ErrAlreadyExist = errors.New("record already exist")
)

type LaptopStore interface {
	Save(laptop *grpc_fundamental.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex   sync.RWMutex
	laptops map[string]*grpc_fundamental.Laptop
}

func NewMapLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		laptops: make(map[string]*grpc_fundamental.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *grpc_fundamental.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.laptops[laptop.Id] != nil {
		return ErrAlreadyExist
	}

	another := &grpc_fundamental.Laptop{}
	err := copier.Copy(another, laptop)
	if err != nil {
		return err
	}

	store.laptops[another.Id] = another
	return nil
}
