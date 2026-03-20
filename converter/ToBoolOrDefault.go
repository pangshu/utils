package converter

// ToBoolOrDefault 将数值类型转换为bool，失败就返回默认值。 / converts numeric types to bool, returns defaultValue on failure.
func (c *Converter[T]) ToBoolOrDefault(v T, defaultValue bool) bool {
	result, err := c.ToBool(v)
	if err != nil {
		return defaultValue
	}
	return result
}
