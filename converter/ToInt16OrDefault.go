package converter

// ToInt16OrDefault 将数值类型转换为Int16，失败就返回默认值。 / converts numeric types to Int8, returns defaultValue on failure.
func (c *Converter[T]) ToInt16OrDefault(v T, defaultValue int16) int16 {
	result, err := c.ToInt16(v)
	if err != nil {
		return defaultValue
	}
	return result
}
