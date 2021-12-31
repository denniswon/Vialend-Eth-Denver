package include

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

const fn = "C:\\Users\\xdotk\\torukmakto\\vialend\\contracts\\vaults\\v2.0\\scripts\\file\\ini"

const cfg_file = "C:\\Users\\xdotk\\torukmakto\\vialend\\contracts\\vaults\\v2.0\\scripts\\file\\config.json"

type Configuration struct {
	STRAT_DEPLOYER string
	VAULT_DEPLOYER string
	VAULT_FACTORY  string
	VAULT_STRATEGY string
	VAULT          string
}

var Cfg Configuration

func ConfigWrite() {

	var Cfg2 Configuration

	Load(cfg_file, &Cfg2)

	if !reflect.DeepEqual(Cfg, Cfg2) {
		Save(cfg_file, &Cfg)
	}

}

func ConfigParser() {
	Load(cfg_file, &Cfg)
}

// Load gets your config from the json file,
// and fills your struct with the option
func Load(filename string, o interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(b, &o)
	}
	return err
}

// Save will save your struct to the given filename,
// this is a good way to create a json template
func Save(filename string, o interface{}) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(filename, j, 0660)
	}
	return err
}

func ReadINI() []string {

	settings := ReadFile(fn)

	if len(settings) < 3 {
		log.Fatal("check ini file. invalid settings")
	}

	return (settings)
}

func WriteINI(msgs []string) {

	WriteFile(fn, msgs, false)
}

//fn "../../file/tick"
func ReadFile(fi string) []string {

	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("ReadFile:Open: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	sln, e := Readln(r)

	fileInArray := []string{}

	for e == nil {

		s := strings.TrimSpace(sln)

		if len(s) > 0 {
			fileInArray = append(fileInArray, s) // a == [3 4]
		}
		sln, e = Readln(r)
	}

	return fileInArray

}

// fn "../file/tick"
func WriteFile(fn string, msgs []string, isAppend bool) {

	var buf []byte

	for i := 0; i < len(msgs); i++ {
		ss := []byte(msgs[i] + "\n")
		buf = append(buf, ss...)
	}

	if isAppend {
		d0, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal("WriteFileAppend:Readfile err:", err)
		}

		s := strings.TrimSpace(string(d0))

		buf = append(buf, s...)
	}
	err := os.WriteFile(fn, buf, 0644)
	if err != nil {
		log.Fatal("writefile err:", err)
	}

}
