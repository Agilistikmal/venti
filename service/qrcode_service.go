package service

import (
	"encoding/base32"
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/helper"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(code string) string {
	qrName := helper.RandomString(4)
	err := qrcode.WriteFile(code, qrcode.Low, 256, "temp/"+qrName+".png")
	handler.HandleError(err)
	return "temp/" + qrName + ".png"
}

func GenerateQRCodeBase64(code string) string {
	data, err := qrcode.Encode(code, qrcode.Low, 256)
	handler.HandleError(err)
	encoded := base32.StdEncoding.EncodeToString(data)
	return "data:image/png;base64," + encoded
}
