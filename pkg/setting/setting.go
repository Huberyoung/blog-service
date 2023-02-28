package setting

import "github.com/spf13/viper"

type Setting struct {
	*viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(sectionName string, v any) error {
	return s.UnmarshalKey(sectionName, v)
}
