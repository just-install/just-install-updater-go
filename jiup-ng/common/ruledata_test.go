package common

import (
	"fmt"
	"testing"
)

func TestRuleData(t *testing.T) {
	type myTag string
	type myOtherTag string

	data := new(RuleData)

	t.Run("BadTypes", func(t *testing.T) {
		if doPanic(func() { data.Get(123) }) == nil {
			t.Error("should have panicked on non-int as tag")
		}
		if doPanic(func() { data.Get("asd") }) == nil {
			t.Error("should have panicked on raw string as tag")
		}
	})

	t.Run("GetSet", func(t *testing.T) {
		if v, ok := data.Get(myTag("asd")); v != nil || ok {
			t.Errorf("should have returned (nil, false) for nonexistent value, got (%#v, %#v)", v, ok)
		}

		data.Set(myTag("asd"), "sdf")
		if v, ok := data.Get(myTag("asd")); v.(string) != "sdf" || !ok {
			t.Errorf("should have returned (\"sdf\", false) for first tag type, got (%#v, %#v)", v, ok)
		}

		if v, ok := data.Get(myOtherTag("asd")); v != nil || ok {
			t.Errorf("should have returned (nil, false) for other tag type, got (%#v, %#v)", v, ok)
		}

		data.Set(myOtherTag("asd"), "dfg")
		if v, ok := data.Get(myOtherTag("asd")); v.(string) != "dfg" || !ok {
			t.Errorf("should have returned (\"dfg\", false) for other tag type, got (%#v, %#v)", v, ok)
		}

		if v, ok := data.Get(myTag("asd")); v.(string) != "sdf" || !ok {
			t.Errorf("should have still returned (\"sdf\", false) for first tag type, got (%#v, %#v)", v, ok)
		}
	})
}

func doPanic(fn func()) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("panic: %v", perr)
		}
	}()
	fn()
	return
}
