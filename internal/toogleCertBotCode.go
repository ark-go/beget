package internal

import (
	//	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ark-go/beget/internal/iface"
)

type datTxt struct {
	Priority int    `json:"priority"`
	Value    string `json:"value"`
}
type dataDns struct {
	Fqdn    string `json:"fqdn"`
	Records struct {
		TXT []datTxt //`json:"TXT"`
	} `json:"records"`
}

// func LoadDnsRecords(setupDns *iface.SetupDns) {
// 	type records map[string]any
// 	type out struct {
// 		Status string
// 		Error  string
// 		Answer struct {
// 			Status string
// 			Result string
// 			//	Records records
// 		}
// 	}
// 	resp, err := http.Post("https://api.beget.com/api/dns/getData?input_format=json&output_format=json"+
// 		`&input_data={"fqdn":"_acme-challenge.h-i-t.store"}`+
// 		`&login=`+setupDns.UserLoginDns+
// 		`&passwd=`+setupDns.UserPasswdDns,
// 		"application/json",
// 		nil)
// 	if err != nil {
// 		log.Fatal("Ошибка (save): не получены данные DNS от API ", err.Error())
// 	}
// 	buf, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal("Ошибка (save) не прочитаны данные DNS от API ", err.Error())
// 	}

//		log.Println("DNS:", string(buf))
//	}
func ToogleCertBotCode(setupDns *iface.SetupDns) {
	if setupDns.CertBotDomain == "" || setupDns.CertBotValidation == "" {
		log.Println("Ошибка (save): Не указан домен или код для Certbot")
		os.Exit(1)
	}
	// ответ beget
	type out struct {
		Status     string `json:"status"`
		Error_text string `json:"error_text"`
		Error_code string `json:"error_code"`
		Answer     struct {
			Status     string `json:"status"`
			Error_text string `json:"error_text"`
			Error_code string `json:"error_code"`
			Result     bool   `json:"result"`
		} `json:"answer"`
	}
	Data := &dataDns{
		Fqdn: setupDns.NameSubDomain + "." + setupDns.CertBotDomain,
		Records: struct {
			TXT []datTxt
		}{
			TXT: []datTxt{
				// {
				// 	Priority: 10,
				// 	Value:    "eetest",
				// },
			},
		},
	}
	if setupDns.Save {
		Data.Records.TXT = append(Data.Records.TXT, datTxt{10, setupDns.CertBotValidation}) // 1345
	}
	bytesR, err := json.Marshal(Data)
	if err != nil {
		log.Fatalln("ошибока", err)
	}
	log.Println("подготовили:", string(bytesR))
	test5 := "input_data=" + string(bytesR)
	// bytes.NewBuffer(bytesR)
	resp, err := http.Post("https://api.beget.com/api/dns/changeRecords?input_format=json&login="+setupDns.UserLoginDns+"&passwd="+setupDns.UserPasswdDns+"&"+test5, "application/json", nil)
	if err != nil {
		log.Fatal("Ошибка (save): не получены данные от API ", err.Error())
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка (save) не прочитаны данные от API ", err.Error())
	}

	//	log.Println("ответ", string(buf))
	res := &out{}
	err2 := json.Unmarshal(buf, res)
	if err2 != nil {
		log.Fatal("Ошибка (save) не разобрать данные API")
	}
	if res.Status == "error" {
		log.Fatal("Ошибка (save) не правильный запрос:", res.Error_text)
	}
	if res.Answer.Status == "error" {
		log.Fatal("Ошибка (save) не правильные параметры запроса:", res.Answer.Error_text)
	}
	if res.Answer.Status == "success" {
		log.Println("Готово (save) ")
	} else {
		log.Fatal("Ошибка (save) ")
	}

}
