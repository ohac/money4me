package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Secret struct {
	ServiceId string
	UserId    string
	Password  string
}

type Secrets struct {
	Secrets []Secret
}

func benefit401k(sec *Secret) (int64, error) {
	url := "https://www.benefit401k.com/"
	bow := surf.NewBrowser()
	err := bow.Open(url + "customer/")
	if err != nil {
		return 0, err
	}
	fm, _ := bow.Form("form#Form1")
	fm.Input("txtUserID", sec.UserId)
	fm.Input("txtPassword", sec.Password)
	if fm.Submit() != nil {
		return 0, err
	}
	bow.Open(url + "customer/RkDCMember/Home/JP_D_MemHome.aspx")
	a := ".BalanceAssets"
	var balance int64
	bow.Dom().Find(a).Each(func(_ int, s *goquery.Selection) {
		balancestr := s.Text()
		balancestr = strings.Replace(balancestr, ",", "", -1)
		balance32, _ := strconv.Atoi(balancestr)
		balance = int64(balance32)
	})
	return balance, nil
}

func main() {
	var secs Secrets
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &secs); err == nil {
		for _, sec := range secs.Secrets {
			switch sec.ServiceId {
			case "benefit401k":
				b, _ := benefit401k(&sec)
				fmt.Println(b) // TODO
			}
		}
	}
}
