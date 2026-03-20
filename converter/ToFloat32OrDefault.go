package converter

// ToFloat32OrDefault 将数值类型转换为float32，失败就返回默认值。 / converts numeric types to float32, returns defaultValue on failure.
func (c *Converter[T]) ToFloat32OrDefault(v T, defaultValue float32) float32 {
	result, err := c.ToFloat32(v)
	if err != nil {
		return defaultValue
	}
	return result
}
