package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"sync"
)

var (
	ErrAlreadyExist = errors.New("record already exist")
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex   sync.RWMutex
	laptops map[string]*pb.Laptop
}

func NewMapLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		laptops: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.laptops[laptop.Id] != nil {
		return ErrAlreadyExist
	}

	another := &pb.Laptop{}
	err := copier.Copy(another, laptop)
	if err != nil {
		return err
	}

	store.laptops[another.Id] = another
	return nil
}
