package main

import (
	"fmt"
	"github.com/beevik/etree"
)

/*
可邑客流摄像机post xml 根节点为 EventNotificationAlert,客流量节点为 peopleCounting 。
然后是报文类型节点statisticalMethods，内容分两种：
一种是实时上报的statisticalMethods值为realTime，节点为realTime，人数为当前触发事件时客流总和；
另一种是按时间段上报的statisticalMethods值为timeRang，有startTime和endTime两个节点，人数为该区间增量。
需要的数据为进入人数enter和离开人数exit。 其它所需信息节点为ipAddress、macAddress、channelName。
*/

func Anaxml(xs string) {
	var doc = etree.NewDocument()
	if err := doc.ReadFromString(xs); err != nil {
		panic(err)
	}
	//doc.WriteTo(os.Stdout)
	root := doc.SelectElement("EventNotificationAlert")
	if root == nil {
		fmt.Println("root节点获取错误！")
		return
	}
	pc := root.SelectElement("peopleCounting")

	fmt.Println(pc.SelectElement("statisticalMethods").Text())
	fmt.Println(pc.SelectElement("enter").Text())
	fmt.Println(pc.SelectElement("exit").Text())
}
