package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/DarthRamone/modulbank-go/models"
	"github.com/DarthRamone/modulbank-go/siggen"
	"io/ioutil"
	"net/http"
)

var merchant = flag.String("merchant", "", "")
var secretKey = flag.String("secret_key", "", "")

func main() {
	flag.Parse()

	fmt.Printf("Merchant: %s\nSecretKey: %s\n\n", *merchant, *secretKey)

	r := models.BillRequest{
		Merchant:       *merchant,
		Amount:         973,
		OrderId:        "14425840",
		Description:    "Заказ №14425840",
		Testing:        true,
		ClientEmail:    "test@test.ru",
		ClientPhone:    "+7 912 9876543",
		SuccessUrl:     "http://myawesomesite.com/payment_success",
		ReceiptContact: "test@mail.com",
		ReceiptItems: []models.ReceiptItem{
			{
				Name:          "Товар 1",
				Quantity:      10,
				Price:         48,
				DiscountSum:   40,
				SNO:           models.OSN,
				PaymentObject: models.COMMODITY,
				PaymentMethod: models.FULL_PREPAYMENT,
				VAT:           models.VAT10,
			},
			{
				Name:          "Товар 2",
				Quantity:      1,
				Price:         533,
				DiscountSum:   0,
				SNO:           models.OSN,
				PaymentObject: models.COMMODITY,
				PaymentMethod: models.FULL_PREPAYMENT,
				VAT:           models.VAT10,
			},
		},
		UnixTimestamp: 1573451160,
		Salt:          "dPUTLtbMfcTGzkaBnGtseKlcQymCLrYI",
		SendLetter:    false,
	}

	sign, err := siggen.GetSignature(context.Background(), *secretKey, r)
	if err != nil {
		panic(err)
	}

	fmt.Println(sign)

	r.Signature = sign

	jsonStr, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nrequest: %s\n\n", jsonStr)

	req, err := http.NewRequest("POST", "https://pay.modulbank.ru/api/v1/bill/", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
