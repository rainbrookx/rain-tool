package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

// C TODO 不要用全局的导出变量
var C Config

// InitConfig 加载配置
func InitConfig() {
	v := viper.New()

	// 设置配置文件路径
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("./configs")
	v.AddConfigPath(".")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		slog.Error(err.Error())
		panic("配置文件读取失败")
	}

	// 解析配置文件，反序列化到结构体
	if err := v.Unmarshal(&C); err != nil {
		slog.Error(err.Error())
		panic("配置解析文件失败")
	}

	// 校验配置文件
	if err := validate(&C); err != nil {
		slog.Error(err.Error())
		panic("配置文件校验失败")
	}
}

type Config struct {
	RainTool RainToolConfig `mapstructure:"rain_tool"`
}

type RainToolConfig struct {
	Addr string `mapstructure:"addr"`
}

func validate(c *Config) error {
	return nil
}
