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

package state

import (
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
	"sync"
)

type VolumesState struct {
	lock    sync.RWMutex
	volumes map[string]types.VolumeSpec
}

func (s *VolumesState) GetVolumes() map[string]types.VolumeSpec {
	log.V(logLevel).Debug("Cache: VolumeCache: get pods")
	return s.volumes
}

func (s *VolumesState) SetVolumes(key string, volumes []*types.VolumeSpec) {
	log.V(logLevel).Debugf("Cache: VolumeCache: set volumes: %#v", volumes)
	for _, secret := range volumes {
		s.volumes[key] = *secret
	}
}

func (s *VolumesState) GetVolume(hash string) *types.VolumeSpec {
	log.V(logLevel).Debugf("Cache: VolumeCache: get secret: %s", hash)
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.volumes[hash]
	if !ok {
		return nil
	}
	return &v
}

func (s *VolumesState) AddVolume(key string, v *types.VolumeSpec) {
	log.V(logLevel).Debugf("Cache: VolumeCache: add volume: %#v", key)
	s.SetVolume(key, v)
}

func (s *VolumesState) SetVolume(key string, volume *types.VolumeSpec) {
	log.V(logLevel).Debugf("Cache: VolumeCache: set volume: %s", key)
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.volumes[key]; ok {
		delete(s.volumes, key)
	}

	s.volumes[key] = *volume
}

func (s *VolumesState) DelVolume(volume *types.Volume) {
	log.V(logLevel).Debugf("Cache: VolumeCache: del volume: %#v", volume)
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.volumes[volume.Meta.Name]; ok {
		delete(s.volumes, volume.Meta.Name)
	}
}
