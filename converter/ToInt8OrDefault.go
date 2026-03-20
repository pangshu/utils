package converter

// ToInt8OrDefault 将数值类型转换为Int8，失败就返回默认值。 / converts numeric types to Int8, returns defaultValue on failure.
func (c *Converter[T]) ToInt8OrDefault(v T, defaultValue int8) int8 {
	result, err := c.ToInt8(v)
	if err != nil {
		return defaultValue
	}
	return result
}
