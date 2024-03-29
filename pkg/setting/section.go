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
	DefaultPageSize         int
	MaxPageSize             int
	LogSavePath             string
	LogFileName             string
	LogFileExt              string
	LogMaxSize              int
	LogMaxAge               int
	LogMaxBackups           int
	LogUseLocalTime         bool
	LogCompress             bool
	UploadSavePath          string
	UploadServerUrl         string
	UploadImageMaxSize      int
	DefaultContextTimeout   time.Duration
	UploadImageAllowExtList []string
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

type JwtS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}
