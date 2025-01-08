package main

import (
	"errors"
	"fmt"
)

func main() {

	m1 := Newmap(3)

	m1["560086"] = "Bangalore-1"
	m1["560096"] = "Bangalore-2"
	m1["560034"] = "Bangalore-3"

	if keys, err := m1.GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("keys:", keys)
	}

	if values, err := m1.GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if err := m1.Delete("560086"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Key successfully deleted")
	}

	if err := m1.Delete("560034386"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Key successfully deleted")
	}

	var m2 map[string]any

	m2 = make(map[string]any)
	m2["560086"] = "Bangalore-1"
	m2["560096"] = "Bangalore-2"
	m2["560034"] = "Bangalore-3"

	if keys, err := mymap(m2).GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("keys:", keys)
	}

	if values, err := mymap(m2).GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	var m3 map[string]string

	m3 = make(map[string]string)
	m3["560086"] = "Bangalore-1"
	m3["560096"] = "Bangalore-2"
	m3["560034"] = "Bangalore-3"

}

type mymap map[string]any

func Newmap(size int) mymap {
	return make(mymap, size)
}

func (m mymap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys, nil
}

func (m mymap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	// values := make([]any, len(m))
	var values []any // nil
	for _, v := range m {
		values = append(values, v) // if slice is nil append can also instantiate
	}
	return values, nil
}

func (m mymap) Delete(key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	_, ok := m[key]
	if !ok {
		return errors.New("key does not exist")
	}
	delete(m, key)
	return nil
}
