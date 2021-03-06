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

package local

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/node/runtime/cpi"
)

type Proxy struct {
	cpi.CPI
}

func (p *Proxy) Info(ctx context.Context) (map[string]*types.EndpointStatus, error) {
	es := make(map[string]*types.EndpointStatus)
	return es, nil
}

func (p *Proxy) Create(ctx context.Context, endpoint *types.EndpointSpec) (*types.EndpointStatus, error) {
	return new(types.EndpointStatus), nil
}

func (p *Proxy) Destroy(ctx context.Context, endpoint *types.EndpointStatus) error {
	return nil
}

func (p *Proxy) Update(ctx context.Context, endpoint *types.EndpointStatus, spec *types.EndpointSpec) (*types.EndpointStatus, error) {
	return new(types.EndpointStatus), nil
}

func New() (*Proxy, error) {
	return &Proxy{}, nil
}
