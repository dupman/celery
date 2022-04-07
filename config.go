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

type Config struct {
	RedisURL string `mapstructure:"CELERY_REDIS_URL" default:"redis://"`
	Workers  int    `mapstructure:"CELERY_WORKERS" default:"5"`
}
