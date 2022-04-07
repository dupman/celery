/*
 * This file is part of the dupman/celery project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package celery

import (
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

const (
	MaxIdle     = 3
	MaxActive   = 3
	IdleTimeout = 240 * time.Second
)

func NewClient(conf *Config) (*gocelery.CeleryClient, error) {
	redisPool := &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(conf.RedisURL)
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
	}

	return gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		gocelery.NewRedisBackend(redisPool),
		conf.Workers,
	)
}
