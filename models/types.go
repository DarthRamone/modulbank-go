package models

import (
	"fmt"
	"strings"
)

type (
	Money   float32
	IntBool bool
)

func (m Money) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("%.0f", m)
	formatted = strings.Replace(formatted, ",", ".", -1)
	return []byte(formatted), nil
}

func (b IntBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}
