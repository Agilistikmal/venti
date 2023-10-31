package product_event

import (
	"fmt"
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/helper/component_helper"
	"github.com/Agilistikmal/venti/helper/embed_helper"
	"github.com/Agilistikmal/venti/service"
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
	"strconv"
	"strings"
)

func PurchaseModalSubmit(bot *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionModalSubmit {
		return
	}
	data := interaction.ModalSubmitData()
	if strings.Contains(data.CustomID, "purchase-confirm") == false {
		return
	}

	err := bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Berhasil, silahkan cek DM anda.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	handler.HandleError(err)

	customId := strings.Split(data.CustomID, "|")
	productId := customId[1]
	product, err := service.FindProductById(productId)
	handler.HandleError(err)
	email := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	var quantity int
	quantity, err = strconv.Atoi(data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value)
	if err != nil || quantity <= 0 {
		quantity = 1
	}
	if quantity > product.Stock {
		quantity = product.Stock
	}
	notes := data.Components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	if notes == "" {
		notes = "-"
	}
	voucherCode := data.Components[3].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	qris, alfamart, voucher, err := service.CreatePayment(product, quantity, voucherCode, email)
	qrFilename := service.GenerateQRCode(qris.QRString)

	file, err := os.Open(qrFilename)
	handler.HandleError(err)

	userDm, _ := bot.UserChannelCreate(interaction.Member.User.ID)

	_, err = bot.ChannelMessageSendComplex(userDm.ID, &discordgo.MessageSend{
		File: &discordgo.File{
			Name:        "qris.png",
			ContentType: "image/png",
			Reader:      io.Reader(file),
		},
		Embed: embed_helper.CustomWithImage(
			interaction.Member.User,
			"Purchase "+qris.ExternalID,
			email+"\nBerikut ini adalah pesanan anda. Harap segera melakukan pembayaran QRIS atau Alfamart dibawah ini\nKlik tombol Sudah Bayar setelah melakukan pembayaran.\n\n- QRIS dapat di scan melalui semua e-wallet dan mobile banking.\n- Kode Alfamart berlaku untuk gerai Alfamart/Alfamidi/Lawson/DAN+DAN",
			[]*discordgo.MessageEmbedField{
				embed_helper.CreateField("Item", fmt.Sprintf("``` 📦 %dx %s ```", quantity, product.Name), false),
				embed_helper.CreateField("Voucher", fmt.Sprintf("``` 🎫 %s (%dx -Rp%d) ```", voucher.Code, quantity, voucher.Discount), true),
				embed_helper.CreateField("Total Price", fmt.Sprintf("``` 🪙 Rp%d ```", (quantity*product.Price)-(quantity*voucher.Discount)), true),
				embed_helper.CreateField("Catatan", notes, false),
				embed_helper.CreateField("Kode Pembayaran Alfamart", fmt.Sprintf("```js\n %s ```\n_*Katakan ke kasir ingin melakukan pembayaran ke **SAFATANC TECHNOLOGY** dan tunjukkan **Kode Pembayaran** diatas._", alfamart.PaymentCode), false),
			},
			"attachment://qris.png",
			embed_helper.WHITE,
		),
		Components: component_helper.CreateButton("Sudah Bayar", discordgo.SecondaryButton, "confirm-payment"),
	})
	handler.HandleError(err)

	file.Close()
	err = os.Remove(qrFilename)
	handler.HandleError(err)
}
