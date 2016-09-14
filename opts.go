package main

import "time"

var opts *Options

type (
	Options struct {
		ChannelSize        int           `help:"process internal data pipeline size" from:"env,flag"`
		PCount             int           `help:"calc process count" from:"env,flag"`
		ValidOffsetMinutes float64       `help:"gpstime offset time.Now valid minutes" from:"env,flag"`
		MinValidDuration   time.Duration `help:"gpstime offset time.Now min duration" from:"env,flag"`
		GpsList            string        `help:"gpsinfo list key name" from:"env,flag"`
		Distance           float64       `help:"distance between two points(max radius)" from:"env,flag"`
		LoadFence          bool          `help:"Whether the fence data is loaded at startup" from:"env,flag"`
		Debug              bool          `help:"logger level" from:"env,flag"`
		FromRedis          RedisOpts
		Tile38             Tile38Opts
		Db                 DbOpts
		Hook               HookOpts
	}

	RedisOpts struct {
		Addr      string        `help:"redis server address(addr:port)" from:"env,flag"`
		MaxActive int           `help:"redis connection pool max active count" from:"env,flag"`
		IdleTime  time.Duration `help:"redis connection pool keepalive duration" from:"env,flag"`
	}

	Tile38Opts struct {
		Addr      string        `help:"tile38 server address(addr:port)" from:"env,flag"`
		MaxActive int           `help:"tile38 connection pool max active count" from:"env,flag"`
		IdleTime  time.Duration `help:"tile38 connection pool keepalive duration" from:"env,flag"`
	}

	DbOpts struct {
		ConnStr string `help:"database connection string" from:"env,flag"`
		Sql     string `help:"query sql" from:"env,flag"`
	}

	HookOpts struct {
		Addr    string `help:"hook server address(addr:port)" from:"env,flag"`
		SetName string `help:"terminal coordinate SET name" from:"env,flag"`
	}
)

//缺省配置
var defaultOpts = &Options{
	ChannelSize:        150,
	PCount:             6,
	ValidOffsetMinutes: 30,
	GpsList:            "dpdb.queue.income.05",
	Distance:           50,
	LoadFence:          false,
	Debug:              false,
	FromRedis: RedisOpts{
		Addr:      "10.88.80.71:6379",
		MaxActive: 10,
		IdleTime:  2 * time.Second,
	},
	Tile38: Tile38Opts{
		Addr:      "127.0.0.1:9851",
		MaxActive: 10,
		IdleTime:  2 * time.Second,
	},
	Db: DbOpts{
		ConnStr: "apts/apts@127.0.0.1:1521/nbusdb",
		Sql:     "select route_code, service_number, bus_stop_code, latituded, longituded, route_id, routesub_id, station_id, order_number from vrm_routestop_gps",
	},
	Hook: HookOpts{
		Addr:    "127.0.0.1:9876",
		SetName: "BusXY",
	},
}
