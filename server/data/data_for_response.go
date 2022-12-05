package data

import (
   	"log"
	"fmt"
   	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

//An instrument represents a combination of ExchangeCode & Exchange Pair Code
type Instrument struct {
	ExchangeCode string `json:"exchange_code"`
	ExchangePairCode string `json:"exchange_pair_code"`
}

//Instruments represents a list of instrument
type Instruments struct {
	Instruments []Instrument  `json:"data"`
}

//instruments is nested map 
//Key: ExchangeCode 
//Values: map of ExchangePairCode (in case of duplicated ExchangeCode-ExchangePairCode)
var instruments = map[string]map[string]string{}

/*
StoreData function's purpose is to store in the nested map "instruments" all unique combination
of ExchangeCode/ExchangePairCode 
*/
func StoreData(datas *Instruments) {
	nb_line := len(datas.Instruments)
	for i := 0; i < nb_line ; i++ {
		exchangeCode := datas.Instruments[i].ExchangeCode
		exchangePairCode := datas.Instruments[i].ExchangePairCode
		values, key_exist := instruments[exchangeCode]
		//exchange code already registered
		if key_exist {
			//check if exchange_pair_code already stored in the values
			_,value_exist := values[exchangePairCode]
			if !value_exist {
				values[exchangePairCode] = ""
			}
		} else {
			newPairCode := map[string]string{exchangePairCode: ""}
			instruments[exchangeCode] = newPairCode
		}
	}
}

/*
The function GetData() retrieve and parse data available from a fix endpoint
*/
func GetData() {
	var datas Instruments
	resp, err := http.Get("https://reference-data-api.kaiko.io/v1/instruments")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//read json
	json.Unmarshal(body, &datas)
	StoreData(&datas)
	fmt.Println("Data storage done")
}

/*
The function DataExist is the response rule - check if the combination ExchangeCode/ExchangePairCode exists
and return an ExistsCode accordingly.
*/
func DataExist(exchangeCode *string, exchangePairCode *string) int {
	values, key_exist := instruments[*exchangeCode]
	if !key_exist {
		//key doesn't exist, return 0
		return 0
	} else {
		//key exists, check for exchange_paire_code
		_,value_exist := values[*exchangePairCode]
		if value_exist {
			return 1
		} else {
			//try with upper case 
			_,value_exist := values[strings.ToUpper(*exchangePairCode)]
			if value_exist {
				return 1
			} else {
				return 2
			}
		}
	}
}