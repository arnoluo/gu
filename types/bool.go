package types

// If 根据条件判断返回不同的值。
func (bt BoolType) If(isTrue, trueValue, falseValue bool) bool {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// And 对 bool 类型进行逻辑与操作。
func (bt BoolType) And(values ...bool) bool {
	result := true
	for _, v := range values {
		result = result && v
		if !result {
			break
		}
	}
	return result
}

// Or 对 bool 类型进行逻辑或操作。
func (bt BoolType) Or(values ...bool) bool {
	result := false
	for _, v := range values {
		result = result || v
		if result {
			break
		}
	}
	return result
}

// Not 对 bool 类型进行逻辑非操作。
func (bt BoolType) Not(value bool) bool {
	return !value
}
