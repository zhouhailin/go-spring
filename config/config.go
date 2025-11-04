package config

import (
	"sync"

	"github.com/spf13/viper"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Env string

	Hertz      Hertz      `yaml:"hertz"`
	MySQL      MySQL      `yaml:"mysql"`
	Redis      Redis      `yaml:"redis"`
	Prometheus Prometheus `yaml:"prometheus"`
	Switch     Switch     `yaml:"switch"`
	Nlp        Nlp        `yaml:"nlp"`
	Speech     Speech     `yaml:"speech"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Switch struct {
	Address  string `yaml:"address"`
	Port     uint   `yaml:"port"`
	Password string `yaml:"password"`
	Timeout  uint   `yaml:"timeout"`
}

type Nlp struct {
	URL           string `yaml:"url" mapstructure:"url"`
	ErrorPlayFile string `yaml:"error" mapstructure:"error-play-file"`
	Hotwords      string `yaml:"hotwords" mapstructure:"hotwords"`
}

type Speech struct {
	Tts Tts `yaml:"tts"`
	Asr Asr `yaml:"asr"`
}

type Tts struct {
	Profile string `yaml:"profile" mapstructure:"profile"`
	Voice   string `yaml:"voice" mapstructure:"voice"`
	Volume  string `yaml:"volume" mapstructure:"volume"`
	Rate    string `yaml:"rate" mapstructure:"rate"`
}

type Asr struct {
	Profile                  string            `yaml:"profile" mapstructure:"profile"`
	GrammarUrl               string            `yaml:"grammar-url" mapstructure:"grammar-url"`
	StartInputTimers         string            `yaml:"start-input-timers" mapstructure:"start-input-timers"`
	NoInputTimeout           string            `yaml:"no-input-timeout" mapstructure:"no-input-timeout"`
	RecognitionTimeout       string            `yaml:"recognition-timeout" mapstructure:"recognition-timeout"`
	SpeechCompleteTimeout    string            `yaml:"speech-complete-timeout" mapstructure:"speech-complete-timeout"`
	SpeechIncompleteTimeout  string            `yaml:"speech-incomplete-timeout" mapstructure:"speech-incomplete-timeout"`
	SensitivityLevel         string            `yaml:"sensitivity-level" mapstructure:"sensitivity-level"`
	ConfidenceThreshold      string            `yaml:"confidence-threshold" mapstructure:"confidence-threshold"`
	DefineGrammar            bool              `yaml:"define-grammar" mapstructure:"define-grammar"`
	VendorSpecificParameters map[string]string `yaml:"vendor-specific-parameters" mapstructure:"vendor-specific-parameters"`
}

type Prometheus struct {
	Address string `yaml:"address"`
	Path    string `yaml:"path"`
}

type Hertz struct {
	Address         string `yaml:"address"`
	EnablePprof     bool   `mapstructure:"enable_pprof"`
	EnableGzip      bool   `mapstructure:"enable_gzip"`
	EnableAccessLog bool   `mapstructure:"enable_access_log"`
	LogLevel        string `mapstructure:"log_level"`
	LogFileName     string `mapstructure:"log_file_name"`
	LogMaxSize      int    `mapstructure:"log_max_size"`
	LogMaxBackups   int    `mapstructure:"log_max_backups"`
	LogMaxAge       int    `mapstructure:"log_max_age"`
}

// Get gets configuration instance
func Get() *Config {
	once.Do(initConf)
	return config
}

func initConf() {
	viper.SetConfigName("iivrf") // name of config file (without extension)
	viper.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/iivrf/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.iivrf") // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.AddConfigPath("./etc")    // optionally look for config in the working directory
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			hlog.Error("config file not found")
		} else {
			// Config file was found but another error was produced
			hlog.Error("config file was found but another error was produced")
		}
	}
	hlog.Infof("Enable pprof: %v", viper.GetBool("enable_pprof"))

	// will be uppercased automatically
	viper.SetEnvPrefix("iivrf")
	err := viper.BindEnv("id")
	if err != nil {
		hlog.Error("bind env failed")
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		hlog.Error("unmarshal config failed")
		return
	}
}

func LogLevel() hlog.Level {
	level := Get().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
