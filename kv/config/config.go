package config

import (
	"fmt"
	"time"

	"github.com/pingcap-incubator/tinykv/log"
)

type Config struct {
	StoreAddr     string
	Raft          bool
	SchedulerAddr string
	LogLevel      string

	DBPath string // 用于存储数据的目录。应该存在并可写。
	//raft_base_tick_间隔是一个基本间隔（ms）。
	RaftBaseTickInterval     time.Duration
	RaftHeartbeatTicks       int
	RaftElectionTimeoutTicks int

	//到gc记录的间隔（ms）。
	RaftLogGCTickInterval time.Duration
	//当条目计数超过此值时，gc将被强制触发。
	RaftLogGcCountLimit uint64

	//检查区域是否需要分割的间隔（毫秒）。
	SplitRegionCheckTickInterval time.Duration
	//删除过时对等点之前的延迟时间
	SchedulerHeartbeatTickInterval      time.Duration
	SchedulerStoreHeartbeatTickInterval time.Duration

	//当区域[a，e]大小满足区域MaxSize时，它将被拆分为
	//几个区域[a，b]、[b，c]、[c，d]、[d，e]以及[a，b]的大小，
	//[b，c]，[c，d）将是区域分裂（可能稍大一点）。
	RegionMaxSize   uint64
	RegionSplitSize uint64
}

func (c *Config) Validate() error {
	if c.RaftHeartbeatTicks == 0 {
		return fmt.Errorf("heartbeat tick must greater than 0")
	}

	if c.RaftElectionTimeoutTicks != 10 {
		log.Warnf("Election timeout ticks needs to be same across all the cluster, " +
			"otherwise it may lead to inconsistency.")
	}

	if c.RaftElectionTimeoutTicks <= c.RaftHeartbeatTicks {
		return fmt.Errorf("election tick must be greater than heartbeat tick.")
	}

	return nil
}

const (
	KB uint64 = 1024
	MB uint64 = 1024 * 1024
)

func NewDefaultConfig() *Config {
	return &Config{
		SchedulerAddr:            "127.0.0.1:2379",
		StoreAddr:                "127.0.0.1:20160",
		LogLevel:                 "info",
		Raft:                     true,
		RaftBaseTickInterval:     1 * time.Second,
		RaftHeartbeatTicks:       2,
		RaftElectionTimeoutTicks: 10,
		RaftLogGCTickInterval:    10 * time.Second,
		//假设条目的平均大小为1k。
		RaftLogGcCountLimit:                 128000,
		SplitRegionCheckTickInterval:        10 * time.Second,
		SchedulerHeartbeatTickInterval:      100 * time.Millisecond,
		SchedulerStoreHeartbeatTickInterval: 10 * time.Second,
		RegionMaxSize:                       144 * MB,
		RegionSplitSize:                     96 * MB,
		DBPath:                              "/tmp/badger",
	}
}

func NewTestConfig() *Config {
	return &Config{
		LogLevel:                 "info",
		Raft:                     true,
		RaftBaseTickInterval:     50 * time.Millisecond,
		RaftHeartbeatTicks:       2,
		RaftElectionTimeoutTicks: 10,
		RaftLogGCTickInterval:    50 * time.Millisecond,
		// Assume the average size of entries is 1k.
		RaftLogGcCountLimit:                 128000,
		SplitRegionCheckTickInterval:        100 * time.Millisecond,
		SchedulerHeartbeatTickInterval:      100 * time.Millisecond,
		SchedulerStoreHeartbeatTickInterval: 500 * time.Millisecond,
		RegionMaxSize:                       144 * MB,
		RegionSplitSize:                     96 * MB,
		DBPath:                              "/tmp/badger",
	}
}
