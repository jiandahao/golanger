package cache

import (
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	crds "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"

	"gorm.io/gorm"
)

type (
	// ExecFn defines the exec method.
	ExecFn func() error

	// QueryFn defines the query method.
	QueryFn func(v interface{}) error

	// A CachedConn is a DB connection with cache capability.
	CachedConn interface {
		// Exec runs given exec on given keys,
		Exec(exec ExecFn, keys ...string) error
		// QueryRow unmarshals into v with given key and query func.
		QueryRow(v interface{}, query QueryFn, key string) error
		// SetCache sets v into cache with given key.
		SetCache(key string, v interface{}) error
		// DelCache deletes cache with keys.
		DelCache(keys ...string) error
		// GetCache unmarshals cache with given key into v.
		GetCache(key string, v interface{}) error
	}

	// Cache interface is used to define the cache implementation.
	Cache interface {
		Del(keys ...string) error
		Get(key string, v interface{}) error
		IsNotFound(err error) bool
		Set(key string, v interface{}) error
		SetWithExpire(key string, v interface{}, expire time.Duration) error
		Take(v interface{}, key string, query func(v interface{}) error) error
		TakeWithExpire(v interface{}, key string, query func(v interface{}, expire time.Duration) error) error
	}

	defaultCachedConn struct {
		dbConn *gorm.DB
		cache  Cache
	}
)

var (
	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	exclusiveCalls = syncx.NewSingleFlight()
	stats          = cache.NewStat("sqlc")
)

// Config cache configure.
type Config struct {
	Addr        string
	Pass        string
	Expiry      time.Duration
	ErrNotFound error
}

func newCacheOption(c Config) []cache.Option {
	var opts []cache.Option
	if c.Expiry > 0 {
		opts = append(opts, cache.WithExpiry(c.Expiry))
	}
	return opts
}

// NewDefaultCache creates a default cache which is provided by go-zero.
func NewDefaultCache(c Config) Cache {
	rds := crds.New(c.Addr, crds.WithPass(c.Pass))
	cc := cache.NewNode(rds, exclusiveCalls, stats, c.ErrNotFound, newCacheOption(c)...)
	return cc
}

// NewDefaultCachedConn creates a cached conn.
func NewDefaultCachedConn(cc Cache) CachedConn {
	return &defaultCachedConn{
		cache: cc,
	}
}

func (cc *defaultCachedConn) Exec(exec ExecFn, keys ...string) error {
	if err := exec(); err != nil {
		return err
	}

	if err := cc.DelCache(keys...); err != nil {
		return err
	}

	return nil
}

func (cc *defaultCachedConn) QueryRow(v interface{}, query QueryFn, key string) error {
	return cc.cache.Take(v, key, func(v interface{}) error {
		return query(v)
	})
}

func (cc *defaultCachedConn) DelCache(keys ...string) error {
	return cc.cache.Del(keys...)
}

func (cc *defaultCachedConn) SetCache(key string, v interface{}) error {
	return cc.cache.Set(key, v)
}

func (cc *defaultCachedConn) GetCache(key string, v interface{}) error {
	return cc.cache.Get(key, v)
}
