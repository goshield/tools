package tools

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/goshield/interfaces"
)

// NewBag returns an instance of Bag
func NewBag() interfaces.Bag {
	return &factoryBag{items: make(map[string]interface{})}
}

// NewBagWithValues returns a bag with provided data
func NewBagWithValues(data map[string]interface{}) interfaces.Bag {
	return &factoryBag{items: data}
}

type factoryBag struct {
	items map[string]interface{}
}

func (b *factoryBag) GetOrDefault(key string, def interface{}) interface{} {
	if b.Has(key) {
		return b.Get(key)
	}
	return def
}

func (b *factoryBag) Get(key string) interface{} {
	return b.items[key]
}

func (b *factoryBag) Set(key string, value interface{}) {
	b.items[key] = value
}

func (b *factoryBag) Remove(key string) {
	delete(b.items, key)
}

func (b *factoryBag) Has(key string) bool {
	_, ok := b.items[key]
	return ok
}

func (b *factoryBag) All() map[string]interface{} {
	return b.items
}

func (b *factoryBag) GetInt(key string) int64 {
	value, ok := b.items[key]
	if ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(value).Int()
		case reflect.String:
			s := reflect.ValueOf(value).String()
			i, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				return i
			}
		}
	}

	return 0
}

func (b *factoryBag) GetUInt(key string) uint64 {
	value, ok := b.items[key]
	if ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(value).Uint()
		case reflect.String:
			s := reflect.ValueOf(value).String()
			i, err := strconv.ParseUint(s, 10, 64)
			if err == nil {
				return i
			}
		}
	}

	return 0
}

func (b *factoryBag) GetFloat(key string) float64 {
	value, ok := b.items[key]
	if ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(value).Float()
		case reflect.String:
			s := reflect.ValueOf(value).String()
			i, err := strconv.ParseFloat(s, 64)
			if err == nil {
				return i
			}
		}
	}

	return 0.0
}

func (b *factoryBag) GetString(key string) string {
	value, ok := b.items[key]
	if ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			return reflect.ValueOf(value).String()
		default:
			return fmt.Sprintf("%v", value)
		}
	}

	return ""
}

func (b *factoryBag) GetBool(key string) bool {
	value, ok := b.items[key]
	if ok {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Bool:
			return reflect.ValueOf(value).Bool()
		case reflect.String:
			v := reflect.ValueOf(value).String()
			if v == "true" || v == "1" {
				return true
			} else if v == "false" || v == "0" {
				return false
			}
		}
	}

	return false
}
