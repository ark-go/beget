package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ark-go/beget/internal"
	"github.com/ark-go/beget/internal/iface"
	_ "github.com/ark-go/beget/internal/util" // читаем Env файл
)

func init() {
}

const (
	AUTH_ERROR        = "ошибка авторизации"
	INCORRECT_REQUEST = "ошибка, говорящая о некорректном запросе к API"
	NO_SUCH_METHOD    = "указанного метода не существует."

	INVALID_DATA  = "ошибка валидации переданных данных"
	LIMIT_ERROR   = "отказ в выполнении из-за достижения лимита"
	METHOD_FAILED = "внутренняя ошибка при выполнении метода."
)

func main() {
	setupDns := &iface.SetupDns{}
	setupDns.UserLoginDns = os.Getenv("UserLoginDns")
	setupDns.UserPasswdDns = os.Getenv("UserPasswdDns")
	setupDns.CertBotDomain = os.Getenv("CERTBOT_DOMAIN")         // Сertbot запросил домен или вставим свой
	setupDns.CertBotValidation = os.Getenv("CERTBOT_VALIDATION") // Код проверки
	flag.StringVar(&setupDns.NameSubDomain, "subdomain", iface.ACME_SUB_DOMEN, "только, если вы задаете собственный субдомен")
	//flag.StringVar(&setupDns.NameSubDomain, "subdomen", iface.ACME_SUB_DOMEN, "имя субдомена, для создания собственого субдомена")
	flag.StringVar(&setupDns.IPvalue, "ip", "", "запись A-type, можно указать IP адрес, пропишется при создании собственного субдомена (-subdomain)")
	flag.BoolVar(&setupDns.Save, "save", false, "Команда save, используется Certbot!")
	flag.BoolVar(&setupDns.Clear, "clear", false, "Команда clear, используется  Certbot!")
	flag.BoolVar(&setupDns.AddAcme, "addAcme", false, "Показать текущие домены, и добавить поддомен "+iface.ACME_SUB_DOMEN)
	flag.Usage = func() {
		fmt.Println("\nПредназначено для Certbot")
		flag.PrintDefaults()
	}
	flag.Parse()
	if setupDns.Clear && setupDns.Save {
		log.Fatal("Нельзя указать оба параметра -save и -clear")
	}
	log.Println("Привет:", setupDns.NameSubDomain, setupDns.UserLoginDns, setupDns.UserPasswdDns, setupDns.Save, setupDns.Clear)

	if setupDns.AddAcme {
		// log.Println("Текущие домены/поддомены:")
		listDom := internal.GetListDomain(setupDns)
		listDom.SelectDomain()
	}
	if setupDns.Save || setupDns.Clear {
		internal.ToogleCertBotCode(setupDns)
	}

}

// https://api.beget.com/api/domain/getList?login=userlogin&passwd=password&output_format=json
