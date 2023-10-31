package service

import (
	"errors"
	"github.com/Agilistikmal/venti/helper"
	"github.com/Agilistikmal/venti/model"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
	"github.com/xendit/xendit-go/retailoutlet"
	"github.com/xendit/xendit-go/transaction"
	"time"
)

func CreatePayment(product model.Product, quantity int, voucherCode string, email string) (*xendit.QRCode, *xendit.RetailOutlet, *model.Voucher, error) {
	var voucher model.Voucher
	var discount int
	voucher, err := FindVoucherByCode(voucherCode)
	if err != nil {
		discount = 0
		voucher = model.Voucher{
			Code:      "INVALID",
			Discount:  0,
			ExpiredAt: time.Now(),
		}
	} else {
		discount = voucher.Discount * quantity
	}

	externalId := "GPR-" + helper.RandomString(8)
	totalAmount := float64((quantity * product.Price) - (quantity * discount))
	qr, _ := qrcode.CreateQRCode(&qrcode.CreateQRCodeParams{
		ExternalID:  externalId,
		Type:        xendit.DynamicQRCode,
		Amount:      totalAmount,
		CallbackURL: "https://safatanc.com/api/callback",
	})

	var alfamart = &xendit.RetailOutlet{
		PaymentCode: "Pembayaran Alfamart hanya bisa untuk nominal Rp10000 keatas",
	}
	if totalAmount >= 10000 {
		alfamart, _ = retailoutlet.CreateFixedPaymentCode(&retailoutlet.CreateFixedPaymentCodeParams{
			ExternalID:       externalId,
			RetailOutletName: xendit.RetailOutletNameAlfamart,
			Name:             email,
			ExpectedAmount:   totalAmount,
		})
		return qr, alfamart, &voucher, nil
	}
	return qr, alfamart, &voucher, nil
}

func GetPayment(externalId string) ([]xendit.Transaction, error) {
	trx, err := transaction.GetListTransaction(&transaction.GetListTransactionParams{
		ReferenceID: externalId,
	})
	if err != nil {
		return trx.Data, err
	}
	if len(trx.Data) == 0 {
		return trx.Data, errors.New("Belum dibayar")
	}
	return trx.Data, nil
}
