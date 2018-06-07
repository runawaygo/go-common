package go_common

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"gopkg.in/yaml.v2"
)

var conMap = map[string]string{
	"k8s":  k8s,
	"dev":  dev,
	"prod": prod,
	"local": local,
	"ci": ci,
}

const (
	k8s    = "http://spring-config-server.ingress.98.cn"
	dev    = "http://spring-config-server.ingress.98.cn"
	local  = "http://spring-config-server.ingress.98.cn"
	ci     = "http://spring-config-server.ingress.98.cn"
	prod   = "http://conf.jinyi999.cn"
)

func GetAppConfig(appName string, v interface{}) {

	environment := os.Getenv("GO_ENV")
	if environment == "" {
		environment = "k8s"
	}

	resp, err := http.Get(fmt.Sprintf("%s/%s.yml", conMap[environment], appName))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalf("ReadAll data error : %v", err)
	}

	err = yaml.Unmarshal([]byte(body), v)
	if err != nil {
		log.Fatalf("Unmarshal data error : %v", err)
	}

	log.Printf("%v",v)

	resp, err = http.Get(fmt.Sprintf("%s/%s-%s.yml", conMap[environment], appName, environment))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	err = yaml.Unmarshal([]byte(body), v)

	if err != nil {
		log.Fatalf("Unmarshal data error : %v", err)
	}

	log.Printf("%v",v)


}
