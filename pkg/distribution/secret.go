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

package distribution

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/lastbackend/pkg/storage"
	"github.com/lastbackend/lastbackend/pkg/storage/store"
	"github.com/lastbackend/lastbackend/pkg/util/generator"
)

const (
	logSecretPrefix = "distribution:secret"
)

type ISecret interface {
	Get(namespace, name string) (*types.Secret, error)
	ListByNamespace(namespace string) (map[string]*types.Secret, error)
	Create(namespace *types.Namespace, opts *types.SecretCreateOptions) (*types.Secret, error)
	Update(secret *types.Secret, namespace *types.Namespace, opts *types.SecretUpdateOptions) (*types.Secret, error)
	Remove(secret *types.Secret) error
}

type Secret struct {
	context context.Context
	storage storage.Storage
}

func (n *Secret) Get(namespace, name string) (*types.Secret, error) {

	log.V(logLevel).Debugf("%s:get:> get secret by id %s/%s", logSecretPrefix, namespace, name)

	item, err := n.storage.Secret().Get(n.context, namespace, name)
	if err != nil {

		if err.Error() == store.ErrEntityNotFound {
			log.V(logLevel).Warnf("%s:get:> in namespace %s by name %s not found", logSecretPrefix, namespace, name)
			return nil, nil
		}

		log.V(logLevel).Errorf("%s:get:> in namespace %s by name %s error: %s", logSecretPrefix, namespace, name, err)
		return nil, err
	}

	return item, nil
}

func (n *Secret) ListByNamespace(namespace string) (map[string]*types.Secret, error) {

	log.V(logLevel).Debugf("%s:listbynamespace:> get secrets list by namespace", logSecretPrefix)

	items, err := n.storage.Secret().ListByNamespace(n.context, namespace)
	if err != nil {
		log.V(logLevel).Error("%s:listbynamespace:> get secrets list by namespace err: %s", logSecretPrefix, err)
		return items, err
	}

	log.V(logLevel).Debugf("%s:listbynamespace:> get secrets list by namespace result: %d", logSecretPrefix, len(items))

	return items, nil
}

func (n *Secret) Create(namespace *types.Namespace, opts *types.SecretCreateOptions) (*types.Secret, error) {

	log.V(logLevel).Debugf("%s:crete:> create secret %#v", logSecretPrefix, opts)

	secret := new(types.Secret)
	secret.Meta.SetDefault()
	secret.Meta.Name = generator.GenerateRandomString(10)
	secret.Meta.Namespace = namespace.Meta.Name
	if opts.Data != nil {
		secret.Data = *opts.Data
	}

	if err := n.storage.Secret().Insert(n.context, secret); err != nil {
		log.V(logLevel).Errorf("%s:crete:> insert secret err: %s", logSecretPrefix, err)
		return nil, err
	}

	return secret, nil
}

func (n *Secret) Update(secret *types.Secret, namespace *types.Namespace, opts *types.SecretUpdateOptions) (*types.Secret, error) {

	log.V(logLevel).Debugf("%s:update:> update secret %s", logSecretPrefix, secret.Meta.Name)

	if opts.Data != nil {
		secret.Data = *opts.Data
	}

	if err := n.storage.Secret().Update(n.context, secret); err != nil {
		log.V(logLevel).Errorf("%s:update:> update secret err: %s", logSecretPrefix, err)
		return nil, err
	}

	return secret, nil
}

func (n *Secret) Remove(secret *types.Secret) error {

	log.V(logLevel).Debugf("%s:remove:> remove secret %#v", logSecretPrefix, secret)

	if err := n.storage.Secret().Remove(n.context, secret); err != nil {
		log.V(logLevel).Errorf("%s:remove:> remove secret  err: %s", logSecretPrefix, err)
		return err
	}

	return nil
}

func NewSecretModel(ctx context.Context, stg storage.Storage) ISecret {
	return &Secret{ctx, stg}
}
