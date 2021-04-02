package main

import (
	"errors"
	"fmt"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	Client *agollo.Client
)

type CountriesSt struct {
	Id        int64  `json:"id"`
	NameLocal string `json:"name_local"`
	NameZh    string `json:"name_zh"`
	NameEn    string `json:"name_en"`
	Code      string `json:"code"`
	Url       string `json:"url"`
}

func main() {
	//address := "http://apollo-service-apollo-configservice:18080"
	address := "http://127.0.0.1:18080"
	if err := detect(address); err != nil {
		log.Println(err)
	}
	c := &config.AppConfig{
		AppID:          "live-account",
		Cluster:        "dev",
		IP:             "http://127.0.0.1:18080",
		NamespaceName:  "application,HL.common,HL.i18n.id,HL.i18n.zh",
		IsBackupConfig: true,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		log.Println(err)
	}
	Client = client

	credentials := Client.GetConfig("HL.common").GetValue("countries")

	countriesResp := make([]CountriesSt, 0)
	//_ = json.Unmarshal([]byte(credentials), &countriesResp)
	_ = yaml.Unmarshal([]byte(credentials), &countriesResp)

	log.Println(countriesResp)
}

func detect(address string) (e error) {
	if address == "" {
		return errors.New("apollo svc address not set yet")
	}
	_, e = url.Parse(address)
	if e != nil {
		return
	}

	resp, e := http.Get(address)
	if e != nil {
		return
	}

	defer resp.Body.Close()

	content, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("apollo server resp code: %d", resp.StatusCode)
	}

	log.Printf("detect: apollo server resp body: %s", content)

	return nil
}
