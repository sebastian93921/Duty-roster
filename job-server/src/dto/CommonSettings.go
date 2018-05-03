package dto

import (

)

var ServerProperties struct{
	ServerStartPort string		`json:"serverStartPort"`
	ServerStartUrl string 		`json:"serverStartUrl"`
	//HTTPS protocol
	SSLCertChain string		`json:"ssl.cert.fullchain"`
	SSLCertKey string		`json:"ssl.cert.key"`
	//Data Source
	CsvFile string		`json:"csv"`
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
