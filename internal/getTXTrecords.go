package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ark-go/beget/internal/iface"
)

func GetTXTrecords(setupDns *iface.SetupDns, fqdn string) ([]string, error) {
	type datTxt struct {
		Ttl     int    `json:"ttl,omitempty"`
		Txtdata string `json:"txtdata,omitempty"`
	}
	type Record struct {
		TXT []datTxt
	}
	type Records struct {
		Records Record
	}
	type out struct {
		Status     string `json:"status"`
		Error      string
		Error_code string
		Error_text string
		Answer     struct {
			Status string
			Result Records
			Errors []struct {
				Error_code string
				Error_text string
			}
		}
	}
	resp, err := http.Post("https://api.beget.com/api/dns/getData?input_format=json&output_format=json"+
		`&input_data={"fqdn":"`+fqdn+`"}`+
		`&login=`+setupDns.UserLoginDns+
		`&passwd=`+setupDns.UserPasswdDns,
		"application/json",
		nil)
	if err != nil {
		log.Fatal("Ошибка (txt): не получены данные DNS от API ", err.Error())
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка (txt) не прочитаны данные DNS от API ", err.Error())
	}
	//log.Println(">>", string(buf))
	res := &out{}
	err2 := json.Unmarshal(buf, res)
	if err2 != nil {
		log.Fatal("Ошибка (txt) не разобрать данные API")
	}
	if res.Status == "error" {
		log.Fatal("Ошибка (txt) не правильный запрос:", res.Error_text)
	}
	if res.Answer.Status == "error" {
		for i, v := range res.Answer.Errors {
			log.Println("Ошибка:", i, v.Error_text)
		}
		log.Fatal("Ошибка (txt) не правильные параметры запроса:")
	}
	if res.Answer.Status == "success" {
		var txt = []string{}
		//	for _, v := range res.Answer.Result {
		//log.Println("22>", res.Answer.Result.Records.TXT)
		for _, v := range res.Answer.Result.Records.TXT {
			if v.Txtdata != "" {
				txt = append(txt, v.Txtdata)
			}
		}
		//	log.Println(">>>>>>", txt, len(res.Answer.Result.Records.TXT), res.Answer.Result.Records.TXT)

		return txt, nil
	} else {
		log.Fatal("Ошибка (txt) ")
	}
	return nil, errors.New("Ошибка 756")
}
