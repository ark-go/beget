package iface

import ()

type Command int

const (
	Domain Command = iota + 1
	SubDomain
)

const ACME_SUB_DOMEN = "_acme-challenge" // точка мы используем для вырезания потом"

type SetupDns struct {
	NameSubDomain     string
	IPvalue           string
	Save              bool
	Clear             bool
	AddAcme           bool
	UserLoginDns      string
	UserPasswdDns     string
	CertBotDomain     string
	CertBotValidation string
	TimePropagation   int64
}

type ResultData struct {
	Id        int64 // запрашиваемый домен
	Fqdn      string
	Domain_id int64 // если есть то родительский домен,  при запросе поддоменов
	SubDomain []ResultData
}

type ResultApi struct {
	Status     string
	Error_text string
	Error_code string
	Answer     struct {
		Status string
		Result []ResultData
		Errors []struct {
			Error_code string
			Error_text string
		}
	}
}

// type ListDomainId struct {
// 	Id       int64
// 	Name     string
// 	IdParent int64
// 	IsSub    bool
// 	Acme     bool
// }
// type ListDomain map[int64]*ListDomainId

// func (ld ListDomain) Sort() (KeySorted []int64) {
// 	keys := make([]int64, 0, len(ld))
// 	for k := range ld {
// 		keys = append(keys, k)
// 	}
// 	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
// 	return keys
// }
// func (ld ListDomain) Print() {
// 	keys := ld.Sort()
// 	for _, k := range keys {
// 		log.Println("--", k, ld[k].Id, ld[k].IdParent, ld[k].Name)
// 	}
// }
// func (ld ListDomain) SelectDomain() (infoDomain *ListDomainId, err error) {
// 	if len(ld) == 0 {
// 		return nil, errors.New("список доменов пустой")
// 	}

// 	reader := bufio.NewReader(os.Stdin)      // reader, откуда будем читать - Stdin
// 	fmt.Printf("Введите номер строки:")      // выведем приглашение к вводу
// 	s, _ := reader.ReadString('\n')          // ждем ввод
// 	s = strings.TrimSpace(s)                 // обрезаем ентер
// 	intS, err := strconv.ParseInt(s, 10, 64) // переводим проверяем в число

// 	if err != nil {
// 		log.Println("Ошибка введите число:", err)
// 		return nil, err
// 	}
// 	d, ok := ld[intS]
// 	if ok {
// 		log.Println("нашли:", d)
// 		return d, nil
// 	} else {
// 		return nil, errors.New("не найден домен по этому номеру")
// 	}

// }
