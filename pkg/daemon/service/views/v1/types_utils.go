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

package v1

import (
	"encoding/json"
	"fmt"
	"github.com/lastbackend/lastbackend/pkg/apis/types"
	"github.com/lastbackend/lastbackend/pkg/daemon/pod/views/v1"
	"github.com/lastbackend/lastbackend/pkg/util/table"
	"strings"
)

func New(obj *types.Service) *Service {
	s := new(Service)

	s.Meta.Name = obj.Meta.Name
	s.Meta.Description = obj.Meta.Description
	s.Meta.Region = obj.Meta.Region
	s.Meta.Updated = obj.Meta.Updated
	s.Meta.Created = obj.Meta.Created
	s.Meta.Replicas = obj.Meta.Replicas

	s.Config.Memory = obj.Config.Memory
	s.Config.Command = strings.Join(obj.Config.Command, " ")
	s.Config.Image = obj.Config.Image

	if len(obj.Pods) == 0 {
		s.Pods = make([]v1.PodInfo, 0)
		return s
	}

	for _, pod := range obj.Pods {
		s.Pods = append(s.Pods, v1.ToPodInfo(pod))
	}

	return s
}

func (obj *Service) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func NewList(obj *types.ServiceList) *ServiceList {
	s := new(ServiceList)
	if obj == nil {
		return nil
	}
	for _, v := range *obj {
		*s = append(*s, *New(&v))
	}
	return s
}

func (s *Service) DrawTable(namespaceName string) {
	table.PrintHorizontal(map[string]interface{}{
		"NAME":        s.Meta.Name,
		"DESCRIPTION": s.Meta.Description,
		"NAMESPACE":   namespaceName,
		"REPLICAS":    s.Meta.Replicas,
		"MEMORY":      s.Config.Memory,
		"IMAGE":       s.Config.Image,
		"CREATED":     s.Meta.Created,
		"UPDATED":     s.Meta.Updated,
	})

	fmt.Println("\n\tPODS")
	for _, pod := range s.Pods {
		table.PrintHorizontal(map[string]interface{}{
			"\tID":                 pod.Meta.ID,
			"\tSTATE":              pod.Meta.State.State,
			"\tSTATUS":             pod.Meta.State.Status,
			"\tTOTAL CONTAINERS":   pod.Meta.State.Containers.Total,
			"\tRUNNING CONTAINERS": pod.Meta.State.Containers.Running,
			"\tCREATED CONTAINERS": pod.Meta.State.Containers.Created,
			"\tSTOPPED CONTAINERS": pod.Meta.State.Containers.Stopped,
			"\tERRORED CONTAINERS": pod.Meta.State.Containers.Errored,
		})

		fmt.Println("\n\t\tCONTAINERS")
		for _, container := range pod.Containers {
			table.PrintHorizontal(map[string]interface{}{
				"\t\tID":      container.ID,
				"\t\tIMAGE":   container.Image,
				"\t\tSTATE":   container.State,
				"\t\tSTATUS":  container.Status,
				"\t\tPORTS":   container.Ports,
				"\t\tCREATED": container.Created,
				"\t\tSTARTED": container.Started,
			})
		}
	}
}

func (obj *ServiceList) ToJson() ([]byte, error) {
	if obj == nil || len(*obj) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(obj)
}

func (sl *ServiceList) DrawTable(namespaceName string) {
	t := table.New([]string{"NAME", "DESCRIPTION", "REPLICAS", "CREATED", "UPDATED"})
	t.VisibleHeader = true

	fmt.Println("NAMESPACE: ", namespaceName)
	for _, s := range *sl {
		t.AddRow(map[string]interface{}{
			"NAME":        s.Meta.Name,
			"DESCRIPTION": s.Meta.Description,
			"REPLICAS":    s.Meta.Replicas,
			"CREATED":     s.Meta.Created.String()[:10],
			"UPDATED":     s.Meta.Updated.String()[:10],
		})
	}

	t.AddRow(map[string]interface{}{})

	t.Print()
}
