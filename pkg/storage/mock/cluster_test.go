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

package mock

import (
	"context"
	"reflect"
	"testing"

	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/storage/storage"
)

// Test cluster storage set status method
func TestClusterStorage_SetStatus(t *testing.T) {

	var (
		stg = newClusterStorage()
		ctx = context.Background()
		c   = types.Cluster{}
	)

	type fields struct {
		stg storage.Cluster
	}

	type args struct {
		ctx     context.Context
		cluster *types.Cluster
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Cluster
		wantErr bool
		err     string
	}{
		{
			"test successful insert",
			fields{stg},
			args{ctx, &c},
			&c,
			false,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := stg.Clear(ctx); err != nil {
				t.Errorf("ClusterStorage.Get() storage setup error = %v", err)
				return
			}

			if err := tt.fields.stg.SetStatus(tt.args.ctx, &tt.args.cluster.Status); (err != nil) != tt.wantErr || (tt.wantErr && (err.Error() != tt.err)) {
				t.Errorf("ClusterStorage.Insert() error = %v, want errorr %v", err, tt.err)
			}
		})
	}
}

// Test cluster storage get return method
func TestClusterStorage_Get(t *testing.T) {

	var (
		stg = newClusterStorage()
		ctx = context.Background()
		c   = getClusterAsset()
	)

	type fields struct {
		stg storage.Cluster
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Cluster
		wantErr bool
	}{
		{
			"get cluster info successful",
			fields{stg},
			args{ctx},
			&c,
			false,
		},
	}

	for _, tt := range tests {

		if err := stg.Clear(ctx); err != nil {
			t.Errorf("ClusterStorage.Get() storage setup error = %v", err)
			return
		}

		if err := stg.SetStatus(ctx, &c.Status); err != nil {
			t.Errorf("ClusterStorage.Get() storage setup error = %v", err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.stg.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterStorage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test storage initialization
func Test_newClusterStorage(t *testing.T) {

	tests := []struct {
		name string
		want storage.Cluster
	}{
		{"initialize storage",
			newClusterStorage(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newClusterStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClusterStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getClusterAsset() types.Cluster {
	var c = types.Cluster{
		Status: types.ClusterStatus{
			Nodes: types.ClusterStatusNodes{
				Total:   2,
				Online:  1,
				Offline: 1,
			},
			Capacity: types.ClusterResources{
				Containers: 1,
				Pods:       1,
				Memory:     1024,
				Cpu:        1,
				Storage:    512,
			},
			Allocated: types.ClusterResources{
				Containers: 1,
				Pods:       1,
				Memory:     1024,
				Cpu:        1,
				Storage:    512,
			},
			Deleted: false,
		},
	}
	return c
}
