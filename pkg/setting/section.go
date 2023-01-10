package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type AppSettingS struct {
	DefaultPageSize         int
	MaxPageSize             int
	LogSavePath             string
	LogFileName             string
	LogFileExt              string
	UploadSavePath          string
	UploadServerUrl         string
	UploadImageMaxSize      int
	UploadImageAllowExtList []string
}

type DatabaseSettingS struct {
	DBType             string
	Username           string
	Password           string
	Host               string
	DBName             string
	TablePrefix        string
	Charset            string
	ParseTime          bool
	MaxIdleConnections int
	MaxOpenConnections int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v any) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
