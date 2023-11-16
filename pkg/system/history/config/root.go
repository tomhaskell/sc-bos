package config

import (
	"github.com/vanti-dev/sc-bos/internal/util/pgxutil"
	"github.com/vanti-dev/sc-bos/pkg/system"
	"github.com/vanti-dev/sc-bos/pkg/util/jsontypes"
)

type Root struct {
	system.Config
	Storage *Storage `json:"storage,omitempty"`
}

type StorageType string

const (
	StorageTypePostgres StorageType = "postgres"
)

type Storage struct {
	Type StorageType `json:"type,omitempty"`
	pgxutil.ConnectConfig
	TTL *TTL `json:"ttl,omitempty"`
}

type TTL struct {
	MaxAge   jsontypes.Duration `json:"maxAge,omitempty"`
	MaxCount int64              `json:"maxCount,omitempty"`
}
