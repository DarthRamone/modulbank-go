package modulbank_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	API interface {
		CreateBill(ctx context.Context, request BillRequest) (Bill, error)
	}

	client struct {
		http *http.Client
		opts ClientOptions
	}

	ClientOptions struct {
		Merchant  string
		SecretKey string
	}
)

func NewClient(opts ClientOptions) API {
	return client{http: http.DefaultClient, opts: opts}
}

func (c client) CreateBill(ctx context.Context, request BillRequest) (bill Bill, err error) {
	request.Merchant = c.opts.Merchant

	sign, err := getSignature(c.opts.SecretKey, request)
	if err != nil {
		return bill, fmt.Errorf("failed to generate signature: %w", err)
	}

	request.Signature = sign

	jsonStr, err := json.Marshal(request)
	if err != nil {
		return bill, fmt.Errorf("failed to marshal request obj: %w", err)
	}

	req, err := http.NewRequest("POST", "https://pay.modulbank.ru/api/v1/bill", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return bill, fmt.Errorf("failed to do http request: %w", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Printf("failed to close request body: %w", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bill, fmt.Errorf("failed to read body: %w", err)
	}

	fmt.Println(string(body))

	var billResponse billResponse
	err = json.Unmarshal(body, &billResponse)
	if err != nil {
		return bill, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	if billResponse.Status != "ok" {
		return bill, fmt.Errorf("status not ok: %w", err)
	}

	return billResponse.Bill, nil
}
