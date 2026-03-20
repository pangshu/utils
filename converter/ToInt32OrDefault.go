package converter

// ToInt32OrDefault 将数值类型转换为Int32，失败就返回默认值。 / converts numeric types to Int32, returns defaultValue on failure.
func (c *Converter[T]) ToInt32OrDefault(v T, defaultValue int32) int32 {
	result, err := c.ToInt32(v)
	if err != nil {
		return defaultValue
	}
	return result
}
