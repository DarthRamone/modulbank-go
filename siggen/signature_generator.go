package siggen

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/DarthRamone/modulbank-go/models"
	"sort"
	"strings"
)

func GetSignature(ctx context.Context, secretKey string, request models.BillRequest) (string, error) {
	kv := make(map[string]interface{})

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	err = json.Unmarshal(requestBytes, &kv)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal request bytes to map: %w", err)
	}

	keys := make([]string, 0)
	for k, _ := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pairs := make([]string, 0)
	for _, key := range keys {
		var valueBytes []byte
		var err error

		if key == "receipt_items" {
			valueBytes, err = json.Marshal(request.ReceiptItems)
			fmt.Printf("receipt bytes: %s\n\n", valueBytes)
			//valueBytes = []byte("[{\"discount_sum\": 40, \"name\": \"Товар 1\", \"payment_method\": \"full_prepayment\", \"payment_object\": \"commodity\", \"price\": 48, \"quantity\": 10, \"sno\": \"osn\", \"vat\": \"vat10\"}, {\"name\": \"Товар 2\", \"payment_method\": \"full_prepayment\", \"payment_object\": \"commodity\", \"price\": 533, \"quantity\": 1, \"sno\": \"osn\", \"vat\": \"vat10\"}]")
		} else {
			valueBytes, err = json.Marshal(kv[key])
		}
		if err != nil {
			return "", fmt.Errorf("failed to marshal map value: %w", err)
		}

		valueStr := string(valueBytes)
		valueStr = strings.Trim(valueStr, "\"")

		if valueStr == "" || key == "signature" {
			continue
		}

		valueBytes = []byte(valueStr)

		value := make([]byte, base64.StdEncoding.EncodedLen(len(valueBytes)))
		base64.StdEncoding.Encode(value, valueBytes)

		pair := fmt.Sprintf("%s=%s", key, string(value))
		pairs = append(pairs, pair)
	}

	combinedPairs := strings.Join(pairs, "&")
	combinedPairs = strings.Replace(combinedPairs, "\"", "", -1)

	fmt.Printf("%s\n\n", combinedPairs)

	//SHA1(secret_key + SHA1(secret_key + values)
	str1 := secretKey + combinedPairs
	hash1 := sha1.Sum([]byte(str1))
	str2 := secretKey + fmt.Sprintf("%x", hash1[:])
	resultHash := sha1.Sum([]byte(str2))

	result := fmt.Sprintf("%x", resultHash[:])

	return result, nil
}
