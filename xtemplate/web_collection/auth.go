package web_collection

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// generate random string only contains letter and number
const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandStr(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return randStr(n)
}

func QQNetCafeHeaderMake(body any) map[string]string {
	headers := make(map[string]string, 3)
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := GetRandStr(8)
	headers["Timestamp"] = ts
	headers["Nonce"] = nonce
	secret := Config.GetString("config.netcafe.secretcall")
	source := Config.GetString("config.netcafe.sourcename")
	bodyJson, err := json.Marshal(body)
	if err != nil {
		// error
		return nil
	}
	headers["Auth"] = HmacSha512(secret, source+"|"+ts+"|"+nonce+"|"+string(bodyJson))
	fmt.Println(headers)
	fmt.Printf("source:%s, timestamp:%s, nonce:%s, payload:%s", source, ts, nonce, string(bodyJson))
	fmt.Println("before hmac: ", source+"|"+ts+"|"+nonce+"|"+string(bodyJson))
	fmt.Println("after hamc:", HmacSha512(secret, source+"|"+ts+"|"+nonce+"|"+string(bodyJson)))
	fmt.Println("secret: ", HmacSha512("1", secret))
	return headers
}
