package main

import (
	"github.com/beevik/etree"
	"strconv"
)

/*
可邑客流摄像机post xml 根节点为 EventNotificationAlert,
客流量节点为 ventNotificationAlert/peopleCounting,
然后是报文类型节点statisticalMethods，内容分两种：
一种是实时上报的realTime，节点为realTime，人数为当前触发事件时客流总和；
另一种是按时间段上报的timeRang，有startTime和endTime两个节点，人数为该区间增量。
需要的数据为进入人数enter和离开人数exit。 其它所需信息节点为ipAddress、macAddress、channelName。
*/

//Parse a XML document at the period of time
func period(xs string) *countInfo {
	count := countInfo{}
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xs); err != nil {
		panic(err)
	}

	inNm, _ := strconv.ParseInt(doc.FindElement("EventNotificationAlert/peopleCounting/enter").Text(), 10, 32)
	leNm, _ := strconv.ParseInt(doc.FindElement("EventNotificationAlert/peopleCounting/exit").Text(), 10, 32)

	count.ip = doc.FindElement("EventNotificationAlert/ipAddress").Text()
	count.mac = doc.FindElement("EventNotificationAlert/macAddress").Text()
	count.starTime = doc.FindElement("EventNotificationAlert/peopleCounting/TimeRange/startTime").Text()
	count.endTime = doc.FindElement("EventNotificationAlert/peopleCounting/TimeRange/endTime").Text()
	count.channelName = doc.FindElement("EventNotificationAlert/channelName").Text()
	count.enter = int32(inNm)
	count.leave = int32(leNm)

	//fmt.Printf("IP :%s  MAC :%s\n from %s to %s\n",count.ip,count.mac,count.starTime,count.endTime)
	//fmt.Printf("%s  in:%d   in:%d\n",count.channelName,count.enter,count.leave)

	return &count
}

/*func real(xs string) {
	var doc = etree.NewDocument()
	err := doc.ReadFromString(xs)
	if err != nil {
		return
	}
	//doc.WriteTo(os.Stdout)
	root := doc.SelectElement("EventNotificationAlert")
	if root == nil {
		fmt.Println("Get root fail")
		panic(err)
	}
	pc := root.SelectElement("peopleCounting")

	inNm, _ := strconv.ParseInt(pc.SelectElement("enter").Text(), 10, 32)
	leNm, _ := strconv.ParseInt(pc.SelectElement("exit").Text(), 10, 32)
	count := countInfo{
		ip:          pc.SelectElement("ipAddress").Text(),
		mac:         pc.SelectElement("macAddress").Text(),
		channelName: pc.SelectElement("channelName").Text(),
		starTime:    pc.SelectElement("starTime").Text(),
		endTime:     pc.SelectElement("endTime").Text(),
		enter:       int32(inNm),
		leave:       int32(leNm),
	}

	fmt.Println(pc.SelectElement("channelName").Text())
	fmt.Println("enter:", inNm, "  exit ", count.leave)
}*/
