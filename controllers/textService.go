package controllers

import (
	// "fmt"
	// "github.com/astaxie/beego/logs"
	"log"
	"strings"
)

func AnalTextContent(content string) (string, error) {
	var retstr string
	content = strings.TrimSpace(content)
	log.Println("收到了" + content)
	order, err := getOrderWithAnalText(content)
	if err != nil {
		//报错
	}

	switch order {
	case "功能":
		retstr = "hi，现在可以查询天气啦~比如“天津天气”，赶紧试试吧~"
	case "我爱你":
		retstr = "亲爱的，我也爱你~"
	case "我想你":
		retstr = "亲爱的，我也想你~"
	case "天气":
		retstr, err = GetChinaWeatherByCity(content)
		if err != nil {
			return "", err
		}
	case "笑话":
		retstr, err = DealGetJoke(content)
		if err != nil {
			return "", err
		}
	default:
		retstr = "亲爱的，我看到你对我说\" " + content + " \"~ 不来说点甜蜜的话嘛？回复“功能”可获取功能列表哦~"
	}
	return retstr, nil
}

func getOrderWithAnalText(content string) (string, error) {
	var retorder string

	if strings.Contains(content, "我爱你") {
		retorder = "我爱你"
		return retorder, nil
	}
	if strings.Contains(content, "我想你") {
		retorder = "我想你"
		return retorder, nil
	}
	if strings.TrimSpace(content) == "功能" {
		retorder = "功能"
		return retorder, nil
	}
	if (len(content) >= 6) && content[len(content)-6:len(content)] == "天气" {
		retorder = "天气"
		return retorder, nil
	}
	if (len(content) >= 6) && content[0:6] == "笑话" {
		retorder = "笑话"
		return retorder, nil
	}

	return "", nil
}
