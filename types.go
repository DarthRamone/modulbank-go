package modulbank_go

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	Money        float64
	IntBool      bool
	ReceiptItems []ReceiptItem
)

func (m Money) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%.2f\"", m)
	formatted = strings.Replace(formatted, ",", ".", -1)
	return []byte(formatted), nil
}

func (b IntBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (b *IntBool) UnmarshalJSON(bytes []byte) error {
	s := string(bytes)
	fmt.Printf("BYTES: %s\n", s)
	s = strings.Trim(s, "\"")
	var result IntBool = false
	if s == "1" {
		result = true
	} else {
		result = false
	}
	b = &result
	return nil
}

func (ri ReceiptItems) MarshalJSON() ([]byte, error) {
	slice := []ReceiptItem(ri)
	bytes, err := json.Marshal(slice)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal slice: %w", err)
	}
	str := string(bytes)
	str = strings.Replace(str, "\"", "\\\"", -1)
	str = fmt.Sprintf("\"%s\"", str)
	return []byte(str), nil
}
