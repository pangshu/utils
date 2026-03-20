package converter

// ToFloat64OrDefault 将数值类型转换为float64，失败就返回默认值。 / converts numeric types to float64, returns defaultValue on failure.
func (c *Converter[T]) ToFloat64OrDefault(v T, defaultValue float64) float64 {
	result, err := c.ToFloat64(v)
	if err != nil {
		return defaultValue
	}
	return result
}
