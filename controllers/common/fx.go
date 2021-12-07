// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package common

import (
	"context"
	"reflect"

	"github.com/chaos-mesh/chaos-mesh/apis/core/v1alpha1"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/chaos-mesh/chaos-mesh/controllers/config"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/builder"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/controller"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector"
)

type ChaosImplPair struct {
	Name   string
	Object v1alpha1.InnerObjectWithSelector
	Impl   ChaosImpl

	ObjectList v1alpha1.GenericChaosList
	Controlls  []client.Object
}

type Params struct {
	fx.In

	Mgr             ctrl.Manager
	Client          client.Client
	Logger          logr.Logger
	Selector        *selector.Selector
	RecorderBuilder *recorder.RecorderBuilder
	Impls           []*ChaosImplPair `group:"impl"`
	Reader          client.Reader    `name:"no-cache"`
}

func Bootstrap(params Params) error {
	logger := params.Logger
	pairs := params.Impls
	mgr := params.Mgr
	kubeclient := params.Client
	reader := params.Reader
	selector := params.Selector
	recorderBuilder := params.RecorderBuilder

	setupLog := logger.WithName("setup-common")
	for _, pair := range pairs {
		name := pair.Name + "-records"
		if !config.ShouldSpawnController(name) {
			return nil
		}

		setupLog.Info("setting up controller", "resource-name", pair.Name)

		builder := builder.Default(mgr).
			For(pair.Object).
			Named(name)

		// Add owning resources
		if len(pair.Controlls) > 0 {
			pair := pair
			for _, obj := range pair.Controlls {
				builder.Watches(&source.Kind{
					Type: obj,
				},
					handler.EnqueueRequestsFromMapFunc(func(obj client.Object) []reconcile.Request {
						reqs := []reconcile.Request{}
						objName := k8sTypes.NamespacedName{
							Namespace: obj.GetNamespace(),
							Name:      obj.GetName(),
						}

						list := pair.ObjectList.DeepCopyList()
						err := kubeclient.List(context.TODO(), list)
						if err != nil {
							setupLog.Error(err, "fail to list object")
						}

						items := reflect.ValueOf(list).Elem().FieldByName("Items")
						for i := 0; i < items.Len(); i++ {
							item := items.Index(i).Addr().Interface().(v1alpha1.InnerObjectWithSelector)
							for _, record := range item.GetStatus().Experiment.Records {
								namespacedName, err := controller.ParseNamespacedName(record.Id)
								if err != nil {
									setupLog.Error(err, "failed to parse record", "record", record.Id)
									continue
								}
								if namespacedName == objName {
									id := k8sTypes.NamespacedName{
										Namespace: item.GetNamespace(),
										Name:      item.GetName(),
									}
									setupLog.Info("mapping requests", "source", objName, "target", id)
									reqs = append(reqs, reconcile.Request{
										NamespacedName: id,
									})
								}
							}
						}
						return reqs
					}),
				)
			}
		}

		err := builder.Complete(&Reconciler{
			Impl:     pair.Impl,
			Object:   pair.Object,
			Client:   kubeclient,
			Reader:   reader,
			Recorder: recorderBuilder.Build("records"),
			Selector: selector,
			Log:      logger.WithName("records"),
		})
		if err != nil {
			return err
		}

	}

	return nil
}
