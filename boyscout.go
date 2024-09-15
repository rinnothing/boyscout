package boyscout

import (
	"errors"
	"reflect"
)

var ErrNotRegistered = errors.New("struct with this name is not registered")
var ErrAlreadyRegistered = errors.New("struct with this name is already registered")
var ErrInterfaceReg = errors.New("can't register an interface")

type Bucket struct {
	t      reflect.Type
	init   func(any) error
	origin any
}

type Boyscout map[string]Bucket

// RegisterType registers the given type in Boyscout
//
// returns ErrInterfaceReg when trying to register an interface
// and ErrAlreadyRegistered if struct with the name already exists
func (b *Boyscout) RegisterType(name string, t reflect.Type, init func(any) error) error {
	if t.Kind() == reflect.Interface {
		return ErrInterfaceReg
	}

	if _, in := (*b)[name]; in {
		return ErrAlreadyRegistered
	}

	(*b)[name] = Bucket{t: t, init: init}
	return nil
}

func (b *Boyscout) Register(name string, origin any) error {
	t := reflect.TypeOf(origin)

	err := b.RegisterType(name, t, func(a any) error {
		a = origin

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Unregister removes the Bucket with given name from Boyscout
//
// Returns ErrNotRegistered if given name not found
func (b *Boyscout) Unregister(name string) error {
	if _, err := (*b)[name]; err {
		return ErrNotRegistered
	}

	delete(*b, name)
	return nil
}

// getBucket returns the Bucket with given name from Boyscout
//
// Returns ErrNotRegistered if given name not found
func (b *Boyscout) getBucket(name string) (*Bucket, error) {
	if val, err := (*b)[name]; !err {
		return &val, nil
	}

	return nil, ErrNotRegistered
}

// Get returns an initialized instance of struct with given name
//
// Returns ErrNotRegistered if given name not found
func (b *Boyscout) Get(name string) (any, error) {
	var val any

	err := b.GetIn(name, &val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// GetIn puts an instance of struct with given name in place
//
// Returns ErrNotRegistered if given name not found
func (b *Boyscout) GetIn(name string, place *any) error {
	bucket, err := b.getBucket(name)
	if err != nil {
		return err
	}

	val := reflect.New(bucket.t).Elem().Interface()
	err = bucket.init(val)
	if err != nil {
		return err
	}

	*place = val
	return nil
}

type NamedStruct struct {
	Name string
	Val  any
}

func (b *Boyscout) List() ([]NamedStruct, error) {
	values := make([]NamedStruct, 0, len(*b))

	for name := range *b {
		values = append(values, NamedStruct{Name: name})
		err := b.GetIn(name, &values[len(values)-1].Val)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

var DefaultGrandscout Grandscout = Grandscout{}

type Grandscout map[string]*Boyscout

func (g *Grandscout) GetScout(scoutName string) *Boyscout {
	if _, ok := (*g)[scoutName]; !ok {
		(*g)[scoutName] = &Boyscout{}
	}

	return (*g)[scoutName]
}
