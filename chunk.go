package chunk

// Index - index
type Index struct {
	From, To int
}

// GetIndex - returns index of specified size length
func GetIndex(length int, chunkSize int) <-chan Index {
	ch := make(chan Index)

	go func() {
		defer close(ch)

		for i := 0; i < length; i += chunkSize {
			idx := Index{i, i + chunkSize}
			if length < idx.To {
				idx.To = length
			}
			ch <- idx
		}
	}()

	return ch
}

// Chunking - chunking the slice
func Chunking[T any](slice []T, chunkSize int) (chunked [][]T) {
	for idx := range GetIndex(len(slice), chunkSize) {
		chunked = append(chunked, slice[idx.From:idx.To])
	}
	return
}
