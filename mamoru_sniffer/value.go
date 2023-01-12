package mamoru_sniffer

import "github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"

type ValueData struct {
	*generated_bindings.FfiValueDataT
}

func NewValueData(value Value) ValueData {
	valueData := generated_bindings.NewValueData(value.FfiValueT)

	return ValueData{valueData}
}

type Value struct {
	*generated_bindings.FfiValueT
}

func NewValueBool(data bool) Value {
	return Value{
		generated_bindings.NewValueBool(data),
	}
}

func NewValueU64(data uint64) Value {
	return Value{
		generated_bindings.NewValueU64(data),
	}
}

func NewValueList(list []Value) Value {
	valueList := generated_bindings.NewValueList()

	for _, item := range list {
		// may return `false` if Value is not ValueList
		// ignoring here, as we guarantee the type is right
		_ = generated_bindings.ValueListAppend(valueList, item.FfiValueT)
	}

	return Value{valueList}
}

func NewValueStruct(ty string, fields map[string]Value) Value {
	valueStruct := generated_bindings.NewValueStruct(ty)

	for field, value := range fields {
		// may return `false` if Value is not ValueStruct of if there is a duplicate field
		// ignoring here, as we guarantee the type is right and there is no duplicates
		_ = generated_bindings.ValueStructAddField(valueStruct, field, value.FfiValueT)
	}

	return Value{valueStruct}
}
