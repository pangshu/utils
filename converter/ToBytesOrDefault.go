package converter

// ToBytesOrDefault 将数值类型转换为[]byte，失败就返回默认值。 / converts numeric types to []byte, returns defaultValue on failure.
func (c *Converter[T]) ToBytesOrDefault(v T, defaultValue []byte) []byte {
	result, err := c.ToBytes(v)
	if err != nil {
		return defaultValue
	}
	return result
}
