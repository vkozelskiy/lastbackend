//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
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
	"fmt"
	"github.com/lastbackend/lastbackend/pkg/apis/types"
	"github.com/lastbackend/lastbackend/pkg/storage/store"
	"github.com/lastbackend/lastbackend/pkg/util/generator"
	"golang.org/x/net/context"
	"time"
)

const ProjectTable string = "projects"

// Project Service type for interface in interfaces folder
type ProjectStorage struct {
	IProject
	Client func() (store.IStore, store.DestroyFunc, error)
}

// Get project by name for user
func (s *ProjectStorage) GetByID(username, id string) (*types.Project, error) {
	var (
		project = new(types.Project)
		key     = fmt.Sprintf("%s/%s/%s/info", ProjectTable, username, id)
	)

	client, destroy, err := s.Client()
	if err != nil {
		return nil, err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Get(ctx, key, &project.Meta); err != nil {
		if err.Error() == store.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}

	project.User = username

	return project, nil
}

// Get project by name for user
func (s *ProjectStorage) GetByName(username, name string) (*types.Project, error) {

	var (
		id string
		// Key example: /helper/projects/<username>/<name>
		key = fmt.Sprintf("/helper/%s/%s/%s", ProjectTable, username, name)
	)

	client, destroy, err := s.Client()
	if err != nil {
		return nil, err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Get(ctx, key, &id); err != nil {
		if err.Error() == store.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}

	return s.GetByID(username, id)
}

// List project by username
func (s *ProjectStorage) ListByUser(username string) (*types.ProjectList, error) {
	var (
		key    = fmt.Sprintf("%s/%s", ProjectTable, username)
		filter = `\b(.+)\/info\b`
	)

	client, destroy, err := s.Client()
	if err != nil {
		return nil, err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	metaList := []types.Meta{}

	if err := client.List(ctx, key, filter, &metaList); err != nil {
		if err.Error() == store.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}

	if metaList == nil {
		return nil, nil
	}

	projectList := new(types.ProjectList)
	for _, meta := range metaList {
		*projectList = append(*projectList, types.Project{Meta: meta, User: username})
	}

	return projectList, nil
}

// Insert new project into storage  for user
func (s *ProjectStorage) Insert(username, name, description string) (*types.Project, error) {
	var (
		id        = generator.GetUUIDV4()
		project   = new(types.Project)
		keyHelper = fmt.Sprintf("/helper/%s/%s/%s", ProjectTable, username, name)
		keyInfo   = fmt.Sprintf("%s/%s/%s/info", ProjectTable, username, id)
	)

	project.ID = id
	project.Name = name
	project.User = username
	project.Description = description
	project.Labels = map[string]string{"tier": "ptoject"}
	project.Updated = time.Now()
	project.Created = time.Now()

	client, destroy, err := s.Client()
	if err != nil {
		return nil, err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx := client.Begin(ctx)

	if err := tx.Create(keyHelper, &id, 0); err != nil {
		return nil, err
	}

	if err := tx.Create(keyInfo, &project.Meta, 0); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return project, nil
}

// Update project model
func (s *ProjectStorage) Update(username string, project *types.Project) (*types.Project, error) {
	var (
		keyInfo = fmt.Sprintf("%s/%s/%s/info", ProjectTable, username, project.ID)
	)

	project.Updated = time.Now()

	client, destroy, err := s.Client()
	if err != nil {
		return project, err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Update(ctx, keyInfo, &project.Meta, nil, 0); err != nil {
		return project, err
	}

	return project, nil
}

// Remove project model
func (s *ProjectStorage) Remove(username, id string) error {
	var (
		key = fmt.Sprintf("%s/%s/%s/info", ProjectTable, username, id)
	)

	client, destroy, err := s.Client()
	if err != nil {
		return err
	}
	defer destroy()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Delete(ctx, key, nil); err != nil {
		return err
	}

	return nil
}

func newProjectStorage(config store.Config) *ProjectStorage {
	s := new(ProjectStorage)
	s.Client = func() (store.IStore, store.DestroyFunc, error) {
		return New(config)
	}
	return s
}
