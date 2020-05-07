package modulbank_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	API interface {
		CreateBill(ctx context.Context, request CreateBillRequest) (Bill, error)
	}

	client struct {
		http *http.Client
		opts MerchantOptions
	}

	MerchantOptions struct {
		Merchant  string
		SecretKey string
	}
)

func NewClient(opts MerchantOptions) API {
	return client{http: http.DefaultClient, opts: opts}
}

func GetBill(ctx context.Context, billId string, opts MerchantOptions, c *http.Client) (bill Bill, err error) {
	request := GetBillRequest{
		Id:            billId,
		Merchant:      opts.Merchant,
		UnixTimestamp: time.Now().Unix(),
		Salt:          uuid.New().String(),
		Signature:     "",
	}

	var signature string
	signature, err = getSignature(opts.SecretKey, request)
	if err != nil {
		return bill, fmt.Errorf("failed to calculate signature")
	}

	request.Signature = signature

	requestStr := fmt.Sprintf(
		"https://pay.modulbank.ru/api/v1/bill?id=%s&merchant=%s&unix_timestamp=%d&signature=%s&salt=%s",
		request.Id, request.Merchant, request.UnixTimestamp, request.Signature, request.Salt)

	fmt.Printf("requestStr: %s\n", requestStr)

	req, err := http.NewRequest("GET", requestStr, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
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
		return bill, fmt.Errorf("status not ok: %s", billResponse.Status)
	}

	return billResponse.Bill, nil
}

func CreateBill(ctx context.Context, request CreateBillRequest, opts MerchantOptions, c *http.Client) (bill Bill, err error) {
	request.Merchant = opts.Merchant

	sign, err := getSignature(opts.SecretKey, request)
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

	resp, err := c.Do(req)
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
		return bill, fmt.Errorf("status not ok: %s", billResponse.Status)
	}

	return billResponse.Bill, nil
}

func (c client) CreateBill(ctx context.Context, request CreateBillRequest) (bill Bill, err error) {
	return CreateBill(ctx, request, c.opts, c.http)
}
