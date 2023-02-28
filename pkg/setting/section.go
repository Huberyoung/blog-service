package setting

import "time"

type ServerS struct {
	RunMode        string
	HttpPort       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type AppS struct {
	DefaultPageSize uint
	MaxPageSize     uint
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseS struct {
	DBType            string
	Username          string
	Password          string
	Host              string
	DBName            string
	TablePrefix       string
	Charset           string
	ParseTime         bool
	MaxIdleConnection int
	MaxOpenConnection int
}