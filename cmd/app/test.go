package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/DarthRamone/modulbank-go"
	"time"
)

var merchant = flag.String("merchant", "", "")
var secretKey = flag.String("secret_key", "", "")

func main() {
	flag.Parse()

	fmt.Printf("Merchant: %s\nSecretKey: %s\n\n", *merchant, *secretKey)

	r := modulbank_go.BillRequest{
		Merchant:       *merchant,
		Amount:         1000,
		Description:    "Заказ №14425840",
		Testing:        true,
		ReceiptContact: "test@mail.com",
		UnixTimestamp:  time.Now().Unix(),
		Salt:           "dPUTLtbMfcTGzkaBnGtseKlcQymCLrYI",
		ReceiptItems: []modulbank_go.ReceiptItem{
			{
				Name:          "Товар 1",
				Quantity:      1,
				Price:         500,
				SNO:           modulbank_go.SNOOsn,
				PaymentObject: modulbank_go.PaymentObjectService,
				PaymentMethod: modulbank_go.PaymentMethodFullPrepayment,
				VAT:           modulbank_go.VATNone,
			},
			{
				Name:          "Товар 2",
				Quantity:      1,
				Price:         500,
				SNO:           modulbank_go.SNOOsn,
				PaymentObject: modulbank_go.PaymentObjectService,
				PaymentMethod: modulbank_go.PaymentMethodFullPrepayment,
				VAT:           modulbank_go.VATNone,
			},
		},
		SendLetter: false,
	}

	client := modulbank_go.NewClient(*secretKey)
	bill, err := client.CreateBill(context.Background(), r)

	if err != nil {
		panic(err)
	}

	fmt.Println(bill)
}
