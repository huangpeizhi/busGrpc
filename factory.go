package main

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/go-playground/log"
	_ "github.com/mattn/go-oci8"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/tidwall/gjson"
	"github.com/tidwall/tile38/client"
)

const (
	rubyDatetimeLayout string = "20060102 150405" //公交使用的日期格式
	ksep                      = ":"
)

type Factory struct {
	fromPool     *pool.Pool
	tile38Pool   *client.Pool
	db           *sql.DB
	chGpsWaiting chan *GpsInfo
}

func (f *Factory) AnalysisLoop() {
	conn, err := f.tile38Pool.Get()
	if err != nil {
		log.WithFields(log.F("func", "Factory.AnalysisLoop.til38Pool.Get")).Fatal(err)
	}
	defer conn.Close()

	retMsg, err := conn.Do("OUTPUT json") //output json
	if err != nil {
		log.WithFields(log.F("func", "Factory.AnalysisLoop.Output")).Warn(err)
	}
	log.WithFields(log.F("func", "Factory.AnalysisLoop.Output"),
		log.F("status", "Done"),
		log.F("cmd", "OUTPUT json")).Debug(string(retMsg))

	for {
		for i := range f.chGpsWaiting {
			busXY := strings.Join([]string{i.RouteCode, i.Service}, ksep)
			cmd := fmt.Sprintf("SET %s %s POINT %f %f", busXY, i.ObuId, i.Lat, i.Lon)

			log.WithFields(log.F("func", "Factory.AnalysisLoop"),
				log.F("status", "Do"),
				log.F("cmd", cmd)).Debug(i.Json())

			retMsg, err := conn.Do(cmd)
			if err != nil {
				log.WithFields(log.F("func", "Factory.AnalysisLoop"),
					log.F("jstr", i.Json()),
					log.F("cmd", cmd)).Warn(err)
			}

			log.WithFields(log.F("func", "Factory.AnalysisLoop"),
				log.F("status", "Done"),
				log.F("cmd", cmd),
				log.F("jstr", i.Json())).Debug(string(retMsg))
		}
	}
}

//采集Gps信息
func (f *Factory) CollectGpsLoop() {
	for {
		jstr, err := f.Get(opts.GpsList)
		if err != nil {
			log.WithFields(log.F("func", "Factory.CollectGpsLoop")).Warn(err)
			continue
		}

		m := &GpsInfo{
			ObuId:     gjson.Get(jstr, "obuid").String(),
			Lat:       gjson.Get(jstr, "latitude").Float(),
			Lon:       gjson.Get(jstr, "longitude").Float(),
			Speed:     gjson.Get(jstr, "speed").Float(),
			Direction: gjson.Get(jstr, "direction").Float(),
			RouteCode: gjson.Get(jstr, "route_code").String(),
			Service:   gjson.Get(jstr, "service").String(),
		}

		if m.RouteCode == "" {
			log.WithFields(log.F("func", "Factory.CollectGpsLoop.RouteCode"), log.F("jstr", jstr)).Warn("routecode is empty")
			continue
		}

		if m.Service == "" {
			log.WithFields(log.F("func", "Factory.CollectGpsLoop.Service"), log.F("jstr", jstr)).Warn("service is empty")
			continue
		}

		t, err := time.ParseInLocation(rubyDatetimeLayout, gjson.Get(jstr, "gps_time").String(), time.Local)
		if err != nil {
			log.WithFields(log.F("func", "Factory.CollectGpsLoop.Parse.GpsTime"), log.F("jstr", jstr)).Warn(err.Error())
			continue
		}

		if math.Abs(time.Now().Sub(t).Minutes()) > opts.ValidOffsetMinutes {
			log.WithFields(log.F("func", "Factory.CollectGpsLoop.ValidOffsetMinutes"),
				log.F("jstr", jstr),
				log.F("GpsTime", t.String())).Info("offset exceeds the threshold")
			continue
		}

		m.GpsTime = t
		log.WithFields(log.F("func", "Factory.CollectGpsLoop")).Debug(m.Json())
		f.chGpsWaiting <- m
	}
}

func (f *Factory) Get(queueName string) (string, error) {
	resp := f.fromPool.Cmd("BRPOP", queueName, 0)
	if resp.Err != nil {
		log.WithFields(log.F("func", "Factory.Get.BRPOP"), log.F("queue", queueName)).Warn(resp.Err)
		return "", resp.Err
	}

	lst, err := resp.List()
	if err != nil || len(lst) != 2 {
		log.WithFields(log.F("func", "Factory.Get.LIST"), log.F("queue", queueName)).Warn(err)
		return "", err
	}

	return lst[1], nil
}

func NewFactory() (*Factory, error) {
	f, err := initRedisPool(opts.FromRedis.Addr, opts.FromRedis.MaxActive, opts.FromRedis.IdleTime)
	if err != nil {
		log.WithFields(log.F("func", "NewFactory.initRedisPool.FromRedis")).Warn(err)
		return nil, err
	}

	t38, err := initTile38Pool(opts.Tile38.Addr)
	if err != nil {
		log.WithFields(log.F("func", "NewFactory.initTile38Pool.Tile38")).Warn(err)
		return nil, err
	}

	db, err := sql.Open("oci8", opts.Db.ConnStr)
	if err != nil {
		return nil, err
	}

	factory := &Factory{
		fromPool:     f,
		tile38Pool:   t38,
		db:           db,
		chGpsWaiting: make(chan *GpsInfo, opts.ChannelSize),
	}

	return factory, nil
}

func (f *Factory) Destory() {
	if f.tile38Pool != nil {
		err := f.tile38Pool.Close()
		if err != nil {
			log.WithFields(log.F("func", "Factory.Destory")).Warn(err)
		}
	}
}

//建立tile38连接池
func initTile38Pool(addr string) (*client.Pool, error) {
	p, err := client.DialPool(addr)
	if err != nil {
		log.WithFields(log.F("func", "initTile38Pool")).Warn(err)
		return nil, err
	}

	return p, nil
}

//建立REDIS连接池
func initRedisPool(addr string, maxActive int, idleTime time.Duration) (*pool.Pool, error) {
	p, err := pool.New("tcp", addr, maxActive)
	if err != nil {
		log.WithFields(log.F("func", "initRedisPool")).Warn(err)
		return nil, err
	}

	go func() {
		for {
			p.Cmd("PING")
			time.Sleep(idleTime)
		}
	}()

	return p, nil
}

func (f *Factory) initFence() error {
	log.WithFields(log.F("func", "Factory.initFence")).Info("init fence start")

	rows, err := f.db.Query(opts.Db.Sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	defer f.db.Close()

	conn, err := f.tile38Pool.Get()
	if err != nil {
		log.WithFields(log.F("func", "Factory.initFence.Pool.Get")).Warn(err)
		return err
	}
	defer conn.Close()

	for rows.Next() {
		s := &FenceXY{}

		err := rows.Scan(
			&s.RouteCode, &s.Service, &s.BusstopCode,
			&s.Lat, &s.Lon,
			&s.RouteId, &s.RouteSubId, &s.StationId, &s.OrderNumber)
		if err != nil {
			log.WithFields(log.F("func", "Factory.initFence.Scan"), log.F("jstr", s.Json())).Warn(err)
			continue
		}

		id := strings.Join([]string{s.RouteCode, s.Service, s.BusstopCode}, ksep)
		busXY := strings.Join([]string{s.RouteCode, s.Service}, ksep)
		hookCmd := fmt.Sprintf("SETHOOK %s grpc://%s NEARBY %s FENCE DETECT enter,exit POINT %f %f %f",
			id,
			opts.Hook.Addr,
			busXY,
			s.Lat, s.Lon, opts.Distance)

		_, err = conn.Do(hookCmd)
		if err != nil {
			log.WithFields(log.F("func", "Factory.initFence.Pool.SETHOOK"),
				log.F("hookcmd", hookCmd),
				log.F("jstr", s.Json())).Warn(err)
		}
	}

	log.WithFields(log.F("func", "Factory.initFence")).Info("init fence end")
	return nil
}
