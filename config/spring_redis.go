package config

type Redis struct {
	Address  string        `json:"address" yaml:"address" xml:"address"`
	Password string        `json:"password" yaml:"password" xml:"password"`
	Database int           `json:"database" yaml:"database" xml:"database"`
	Cluster  ClusterRedis  `json:"cluster" yaml:"cluster" xml:"cluster" mapstructure:"cluster"`
	Sentinel SentinelRedis `json:"sentinel" yaml:"sentinel" xml:"sentinel" mapstructure:"sentinel"`
}

type ClusterRedis struct {
	Nodes        string `json:"nodes" yaml:"nodes" xml:"nodes"`
	MaxRedirects int    `json:"max-redirects" yaml:"max-redirects" xml:"max-redirects" mapstructure:"max-redirects"`
	Username     string `json:"username" yaml:"username" xml:"username"`
	Password     string `json:"password" yaml:"password" xml:"password"`
}

type SentinelRedis struct {
	Nodes    string `json:"nodes" yaml:"nodes" xml:"nodes"`
	Master   string `json:"master" yaml:"master" xml:"master"`
	Username string `json:"username" yaml:"username" xml:"username"`
	Password string `json:"password" yaml:"password" xml:"password"`
}
