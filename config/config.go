package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/ardanlabs/conf/v3"
	"github.com/joho/godotenv"
)

// App config struct
type Config struct {
	Logger           Logger
	MavisRpc         string `json:"mavis_rpc" conf:"default:https://api.roninchain.com/rpc,env:MAVIS_RPC"`
	InfinityRpc      string `json:"infinity_rpc" conf:"env:INFINITY_RPC"`
	InfinityNvRpc    string `json:"infinity_nv_rpc" conf:"env:INFINITY_NV_RPC"`
	EternityRpc      string `json:"eternity_rpc" conf:"env:ETERNITY_RPC"`
	CatalystRpc      string `json:"catalyst_rpc" conf:"env:CATALYST_RPC"`
	InfinityGroupId  int    `json:"infinity_group_id" conf:"default:4282374336,env:INFINITY_GROUP_ID"`
	RoninNodeGroupId int    `json:"ronin_node_group_id" conf:"default:947505775,env:RONIN_NODE_GROUP_ID"`
	MaxBlockDelay    uint64 `json:"max_block_delay" conf:"default:5,env:MAX_BLOCK_DELAY"`
}

// Logger config
type Logger struct {
	Development       bool   `json:"log_development" conf:"default:false,env:LOG_DEV_MODE"`
	DisableCaller     bool   `json:"log_caller" conf:"default:false,env:LOG_CALLER"`
	DisableStacktrace bool   `json:"log_stacktrace" conf:"default:false,env:LOG_STACKTRACE"`
	Encoding          string `json:"log_encoding" conf:"default:json,env:LOG_ENCODING"`
	Level             string `json:"log_level" conf:"default:info,env:LOG_LEVEL"`
}

// Parse config file
func LoadConfig() (*Config, error) {
	godotenv.Load()
	cfg := &Config{}
	help, err := conf.Parse("", cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
		}

		log.Fatal(err)
	}

	return cfg, nil
}
