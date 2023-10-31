package service

import (
	"fmt"
	"github.com/Agilistikmal/venti/database"
	"github.com/Agilistikmal/venti/model"
)

func CreateVoucher(request *model.Voucher) *model.Voucher {
	database.DB.Create(&request)
	return request
}

func UpdateVoucher(code string, request *model.Voucher) (*model.Voucher, error) {
	_, err := FindVoucherByCode(code)
	if err != nil {
		return nil, err
	}
	var voucher *model.Voucher
	database.DB.First(voucher, "code = ?", code)

	voucher.Code = request.Code
	voucher.Discount = request.Discount

	database.DB.Save(&voucher)
	return voucher, nil
}

func DeleteVoucher(code string) error {
	_, err := FindVoucherByCode(code)
	if err != nil {
		return err
	}
	var voucher *model.Voucher
	database.DB.Where("code = ?", code).Delete(&voucher)
	return nil
}

func FindAllVoucher() []model.Voucher {
	var vouchers []model.Voucher
	database.DB.Find(&vouchers)
	return vouchers
}

func FindVoucherByCode(code string) (model.Voucher, error) {
	if code == "" {
		return model.Voucher{}, fmt.Errorf("invalid voucher")
	}
	var voucher model.Voucher
	result := database.DB.First(&voucher, "code = ?", code)
	if result.Error != nil {
		return model.Voucher{}, fmt.Errorf("voucher with code %s not found", code)
	}
	return voucher, nil
}
