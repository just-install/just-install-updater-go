package jiup

import (
	"fmt"
	"reflect"
	"sync"
)

// RuleData stores shared data between Versioners and Downloaders (for example,
// extra metadata used to get the download link corresponding to a version).
//
// To enforce access, you must define a public or private string type to use as
// a tag to access the data, like:
//
//     type myPrivateData string
//     ...
//     (*RuleData).Set(myPrivateData("something"), "whatever")
//     ...
//     (*RuleData).Get(myPrivateData("something"))
type RuleData struct {
	mu   sync.Mutex
	data map[string]map[string]interface{}
}

// MustGet is like Get, but panics if the tag is not set.
func (r *RuleData) MustGet(tag interface{}) interface{} {
	v, ok := r.Get(tag)
	if !ok {
		panic(fmt.Sprintf("tag %#v not set", tag))
	}
	return v
}

// Get gets the data assigned to the specified tag.
func (r *RuleData) Get(tag interface{}) (interface{}, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tkey, vkey := splitTag(tag)
	if r.data == nil || r.data[tkey] == nil {
		return nil, false
	}
	v, ok := r.data[tkey][vkey]
	return v, ok
}

// Set assigns data to the specified tag.
func (r *RuleData) Set(tag interface{}, data interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tkey, vkey := splitTag(tag)
	switch {
	case r.data == nil:
		r.data = map[string]map[string]interface{}{}
		fallthrough
	case r.data[tkey] == nil:
		r.data[tkey] = map[string]interface{}{}
	}
	r.data[tkey][vkey] = data
}

func splitTag(tag interface{}) (string, string) {
	v := reflect.ValueOf(tag)
	if v.Kind() != reflect.String {
		panic("tag is not a string type")
	} else if v.Type().String() == "string" {
		panic("type is a raw string")
	}
	return v.Type().String(), v.String()
}
