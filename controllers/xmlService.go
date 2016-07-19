package controllers

import (
	"encoding/xml"
	"time"
)

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type ResponseTextXml struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATAText     `xml:"ToUserName"`
	FromUserName CDATAText     `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      CDATAText     `xml:"MsgType"`
	Content      CDATAText     `xml:"Content"`
}

func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func ResponseText(to, from, msgtype, content string) (ResponseTextXml, error) {
	v := ResponseTextXml{}
	v.ToUserName = value2CDATA(to)
	v.FromUserName = value2CDATA(from)
	v.CreateTime = time.Duration(time.Now().Unix())
	v.MsgType = value2CDATA(msgtype)
	v.Content = value2CDATA(content)
	return v, nil
}
