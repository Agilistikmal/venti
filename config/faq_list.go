package config

import "github.com/bwmarrin/discordgo"

var FAQList = []discordgo.SelectMenuOption{
	{
		Label: "Bagaimana cara order?",
		Value: "faq-how-to-order",
	},
	{
		Label: "Saya mau bertanya, dimana tempatnya?",
		Value: "faq-how-to-contact",
	},
	{
		Label: "Ada metode pembayaran apa saja?",
		Value: "faq-payment-list",
	},
	{
		Label: "Apakah aman/terpercaya?",
		Value: "faq-trusted",
	},
	{
		Label: "Cara redeem Minecraft",
		Value: "faq-how-to-redeem-minecraft",
	},
	{
		Label: "Bagaimana cara menggunakan VPN",
		Value: "faq-how-to-use-vpn",
	},
	{
		Label: "Bisa untuk POJAV Launcher?",
		Value: "faq-pojav",
	},
}

type faqResponse struct {
	Value    string
	Response string
}

var FAQResponse = []faqResponse{
	{
		Response: "Pembelian secara otomatis dapat dilakukan lewat channel <#1039148098808725624>. Pilih produk yang ingin dibeli lalu klik tombol Purchase.",
		Value:    "faq-how-to-order",
	},
	{
		Response: "Anda dapat menghubungi kami melalui channel <#1144821480002170960> atau DM Admin kami.",
		Value:    "faq-how-to-contact",
	},
	{
		Response: "Untuk pembelian melalui <#1039148098808725624> hanya bisa menggunakan **QRIS** (support semua e-Wallet/Mobile Banking). Untuk pembayaran lainnya bisa melalui [**Shopee GPrestore**](https://shopee.co.id/gprestore).",
		Value:    "faq-payment-list",
	},
	{
		Response: "Sudah terbukti aman oleh ratusan orang. Bisa cek rating dari mereka di channel <#1041685950952124446> atau [**Shopee GPrestore**](https://shopee.co.id/gprestore).",
		Value:    "faq-trusted",
	},
	{
		Response: "Video Tutorial Coming Soon...",
		Value:    "faq-how-to-redeem-minecraft",
	},
	{
		Response: `- Buka halaman website VPN yang direkomendasikan (bisa menggunakan yang lain jika punya)
					- Download VPN tersebut atau Add to browser extension
					- Buka VPN tersebut (Jika menggunakan extension, buka melalui halaman tab redeem
					- Pilih negara yang diberi tau saat pembelian
					- Lalu klik tombol Connect
					- Selesai`,
		Value: "faq-how-to-use-vpn",
	},
	{
		Response: "Bisa",
		Value:    "faq-pojav",
	},
}
