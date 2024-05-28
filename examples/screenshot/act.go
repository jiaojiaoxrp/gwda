package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/http2"
)

func main() {

	// 生成当前时间戳和nonce
	timestamp := time.Now().Unix()
	nonce := timestamp

	// 将时间戳转换为格式化字符串
	timeFormat := "20060102T150405Z"
	currentTime := time.Now().UTC().Format(timeFormat)

	url := "https://flex-capacity-na.amazon.com/AcceptOffer"
	data := `{"offerId":"Ok9mZmVySWQuRW5jcnlwdGlvbktleS02aGlZbDQAAACDDtWCQwmuLr5uKZbInTe3foh3t7qLFJ9Ixyqr4JddvRTzgcemZ59HlwxdDuqk3nNaxkr1rvzzP3b3PyA6KgD0pfDSC55Wrh5p4/pDRAHLSB5VfBw/Ttpg7LCARTeuhoJxD7kJbZGA82LKwui8xugleoCniG3YZgqLjKBK9j70PDPxumcbtZzvUufXixh6axyouCsYdJnK30gYtD6vztteUAXbtGMFEuy6vMjqpYghplo+dwopMUb6yGJSgdpkWmKD0vtEli6BSRo0ljPOEut0GkzeNZVpJcVsd6Ms0Hpvsc/wS1U+NtbL/+4KR/XU5Tt7Hes9eTVaR9EOGfldonFWLYB/J1xOFGTtr+q5UwU/J2ot3EdHbBbadNocvKlnDlStyUoOA0zTWxlrMA1+tCt7+IXfBFTiuMz7Q9QLVIDTJ9fWTqzCyU1fuT7b85QCYTAuO50q68CQ2iNHjYxo3P/PGswmKPkmnAslWulEczVLs/zaIUw4+hsC7YzML2vn1lTqhkroQet1tWUsLmwE5ILJXe7m95M+ReIz9+vXwmRpX590kSC6C14=|aPpKv5kJPxTHBFqHlM/eXBU65aLGLFUt8fvdttFFLp0="}`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	token := "Atna|EwICINUFmbX4YMaco6noBVXEjyprrL5EEt1ib2rJ82KyCcxAuOhjIy0bNtM6vhK0ISzMxFWJCGJ3DukVJ4QBDwR1qPE98BbEVz0Gbr-8ll_jZV_bSAeTXat4Yz2udzI5fuc6i6-by83lHOQ_gwDKmBjPLTMFiM6Yfa-uXnWyDAaabAIbR6vL3yNTeV6yZZSb3I0hyHUMS92IOd64K0EN0Mzr9nvpf3QpSglhRgzi8Ji5pX7DDOhVuD7fQfiReOLphccrXXgl-Pg5LeVI5Csa0pXYtcSGzGyIz2lrec5ny9to74eCFfCZi41mA5CeKU0vEy-9Iaj7RQK547VCqAYOyUgMtW0FhQI7CnhM0TAqgaoXSxJwZg"
	req.Header.Set("Host", "flex-capacity-na.amazon.com")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-amz-access-token", token)
	req.Header.Set("Authorization", "RABBIT3-HMAC-SHA256 SignedHeaders=x-amz-access-token;x-amz-date,Signature=058bb18c4f07ff7ee43abd99383b3781d3eaf017587f9f9ed966d9068794be51")
	req.Header.Set("X-Amz-Date", currentTime)
	req.Header.Set("x-flex-instance-id", "A30D1951-2749-4BD4-9B37-4D540991A468")
	req.Header.Set("Signature-Input", "x-amzn-attest=(\"@path\" \"x-amzn-marketplace-id\" \"user-agent\");created="+strconv.Itoa(int(nonce))+";nonce="+strconv.Itoa(int(nonce))+";alg=\"apple-attest\";keyid=\"AQICAHg21E9N4cqPKrpywXBTFl2Ivrk5dxuqSkVKL5WYaXPSMgHFrzNT7/afBqmH7vbkCQSpAAABLTCCASkGCSqGSIb3DQEHBqCCARowggEWAgEAMIIBDwYJKoZIhvcNAQcBMB4GCWCGSAFlAwQBLjARBAx3npftO9KcMfrV/RQCARCAgeHfZ75KfNrXTEQJF8/SiCDqKEq99/lh4etxx8y6nTuhnG7nr2F4b+X2FFF8TXpXVxVpCh9bchHyoU7wcTbUs+0iJ+fpFZ+u1KiopwtDpBXXUM1eCWalLGpgCdqHXopQ5tFly8eA5o7jay48kMGdSv+ovUWtt5sQISrR0jPq+iMiE2tIbOGEvCJrP+h1m3e5aT6C11ccUJ+TRKFkvFtA0WDQ/M/PZ6F/d8P6/jwBm1zFIpfScrL/IYRXNM+Ytc5ZWF/xqL3dk5MrtuYBLbNopDxRcCSCJcEDcEl6mYVVsGnrVDM=")
	req.Header.Set("Signature", "x-amzn-attest=:omlzaWduYXR1cmVYSDBGAiEA9mFL9h8MPkvn+Yd4A3raZMeZjCcXVk/SQiDlXzrJzisCIQD7+PawHbMPItft+3JSwYwDzRWmyzxzyUt+iTxjCJ9Mw3FhdXRoZW50aWNhdG9yRGF0YVgl7fPChS0bmEM9b1O14jz2TkRRyM28WLBf6EIHxvfTZflAAAAVww==:")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "iOS/17.5 (iPhone Darwin) Model/iPhone Platform/iPhone15,3 RabbitiOS/2.130.1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", `at-main="Atza|IwEBIGO3ZYOMO63M2i4K6KzBiMQCVCdD6VZaiiqKQfDCr0RvpI9R1mhObTYmsCFe64nJaWfPDRoLkenA0yxb-zYYRqIMm94DVzqMM6DEfnXzN9sfmmxlyu8EOa_nom1d4OQNVoVp9wp6Opa3M9i-vEWt2DXJmjPKnCyodVFmbwguQBopmsVn07CXkYAt4zBdYLKbelGaICqt1Sn28fRLJY9OvPjUgua1RKlUrHGUJKUDBpP8GwFF6KwcMtI-ZSnB8K7xeL71IatueRKIYNfcKV_S9qXddfxASXXlZLoRID61rS2Wtl18qrhORoD3QgiI2Vw_FSriC2lOI2MjsFNOCxzzduD6"; sess-at-main="LZT3d1Qs6HEp1+yFKT0My0B+eEeDIj84thkv7I6SnEc="; session-id=139-1933813-2345366; session-id-time=2347566371l; session-token=FMIqgUaquZiPZOsL+i6gwHWBj0ovYn1k0pu37FsZEQKs/lLiTjw6uYAZgRZJUcI43/J8FFT7pKpQPUNhZC7T1tiMAD/h1G6dTaKHDNPdrSfRVq5yoJpTMAPpWy0wD5In7DhRaztyo//16YkrArs/Rk5XBUovHOKYzsojdNSZqRK23zWdB3kp8Naf8vJ6M0Zzsdf+jQPgnZ0eOkxqs/WIaTgkP3oypXGRdBXm7lmfR3HRX7tW4SnwtIdyGcsWEL6Lq8pokE5q0scX4gqwvWfLxql4zrpHD87/EjU7Us/Dx8ASHIgxwud7b6GHAx6t6D/L1qiTdvmfUP9nwg9631YtpxnU4XwDhTigxnGddtMuD1mf2kJ/+COho4fzdQy8xejV; ubid-main=130-0626653-7877250; x-main="X6ulEY2dvj9C4dfRCm9j4@QWnoxuGKHmcRrhZ9Cw1Q6Nn?veVfvbMnbM1n3AGgGe"; lc-main=en_US; i18n-prefs=USD; sp-cdn=J4F7`)
	req.Header.Set("x-amzn-marketplace-id", "ATVPDKIKX0DER")

	// 创建自定义的Transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: "flex-capacity-na.amazon.com",
		},
	}

	// 启用HTTP/2
	http2.ConfigureTransport(tr)

	// 创建HTTP客户端
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
