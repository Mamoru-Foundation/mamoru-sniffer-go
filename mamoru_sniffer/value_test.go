package mamoru_sniffer

import (
	"testing"
)

func TestNewValueBoolDoesNotFail(t *testing.T) {
	_ = NewValueBool(false)
	_ = NewValueBool(true)
}

func TestNewValueListDoesNotFail(t *testing.T) {
	_ = NewValueList([]Value{
		NewValueBool(false),
		NewValueBool(true),
	})

	_ = NewValueList([]Value{
		NewValueU64(0),
		NewValueU64(1),
		NewValueU64(42),
	})

	_ = NewValueList([]Value{
		NewValueList([]Value{
			NewValueStruct("hello!", map[string]Value{
				"field":  NewValueBool(false),
				"field2": NewValueU64(42),
			}),
		}),
	})
}

func TestNewValueStructDoesNotFail(t *testing.T) {
	_ = NewValueStruct("hello!", map[string]Value{
		"field":  NewValueBool(false),
		"field2": NewValueU64(42),
		"fieldStruct": NewValueStruct("meow :3", map[string]Value{
			"meowField": NewValueBool(true),
		}),
	})
}

func TestNewValueU64DoesNotFail(t *testing.T) {
	_ = NewValueU64(0)
	_ = NewValueU64(42)
	_ = NewValueU64(5000000000)
}

func TestNewValueDataDoesNotFail(t *testing.T) {
	NewValueData(NewValueU64(1))
}
