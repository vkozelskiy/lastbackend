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

package storage

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/storage/storage"
)

type Util interface {
	Key(ctx context.Context, pattern ...string) string
}

type Storage interface {
	Cluster() storage.Cluster
	Deployment() storage.Deployment
	Namespace() storage.Namespace
	Node() storage.Node
	Ingress() storage.Ingress
	Pod() storage.Pod
	Route() storage.Route
	Secret() storage.Secret
	Service() storage.Service
	System() storage.System
	Endpoint() storage.Endpoint
	Trigger() storage.Trigger
	Volume() storage.Volume
	IPAM() storage.IPAM
}
