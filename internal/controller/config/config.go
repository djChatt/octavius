package config

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetIntDefault(viper *viper.Viper, key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}

var once sync.Once
var config OctaviusConfig
var err error

type OctaviusConfig struct {
	viper                *viper.Viper
	LogLevel             string
	AppPort              string
	EtcdPort             string
	EtcdHost             string
	EtcdDialTimeout      time.Duration
	ExecutorPingDeadline time.Duration
	LogFilePath          string
	LogFileSize          int
}

func load() (OctaviusConfig, error) {
	fang := viper.New()

	fang.SetConfigType("json")
	fang.SetConfigName("controller_config")
	fang.AddConfigPath(".")

	//will be nil if file is read properly
	err := fang.ReadInConfig()
	if err != nil {
		return OctaviusConfig{}, err
	}
	octaviusConfig := OctaviusConfig{
		viper:                fang,
		LogLevel:             GetStringDefault(fang, "log_level", "info"),
		EtcdPort:             fang.GetString("etcd_port"),
		EtcdHost:             fang.GetString("etcd_host"),
		EtcdDialTimeout:      time.Duration(GetIntDefault(fang, "etcd_dial_timeout", 30)) * time.Second,
		AppPort:              fang.GetString("app_port"),
		ExecutorPingDeadline: time.Duration(GetIntDefault(fang, "executor_ping_deadline", 30)) * time.Second,
		LogFilePath:          GetStringDefault(fang, "log_file_path", "controller.log"),
		LogFileSize:          fang.GetInt("log_file_max_size_in_mb"),
	}
	return octaviusConfig, nil
}

type AtomBool struct{ flag int32 }

func (b *AtomBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), int32(i))
}

func (b *AtomBool) Get() bool {
	return atomic.LoadInt32(&(b.flag)) != 0
}

var reset = new(AtomBool)

func init() {
	reset.Set(false)
}

func Reset() {
	reset.Set(true)
}

func Loader() (OctaviusConfig, error) {
	once.Do(func() {
		config, err = load()
	})

	if reset.Get() {
		config, err = load()
		reset.Set(false)
	}
	if err != nil {
		return OctaviusConfig{}, err
	}
	return config, nil
}
