package main

import "time"

type human []struct {
	DateTime         time.Time `json:"dateTime"`
	IPAddress        string    `json:"ipAddress"`
	Protocol         string    `json:"protocol"`
	MacAddress       string    `json:"macAddress"`
	ChannelID        int       `json:"channelID"`
	DeviceID         string    `json:"deviceID"`
	ActivePostCount  int       `json:"activePostCount"`
	EventType        string    `json:"eventType"`
	EventState       string    `json:"eventState"`
	EventDescription string    `json:"eventDescription"`
	ChannelName      string    `json:"channelName"`
	CaptureResult    []struct {
		TargetID int `json:"targetID"`
		Human    struct {
			Rect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"Rect"`
			Property []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"Property"`
			ContentID1 string `json:"contentID1"`
			PID1       string `json:"pId1"`
		} `json:"Human"`
		NonMotor struct {
			Rect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"Rect"`
			Property []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"Property"`
			ContentID1 string `json:"contentID1"`
			PID1       string `json:"pId1"`
		} `json:"NonMotor"`
		Face struct {
			Rect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"Rect"`
			Property []struct {
				Description string `json:"description"`
				Value       int    `json:"value"`
			} `json:"Property"`
			ContentID1      string `json:"contentID1"`
			PID1            string `json:"pId1"`
			FacePictureRect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"FacePictureRect"`
		} `json:"Face"`
		UID string `json:"uid"`
	} `json:"CaptureResult"`
}

type vehic []struct {
	DateTime         time.Time `json:"dateTime"`
	IPAddress        string    `json:"ipAddress"`
	Protocol         string    `json:"protocol"`
	MacAddress       string    `json:"macAddress"`
	ChannelID        int       `json:"channelID"`
	DeviceID         string    `json:"deviceID"`
	ActivePostCount  int       `json:"activePostCount"`
	EventType        string    `json:"eventType"`
	EventState       string    `json:"eventState"`
	EventDescription string    `json:"eventDescription"`
	ChannelName      string    `json:"channelName"`
	CaptureResult    []struct {
		TargetID int `json:"targetID"`
		Vehicle  struct {
			VehicleRect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"VehicleRect"`
			PlateRect struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"PlateRect"`
			Property []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"Property"`
			ContentID1 string `json:"contentID1"`
			ContentID2 string `json:"contentID2"`
			PID1       string `json:"pId1"`
			PID2       string `json:"pId2"`
		} `json:"Vehicle"`
		UID string `json:"uid"`
	} `json:"CaptureResult"`
}