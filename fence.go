package main

import (
	"encoding/json"

	"github.com/go-playground/log"
)

type FenceXY struct {
	RouteCode   string  `json:"RouteCode"`   //围栏ID,线路代码
	Service     string  `json:"Service"`     //围栏ID,服务
	BusstopCode string  `json:"BusstopCode"` //围栏ID,站点编码
	Lat         float64 `json:"Lat"`
	Lon         float64 `json:"Lon"`
	RouteId     int64   `json:"RouteId"`     //附加参数
	RouteSubId  int64   `json:"RouteSubId"`  //附加参数
	StationId   int64   `json:"StationId"`   //附加参数
	OrderNumber int64   `json:"OrderNumber"` //附加参数
}

func (r *FenceXY) Json() string {
	bs, err := json.Marshal(r)
	if err != nil {
		log.WithFields(log.F("func", "FenceXY.Json")).Warn(err)
		return ""
	}
	return string(bs)
}
