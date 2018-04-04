package miner

import (
	"github.com/jinzhu/configor"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	"github.com/sonm-io/core/accounts"
	"github.com/sonm-io/core/insonmnia/benchmarks"
	"github.com/sonm-io/core/insonmnia/logging"
	"github.com/sonm-io/core/insonmnia/miner/plugin"
	"go.uber.org/zap/zapcore"
)

// HubConfig describes Hub configuration.
type HubConfig struct {
	EthAddr   string           `required:"true" yaml:"eth_addr"`
	Endpoints []string         `required:"false" yaml:"endpoints"`
	CGroups   *ResourcesConfig `required:"false" yaml:"resources"`
}

// FirewallConfig describes firewall detection settings.
type FirewallConfig struct {
	// STUN server endpoint (with port).
	Server string `yaml:"server"`
}

type SSHConfig struct {
	BindEndpoint   string `required:"true" yaml:"bind"`
	PrivateKeyPath string `required:"true" yaml:"private_key_path"`
}

type LoggingConfig struct {
	Level       string `required:"true" default:"debug"`
	parsedLevel zapcore.Level
}

type ResourcesConfig struct {
	Cgroup    string                `required:"true" yaml:"cgroup"`
	Resources *specs.LinuxResources `required:"false" yaml:"resources"`
}

type DevConfig struct {
	DevAddr  string `yaml:"listen"`
	Insecure bool   `yaml:"insecure"`
}

type storeConfig struct {
	Path   string `required:"true" yaml:"path" default:"/var/lib/sonm/worker.boltdb"`
	Bucket string `required:"true" yaml:"bucket" default:"sonm"`
}

type config struct {
	HubConfig               HubConfig           `required:"true" yaml:"hub"`
	Eth                     *accounts.EthConfig `yaml:"ethereum"`
	SSHConfig               *SSHConfig          `required:"false" yaml:"ssh"`
	LoggingConfig           LoggingConfig       `yaml:"logging"`
	PublicIPsConfig         []string            `required:"false" yaml:"public_ip_addrs"`
	MetricsListenAddrConfig string              `yaml:"metrics_listen_addr" default:"127.0.0.1:14001"`
	PluginsConfig           plugin.Config       `yaml:"plugins"`
	StoreConfig             storeConfig         `yaml:"store"`
	BenchConfig             benchmarks.Config   `yaml:"benchmarks"`
	DevConfig               *DevConfig          `yaml:"yes_i_want_to_use_dev-only_features"`
}

func (c *config) LogLevel() zapcore.Level {
	return c.LoggingConfig.parsedLevel
}

func (c *config) HubEthAddr() string {
	return c.HubConfig.EthAddr
}

func (c *config) HubEndpoints() (endpoints []string) {
	return c.HubConfig.Endpoints
}

func (c *config) HubResources() *ResourcesConfig {
	return c.HubConfig.CGroups
}

func (c *config) PublicIPs() []string {
	return c.PublicIPsConfig
}

func (c *config) SSH() *SSHConfig {
	return c.SSHConfig
}

func (c *config) ETH() *accounts.EthConfig {
	return c.Eth
}

func (c *config) MetricsListenAddr() string {
	return c.MetricsListenAddrConfig
}

func (c *config) Plugins() plugin.Config {
	return c.PluginsConfig
}

func (c *config) Dev() *DevConfig {
	return c.DevConfig
}

func (c *config) StorePath() string {
	return c.StoreConfig.Path
}

func (c *config) StoreBucket() string {
	return c.StoreConfig.Bucket
}

func (c *config) Benchmarks() benchmarks.Config {
	return c.BenchConfig
}

func (c *config) validate() error {
	if len(c.HubConfig.EthAddr) == 0 {
		return errors.New("hub's ethereum address should be specified")
	}

	return nil
}

// NewConfig creates a new Miner config from the specified YAML file.
func NewConfig(path string) (Config, error) {
	cfg := &config{}
	err := configor.Load(cfg, path)

	if err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	lvl, err := logging.ParseLogLevel(cfg.LoggingConfig.Level)
	if err != nil {
		return nil, err
	}
	cfg.LoggingConfig.parsedLevel = lvl

	return cfg, nil
}

// Config represents a Miner configuration interface.
type Config interface {
	logging.Leveler

	// HubEndpoints returns a string representation of a Hub endpoint to communicate with.
	HubEndpoints() []string
	// HubEthAddr returns hub's ethereum address.
	HubEthAddr() string
	// HubResources returns resources allocated for a Hub.
	HubResources() *ResourcesConfig
	// PublicIPs returns all IPs that can be used to communicate with the miner.
	PublicIPs() []string
	// SSH returns settings for built-in ssh server
	SSH() *SSHConfig
	// StorePath returns path to boltdb which keeps Worker's state.
	StorePath() string
	// StoreBucket returns boltdb data-bucket name.
	StoreBucket() string
	// ETH returns ethereum configuration
	ETH() *accounts.EthConfig
	// MetricsListenAddr returns the address that can be used by Prometheus to get
	// metrics.
	MetricsListenAddr() string
	// Plugins returns plugins settings.
	Plugins() plugin.Config
	// Benchmarks returns benchmarking settings.
	Benchmarks() benchmarks.Config
	// DevAddr to listen on. For dev purposes only!
	Dev() *DevConfig
}
