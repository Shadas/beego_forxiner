package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego"
)

type ChinaWeatherXml struct {
	City []ChinaWeatherXmlCity `xml:"city"`
}

type ChinaWeatherXmlCity struct {
	QuName        string `xml:"quName,attr"`
	PyName        string `xml:"pyName,attr",attr`
	CityName      string `xml:"cityname,attr"`
	State1        string `xml:"state1,attr"`
	State2        string `xml:"state2,attr"`
	StateDetailed string `xml:"stateDetailed,attr"`
	Tem1          string `xml:"tem1,attr"`
	Tem2          string `xml:"tem2,attr"`
	WindState     string `xml:"windState,attr"`
}

func GetChinaWeatherByCity(content string) (string, error) {
	contentbyte := []byte(content)
	citybyte := contentbyte[0 : len(contentbyte)-6]
	cwx, err := GetXmlChinaWeather()
	if err != nil {
		return "", err
	}

	var wr *ChinaWeatherXmlCity

	for _, c := range cwx.City {
		if c.CityName == string(citybyte) {
			wr = &c
			break
		}
	}

	var retStr string

	if wr == nil {
		retStr = "未查找到城市"
	} else {
		retStr = "城市：“" + wr.CityName + "”，省份：“" + wr.QuName + "”，温度：“" + wr.Tem2 + "℃~" + wr.Tem1 + "℃”，天气：“" + wr.StateDetailed + "”，风向：“" + wr.WindState + "”"
	}
	return retStr, nil
}

func GetXmlChinaWeather() (*ChinaWeatherXml, error) {

	cwx := ChinaWeatherXml{}
	req := httplib.Get("http://flash.weather.com.cn/wmaps/xml/china.xml")
	byteBody, err := req.Bytes()
	if err != nil {
		// log.Println(err.Error())
		return nil, err
	}

	// log.Println(string(byteBody))
	// byteBody, err := DoRequest("GET", "http://flash.weather.com.cn/wmaps/xml/china.xml", []byte{})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return nil, err
	// }
	err = xml.Unmarshal(byteBody, &cwx)
	if err != nil {
		// log.Println(err.Error())
		return nil, err
	}
	return &cwx, nil
}
