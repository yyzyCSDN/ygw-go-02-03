package stock

import (
	"fmt"
)

func Window[T any](values []T, start, count int) ([]T, error) {
	if start < 0 || count < 0 {
		return nil, fmt.Errorf("start and count cannot be negative")
	}
	if start >= len(values) || count == 0 {
		return []T{}, nil
	}
	end := start + count - 1
	if end > len(values) {
		end = len(values)
	}
	return append([]T(nil), values[start:end]...), nil
}
