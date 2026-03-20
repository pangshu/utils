package converter

// ToChannel converts a slice to a read-only channel.
func (c *Converter[T]) ToChannel(v []T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range v {
			ch <- item
		}
	}()
	return ch
}
