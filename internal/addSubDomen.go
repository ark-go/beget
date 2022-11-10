package internal

import (
	//	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ark-go/beget/internal/iface"
)

func AddSubDomen(setupDns *iface.SetupDns, infoDomain *ListDomainId) {
	// if setupDns.CertBotDomain == "" || setupDns.CertBotValidation == "" {
	// 	log.Println("Ошибка (save): Не указан домен или код для Certbot")
	// 	os.Exit(1)
	// }
	// ответ beget
	type outErr struct {
		Error_text string `json:"error_text"`
		Error_code string `json:"error_code"`
	}
	// ответ {"status":"success","answer":{"status":"success","result":10072829}}
	type out struct {
		Status     string `json:"status"`
		Error_text string `json:"error_text"`
		Error_code string `json:"error_code"`
		Answer     struct {
			Status     string   `json:"status"`
			Errors     []outErr `json:"errors"`
			Error_text string   `json:"error_text"`
			Error_code string   `json:"error_code"`
			Result     int64    `json:"result"`
		} `json:"answer"`
	}

	type subdomain = struct {
		Subdomain string `json:"subdomain"`
		Domain_id int64  `json:"domain_id"`
	}

	// Data := struct {
	// 	Subdomain subdomain `json:"subdomain"`
	// }{
	// 	Subdomain: subdomain{
	// 		Subdomain: setupDns.NameSubDomain, // + "." + infoDomain.Name,
	// 		Domain_id: infoDomain.Id,
	// 	},
	// }
	Data := subdomain{
		Subdomain: setupDns.NameSubDomain, // + "." + infoDomain.Name,
		Domain_id: infoDomain.Id,
	}

	bytesR, err := json.Marshal(Data)
	if err != nil {
		log.Fatalln("ошибока", err)
	}
	log.Println("подготовили:", string(bytesR))
	test5 := "input_data=" + string(bytesR)

	resp, err := http.Post("https://api.beget.com/api/domain/addSubdomainVirtual?input_format=json&login="+setupDns.UserLoginDns+"&passwd="+setupDns.UserPasswdDns+"&"+test5, "application/json", nil)
	if err != nil {
		log.Fatal("Ошибка (add): не получены данные от API ", err.Error())
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка (add) не прочитаны данные от API ", err.Error())
	}

	log.Println("ответ", string(buf))
	res := &out{}
	err2 := json.Unmarshal(buf, res)
	if err2 != nil {
		log.Fatal("Ошибка (add) не разобрать данные API", err2.Error())
	}
	if res.Status == "error" {
		log.Fatal("Ошибка (add) не правильный запрос:", res.Error_text)
	}
	if res.Answer.Status == "error" {
		for i, v := range res.Answer.Errors {
			log.Println(i+1, v.Error_text)
		}
		log.Fatal("Ошибка (add) не правильные параметры запроса:", res.Answer.Error_text)
	}
	if res.Answer.Status == "success" {

		log.Println("Готово (add) ", Data.Subdomain+"."+infoDomain.Name) //TODO сертбот ждет только пустой ответ 0 завершения программы
	} else {
		log.Fatal("Ошибка (add) ")
	}

	//id, ok := int64(buf)

}
