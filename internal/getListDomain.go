package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ark-go/beget/internal/iface"
)

type commandApi struct {
	// Login         string `json:"login"`
	// Passwd        string `json:"passwd"`
	Input_format  string `json:"input_format"`
	Output_format string `json:"output_format"`
	// Input_data    userInfo `json:"input_data"`
}

func GetListDomain(setupDns *iface.SetupDns) *ListDomain {
	dom := getListDmn(setupDns, iface.Domain)
	subDom := getListDmn(setupDns, iface.SubDomain)
	listDomain := make(ListDomain)

	var num int64 = 1
	var numParent int64
	for _, v := range dom.Answer.Result {
		//TODO	w := slices.IndexFunc(subDom.Answer.Result, func(c iface.ResultData) bool { return c.Domain_id == v.Id })
		numParent = num
		listDomain[num] = &ListDomainId{Name: v.Fqdn, Id: v.Id, IsSub: false}
		num++
		for _, v2 := range subDom.Answer.Result { // перебор субдоменов
			if v2.Domain_id == v.Id { // в субдомене есть наш родитель

				listDomain[num] = &ListDomainId{Name: v2.Fqdn, Id: v2.Id, IdParent: v2.Domain_id, IsSub: true}
				if strings.HasPrefix(v2.Fqdn, setupDns.NameSubDomain) { // есть acme
					listDomain[num].Acme = true                            // помечаем суб-домен
					listDomain[num].AcmeCurr = true                        // это сам поддомен
					x := v2.Fqdn[len(setupDns.NameSubDomain):len(v2.Fqdn)] // выделим домен
					if listDomain[numParent].Name == x {                   // у родительского домена есть субдомен _acme
						listDomain[numParent].Acme = true // помечаем родительский домен
					}

					for _, vv := range listDomain {
						if vv.Name == x {
							vv.Acme = true
						}
					}
				}
				num++
			}
		}
	}
	listDomain.Print()
	return &listDomain
}

func getListDmn(s *iface.SetupDns, cmd iface.Command) *iface.ResultApi {
	query := &commandApi{
		Input_format:  "json",
		Output_format: "json",
	}
	bytesRepresentation, err := json.Marshal(query)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("%v", string(bytesRepresentation))
	//resp, err := http.Get("https://go.x.arkadii.ru/api/products")// getSubdomainList
	var resp *http.Response
	var err1 error
	if cmd == iface.Domain {
		resp, err1 = http.Post("https://api.beget.com/api/domain/getList?login=arkadi1i&passwd=ugYMEcL4", "application/json", bytes.NewBuffer(bytesRepresentation))
	}
	if cmd == iface.SubDomain {
		resp, err1 = http.Post("https://api.beget.com/api/domain/getSubdomainList?login=arkadi1i&passwd=ugYMEcL4", "application/json", bytes.NewBuffer(bytesRepresentation))
	}
	if err1 != nil {
		log.Fatalln(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := &iface.ResultApi{}

	err2 := json.Unmarshal(buf, &res)
	if err2 != nil {
		log.Println("ошибка Unmarshal")
	}
	if res.Status == "error" {
		log.Fatalln("Ошибка:", res.Error_text)
	}
	if res.Answer.Status == "error" {
		for i, v := range res.Answer.Errors {
			log.Println("Ошибка:", i, v.Error_text)
		}
		log.Fatalln("Выход с ошибкой")
	}
	// log.Printf("\n%+v", res.Answer.Result)
	// for i, v := range res.Answer.Result {
	// 	log.Println(i, v.Fqdn)
	// }

	return res
}
