package main

import (
	"context"
	"flag"
	"fmt"
	modulbank_go "github.com/DarthRamone/modulbank-go"
	"net/http"
	"time"
)

var merchant = flag.String("merchant", "", "")
var secretKey = flag.String("secret_key", "", "")

func main() {
	flag.Parse()

	r := modulbank_go.CreateBillRequest{
		Merchant:       *merchant,
		Amount:         1000,
		Description:    "Заказ №14425840",
		Testing:        true,
		ReceiptContact: "test@mail.com",
		UnixTimestamp:  time.Now().Unix(),
		Salt:           "dPUTLtbMfcTGzkaBnGtseKlcQymCLrYI",
		CustomOrderId:  "sdfsdf 1",
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
	opts := modulbank_go.MerchantOptions{
		Merchant:  *merchant,
		SecretKey: *secretKey,
	}
	bill, err := modulbank_go.CreateBill(context.Background(), r, opts, http.DefaultClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(bill.Url)
	//
	//gbr := modulbank_go.GetBillRequest{
	//	Id:            bill.Id,
	//	Merchant:      opts.Merchant,
	//	UnixTimestamp: time.Now().Unix(),
	//	Salt:          "dPUTLtbMfcTGzkaBnGtseKlcQymCLrYI",
	//}

	b2, err := modulbank_go.GetBill(context.Background(), bill.Id, opts, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	fmt.Println(b2.Url)
}
