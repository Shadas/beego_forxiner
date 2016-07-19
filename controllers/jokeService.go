package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"strconv"
	"time"
)

type JokeResult struct {
	Code   int64  `json:"error_code"`
	Msg    string `json:"reason"`
	Result struct {
		Data []struct {
			Content string `json:"content"`
			// HashId string `json:"hashId"`

		} `json:"data"`
	} `json:"result"`
}

func DealGetJoke(content string) (string, error) {

	pagestr := content[6:len(content)]
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return "您的“笑话”后面没有跟数字，试试例如“笑话6”获取不同笑话~", nil
	}

	jr, err := GetJokeByInterface(int64(page)+1, 3)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	retstr := ""

	for i, data := range jr.Result.Data {
		retstr = retstr + fmt.Sprintf("%v", i+1) + ". " + data.Content + " \n"
	}

	return retstr, nil
}

func GetJokeByInterface(page, pagesize int64) (JokeResult, error) {
	jr := JokeResult{}
	//获得当前时间戳
	t := time.Now().AddDate(0, -1, 0).Unix()

	url := "http://japi.juhe.cn/joke/content/list.from?key=a242ed0e27ff19b27f94f12b10324f57&page=" + fmt.Sprintf("%v", page) + "&pagesize=" + fmt.Sprintf("%v", pagesize) + "&sort=asc&time=" + fmt.Sprintf("%v", t)

	req := httplib.Get(url)
	byteBody, err := req.Bytes()
	if err != nil {
		log.Println(err.Error())
		return jr, err
	}

	// byteBody, err := DoRequest("GET", url, []byte{})
	// if err != nil {
	// 	return jr, err
	// }

	if err := json.Unmarshal(byteBody, &jr); err != nil {
		return jr, errors.New("json result格式解析错误。")
	}

	return jr, nil

}
