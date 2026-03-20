package converter

// ToInt64OrDefault 将数值类型转换为Int64，失败就返回默认值。 / converts numeric types to Int64, returns defaultValue on failure.
func (c *Converter[T]) ToInt64OrDefault(v T, defaultValue int64) int64 {
	result, err := c.ToInt64(v)
	if err != nil {
		return defaultValue
	}
	return result
}
