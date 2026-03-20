package converter

// ToIntOrDefault 将数值类型转换为Int，失败就返回默认值。 / converts numeric types to Int, returns defaultValue on failure.
func (c *Converter[T]) ToIntOrDefault(v T, defaultValue int) int {
	result, err := c.ToInt(v)
	if err != nil {
		return defaultValue
	}
	return result
}
