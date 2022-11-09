package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type ListDomainId struct {
	Id       int64
	Name     string
	IdParent int64
	IsSub    bool
	Acme     bool
	AcmeCurr bool
}
type ListDomain map[int64]*ListDomainId

func (ld ListDomain) Sort() (KeySorted []int64) {
	keys := make([]int64, 0, len(ld))
	for k := range ld {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
func (ld ListDomain) Print() {
	log.Println("Текущие домены/поддомены:")
	keys := ld.Sort()
	for _, k := range keys {
		space := ""
		if ld[k].IsSub {
			space = "  "
		}
		if !ld[k].Acme {
			color.Set(color.FgGreen)
			fmt.Printf("%3d %s %s\n", k, space, ld[k].Name) //, ld[k].Id, ld[k].IdParent)
			color.Unset()
		} else {
			fmt.Printf("%3d %s %s\n", k, space, ld[k].Name) //, ld[k].Id, ld[k].IdParent)
		}
	}
}
func (ld ListDomain) SelectDomain() (infoDomain *ListDomainId, err error) {
	if len(ld) == 0 {
		return nil, errors.New("список доменов пустой")
	}

	reader := bufio.NewReader(os.Stdin) // reader, откуда будем читать - Stdin
	for {
		fmt.Println("Выберите домен для добавления поддомена _acme") // выведем приглашение к вводу
		fmt.Printf("Введите номер строки (выход Enter):")            // выведем приглашение к вводу
		s, _ := reader.ReadString('\n')                              // ждем ввод
		s = strings.TrimSpace(s)                                     // обрезаем ентер
		if len(s) == 0 {
			break
		}
		intS, err := strconv.ParseInt(s, 10, 64) // переводим проверяем в число

		if err != nil {
			log.Println("Ошибка введите число:", err)
			continue
			//return nil, err
		}
		d, ok := ld[intS]
		if ok {
			if d.Acme {
				if d.IsSub {
					if d.AcmeCurr {
						log.Println("Это сам поддомен acme:", d.Name)
					} else {
						log.Println("Поддодмен _acme_challenge, уже установлен для домена:", d.Name)
					}
				} else {
					log.Println("Поддодмен _acme_challenge, уже установлен для домена:", d.Name)
				}
				continue
				//return nil, errors.New("Уже есть поддодмен _acme_challenge")
			}
			log.Println("нашли:", d)
			return d, nil
		} else {
			log.Println("в списке нет такого номера", intS)
			continue
			//return nil, errors.New("не найден домен по этому номеру")
		}
	}
	return nil, errors.New("ошибка 243")
}
