//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package envs

import (
	"github.com/lastbackend/lastbackend/pkg/discovery/cache"
	"github.com/lastbackend/lastbackend/pkg/storage"
)

var _env Env

type Env struct {
	storage storage.Storage
	cache   *cache.Cache
}

func Get() *Env {
	return &_env
}

func (c *Env) SetStorage(storage storage.Storage) {
	c.storage = storage
}

func (c *Env) GetStorage() storage.Storage {
	return c.storage
}

func (c *Env) SetCache(cache *cache.Cache) {
	c.cache = cache
}

func (c *Env) GetCache() *cache.Cache {
	return c.cache
}
