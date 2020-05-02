package modulbank

import (
	"context"
	"github.com/DarthRamone/modulbank-go/models"
)

type (
	ClientOptions struct {
		//Идентификатор магазина, который выдается в личном кабинете на этапе интеграции.
		Merchant  string
		SecretKey string
	}

	API interface {
		CreateBill(ctx context.Context, request models.BillRequest) (string, error)
	}
)
