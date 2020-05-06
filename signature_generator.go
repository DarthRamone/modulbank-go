package modulbank_go

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//Calculates signature for bill creation request
func getSignature(secretKey string, request BillRequest) (string, error) {
	request.Merchant = "25787de5-4f38-4e4a-939d-0de20a6e0698"

	//Future request payload
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	//Key value object represents future JSON structure
	kv := make(map[string]interface{})
	err = json.Unmarshal(requestBytes, &kv)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal request bytes to map: %w", err)
	}

	//Sorting all JSON fields by key
	keys := make([]string, 0)
	for k, _ := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//Retrieving all key-value pairs with base64 encoded value
	pairs := make([]string, 0)
	for _, key := range keys {
		valueBytes, err := json.Marshal(kv[key])
		if err != nil {
			return "", fmt.Errorf("failed to marshal map value: %w", err)
		}

		valueStr := string(valueBytes)

		//As long as receipt_items object is passed as string(not as object) - all quotes is escaped in marshaller
		//For signature generating we should unescape them
		valueStr = strings.Replace(valueStr, "\\\"", "\"", -1)
		//Also we should trim all leading and trailing quotes
		valueStr = strings.Trim(valueStr, "\"")

		//Skip all empty and unnecessary fields
		if valueStr == "" || key == "signature" {
			continue
		}

		//Back to bytes
		valueBytes = []byte(valueStr)

		//Encoding value to base64
		value := make([]byte, base64.StdEncoding.EncodedLen(len(valueBytes)))
		base64.StdEncoding.Encode(value, valueBytes)

		//Building encoded key-value pair
		pair := fmt.Sprintf("%s=%s", key, string(value))
		pairs = append(pairs, pair)
	}

	//Combine all pairs all together
	combinedPairs := strings.Join(pairs, "&")

	//Calculating signature following formula: SHA1(secret_key + SHA1(secret_key + values)
	str1 := secretKey + combinedPairs
	hash1 := sha1.Sum([]byte(str1))
	str2 := secretKey + fmt.Sprintf("%x", hash1[:])
	resultHash := sha1.Sum([]byte(str2))

	result := fmt.Sprintf("%x", resultHash[:])

	return result, nil
}
