package config

import (
	"fmt"
	"github.com/lightningnetwork/lnd/kvdb"
	"time"
)

const (
	defaultDbName = "vigilante.db"
)

type DBConfig struct {
	// DBPath is the directory path in which the database file should be
	// stored.
	DBPath string `long:"dbpath" description:"The directory path in which the database file should be stored."`

	// DBFileName is the name of the database file.
	DBFileName string `long:"dbfilename" description:"The name of the database file."`

	// NoFreelistSync, if true, prevents the database from syncing its
	// freelist to disk, resulting in improved performance at the expense of
	// increased startup time.
	NoFreelistSync bool `long:"nofreelistsync" description:"Prevents the database from syncing its freelist to disk, resulting in improved performance at the expense of increased startup time."`

	// AutoCompact specifies if a Bolt based database backend should be
	// automatically compacted on startup (if the minimum age of the
	// database file is reached). This will require additional disk space
	// for the compacted copy of the database but will result in an overall
	// lower database size after the compaction.
	AutoCompact bool `long:"autocompact" description:"Specifies if a Bolt based database backend should be automatically compacted on startup (if the minimum age of the database file is reached). This will require additional disk space for the compacted copy of the database but will result in an overall lower database size after the compaction."`

	// AutoCompactMinAge specifies the minimum time that must have passed
	// since a bolt database file was last compacted for the compaction to
	// be considered again.
	AutoCompactMinAge time.Duration `long:"autocompactminage" description:"Specifies the minimum time that must have passed since a bolt database file was last compacted for the compaction to be considered again."`

	// DBTimeout specifies the timeout value to use when opening the wallet
	// database.
	DBTimeout time.Duration `long:"dbtimeout" description:"Specifies the timeout value to use when opening the wallet database."`
}

func DefaultDBConfig() *DBConfig {
	return DefaultDBConfigWithHomePath(defaultAppDataDir)
}

func DefaultDBConfigWithHomePath(homePath string) *DBConfig {
	return &DBConfig{
		DBPath:            DataDir(homePath),
		DBFileName:        defaultDbName,
		NoFreelistSync:    true,
		AutoCompact:       false,
		AutoCompactMinAge: kvdb.DefaultBoltAutoCompactMinAge,
		DBTimeout:         kvdb.DefaultDBTimeout,
	}
}

func (cfg *DBConfig) ToBoltBackendConfig() *kvdb.BoltBackendConfig {
	return &kvdb.BoltBackendConfig{
		DBPath:            cfg.DBPath,
		DBFileName:        cfg.DBFileName,
		NoFreelistSync:    cfg.NoFreelistSync,
		AutoCompact:       cfg.AutoCompact,
		AutoCompactMinAge: cfg.AutoCompactMinAge,
		DBTimeout:         cfg.DBTimeout,
	}
}

func (cfg *DBConfig) Validate() error {
	if cfg.DBPath == "" {
		return fmt.Errorf("DB path cannot be empty")
	}

	if cfg.DBFileName == "" {
		return fmt.Errorf("DB file name cannot be empty")
	}
	return nil
}

func (cfg *DBConfig) GetDbBackend() (kvdb.Backend, error) {
	return kvdb.GetBoltBackend(cfg.ToBoltBackendConfig())
}
