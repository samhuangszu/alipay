package alipay_test

import (
	"testing"

	"github.com/samhuangszu/alipay"
)

var (
	appID     = "2018070460555345"
	partnerID = "2088102169227503"

	// RSA2(SHA256)
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmZB8N5J3zxakNtxZ0vw3v3uLDW1cVC2sqH+G4YioeMGAjAHbxBYj+LKPTvS6GWOy8+lk+y3Kpg1F+z68ggl4vqYATiugFkppZihN0Gv1DeHoWWhH7uRAQJ4zAx6HO3NJro6bZrH8Co1Mwt9eb0NqmFWt1B1PJ6vIyBUF82FtMWO09WUS/NdIEHuTRIANyfqQqkyy32wApdQSonItHoQ5ThvEYpRMGx8bc1fuYnnF9vawPc5ipieg7uHYw+RRVx6Qb3WOUtP9xKrwYlazT1VUSdfNvMRzBk9kHMxRpp8kj3embvxcN4BrkmP9FTCb8hZopGsFdKcc74g96EYkVerrfQIDAQAB"

	appPrivateKey = "MIIEowIBAAKCAQEA08OVIKodtdDKgWtnjHWjlEsLr/yQRSZAOqZ/iw6pZM++NetOzClR8/pl9xRRTMvp9/bosvvXtS2es6pVA3tJKKv4rZpPgVfDnkMvmwhe9uqW2bHwKOtI18pfCGtcrdAqjfiSFWHPcBwOfzpy77fehzJbHNAJYRy11d1pN1xSHI/+idvNiPc05gwyhQhyzQg4fxAfgLn/KD7D1cfYwSfFHhSajoXhe7QZiXSNZtWi0X6LpAnJXBTYdc+z9NSsGxpFSybSNFzFjM/mtE04R0CEUj8t8kMRm3ylZh02R5SoIoOpYBrtwItzMtxzO0Bpsh1H2wIX6pdUI1qjb+n4s8+UNwIDAQABAoIBAHWmAWHmYR8z4m3Id/znFnw6vUGGtbhVKfXrroxZFu81I13dXye6BRfhE2kT1p8t4syqVlmp5Qt5TuROS3Dlu51m/X+lbWpUkg+JEBe3rNrNNgW3/88b+Jo6ilWpIBJIxnudj7tlXdPRIpS4qkZYBr9h0JDPW4aiZrtsiRoCpnPK4EA1Qm667KMCyK8az7yjMWZzAnFtc9slhmsTC/hEVc1mD3xj8hLcAbrUCGIh6ZhuvV2mVWToH56LNzaqCvMLrGh+XmekFbKcA7FdcxqrKPOJAOH6tiC2MJbc75ZrFn3oaHLoDNmE7uH9V/61No8D3IJJCno5OKKzH8O6OYmgRMECgYEA/tw4AHGeU3WiuT57Yd31rKyrqcv6QvlliXFr5nd7rWPennLo5JAUZ6tK7eVkdmH1nYrrCROmAkfiCJ8fzdzl5Whi93GNcNk3lfcOuDh5mWUAdbdf4OsMB1l5nlbzemta5XTJ7KYjzjPtba1KcR4XZV2uwQ1OTa6z3LqBtZjyrAcCgYEA1LYGN5G74c4gLzKHuw+JBNQXTxpDh6SUn0+xEiSoogecBusd/Xoroun59UwTDtxYRtdTZ/4XFesHfSt7vJW1XqlGGu0XUbrExww8oG2C5vE5kesEi2+RwKYS2XWuH5N6OXDVEBxDOI6FF5wgfF3IWjogRkpWtzPQHYzLzHNCKlECgYEArHbf2sUITBgV5t92ZRLr1k7+16d2Em5snKbJSqteYUZs9rJyEYHCnSjYSsZpxoahzFuek4TdWTvFpOnxfsDWPfj2x9XqzgvgRrZqGpX1C63CsuZFlpHYkPymhVT37MbLOu1eW6tOHZMcP5T+BDBFys2rnW5gp8bqZhs7/WSkC28CgYBr2M5WJmGoHyZaR7hhs4K5G/+lb5+FioCANZuFo2iMnmcRauwNtH+jXhYJtMSE6Fspr8ruEvoJdtyZLg7SkuUeVZzh0gvcuGFEuoOSo3OXO/8AaLQpLiwsDQfJkczy3Yc+0GJ2hJ3gcwIv8kzWQYZvmMD8YnyM7ow6L0S1KEsHwQKBgCYi7gz9UKLMsgmbBDb14AgOzAbZ1F79UgFTZFQhS2KixWC2jCRYG3aLTOzbZ9AdTyrvTGyovZcwojhEWNM8ZGiEnza+9n0EUQ05SQZ5MKBC6/a8GxNaSnef6d2Q/NXmPZSgYKKMWRPBMJYfU0EsQBKNaJOf8GaS7zwwxOqd7WWZ"
	appPublicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA08OVIKodtdDKgWtnjHWjlEsLr/yQRSZAOqZ/iw6pZM++NetOzClR8/pl9xRRTMvp9/bosvvXtS2es6pVA3tJKKv4rZpPgVfDnkMvmwhe9uqW2bHwKOtI18pfCGtcrdAqjfiSFWHPcBwOfzpy77fehzJbHNAJYRy11d1pN1xSHI/+idvNiPc05gwyhQhyzQg4fxAfgLn/KD7D1cfYwSfFHhSajoXhe7QZiXSNZtWi0X6LpAnJXBTYdc+z9NSsGxpFSybSNFzFjM/mtE04R0CEUj8t8kMRm3ylZh02R5SoIoOpYBrtwItzMtxzO0Bpsh1H2wIX6pdUI1qjb+n4s8+UNwIDAQAB"
)

var client *alipay.AliPay

func init() {
	client = alipay.New(appID, partnerID, aliPublicKey, appPrivateKey, appPublicKey, false)
}

func TestAliPay_TradeCreate(t *testing.T) {
	// t.Log("========== TradeWapPay ==========")
	var p = alipay.AliPayTradeCreate{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	// p.ReturnURL = "http://203.86.24.181:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_20180623021"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	// p.BuyerId = "2088102147948060"
	data, err := client.TradeCreate(p)
	t.Log(data)
	t.Error(data)
	if err != nil {
		t.FailNow()
	}
}
