package conf

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server   ServerCfg        `yaml:"server"`
	Mysql    MysqlCfg          `yaml:"mysql"`
	Logger   LoggerConfig      `yaml:"logger"`
	Stat     StatConfig        `yaml:"stat"`
}

type ServerCfg struct {
	Port     string    `yaml:"port"`
	Mode     string    `yaml:"mode"`
	Version  string    `yaml:"version"`
}

type MysqlCfg struct {
	Db  struct {
		SpecialData struct {
			Name    string    `yaml:"name"`
			Host    string    `yaml:"host"`
			Port    int       `yaml:"port"`
			User    string    `yaml:"user"`
			Passwd  string    `yaml:"passwd"`
			Tables   struct {
				BlackUrlList        string     `yaml:"black_url_list"`
				BlackInfoHashList   string     `yaml:"black_infohash_list"`
				BlackEmulehashList  string     `yaml:"black_emulehash_list"`
			}    `yaml:"tables"`
		}            `yaml:"special_data"`
    }    `yaml:"db"`
}


type LoggerConfig struct {
	Level          zap.AtomicLevel      `yaml:"level"`
	OutputPaths     []string             `yaml:"output_paths"`
}

type StatConfig struct {
	PrometheusHttpAddr    string  `yaml:"prometheus_http_addr"`
}


func (cfg *Config)Load(filename string) ( error) {
	data, err  := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return err
	}
	return nil
}

func  NewConfig() *Config {
	return &Config{}
}