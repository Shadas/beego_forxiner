package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
	"log"
)

type AutoReturnController struct {
	beego.Controller
}

type TextXml struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
}

func (c *AutoReturnController) AutoReturnHandler() {

	byteBody := c.Ctx.Input.RequestBody
	getxml := TextXml{}
	err := xml.Unmarshal(byteBody, &getxml)
	if err != nil {
		log.Println(err.Error())
	}
	//获取初次内容
	getcontent := getxml.Content

	//传入文字，分析
	retstr, err := AnalTextContent(getcontent)
	if err != nil {
		//处理错误
	}

	retbyte, _ := ResponseText(getxml.FromUserName, getxml.ToUserName, getxml.MsgType, retstr)
	log.Println(retbyte)
	c.Data["xml"] = &retbyte
	c.ServeXML()
}
