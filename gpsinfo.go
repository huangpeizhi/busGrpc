package main

import (
	"encoding/json"
	"time"

	"github.com/go-playground/log"
	geo "github.com/kellydunn/golang-geo"
)

//GPS消息
type GpsInfo struct {
	ObuId     string    `json:"ObuId"`
	Lat       float64   `json:"Lat"`       //纬度
	Lon       float64   `json:"Lon"`       //经度
	Speed     float64   `json:"Speed"`     //速度
	Direction float64   `json:"Direction"` //方向
	GpsTime   time.Time `json:"GpsTime"`   //时间
	RouteCode string    `json:"RouteCode"` //线路，前缀匹配符之一
	Service   string    `json:"Service"`   //服务，前缀匹配符之一
}

func (r *GpsInfo) Distance(i *GpsInfo) float64 {
	p1 := geo.NewPoint(r.Lat, r.Lon)
	p2 := geo.NewPoint(i.Lat, i.Lon)

	return p1.GreatCircleDistance(p2) * 1000
}

func (r *GpsInfo) Json() string {
	bs, err := json.Marshal(r)
	if err != nil {
		log.WithFields(log.F("func", "GpsInfo.Json")).Warn(err)
		return ""
	}
	return string(bs)
}
