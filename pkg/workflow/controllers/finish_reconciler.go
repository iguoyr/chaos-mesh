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

package controllers

import (
	"context"
	"fmt"
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type FinishReconciler struct {
	*ChildNodesFetcher
	kubeClient    client.Client
	eventRecorder recorder.ChaosRecorder
	logger        logr.Logger
}

func NewFinishReconciler(kubeClient client.Client, eventRecorder recorder.ChaosRecorder, logger logr.Logger) *FinishReconciler {
	return &FinishReconciler{
		ChildNodesFetcher: NewChildNodesFetcher(kubeClient, logger),
		kubeClient:        kubeClient,
		eventRecorder:     eventRecorder,
		logger:            logger}
}

func (it *FinishReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	node := v1alpha1.WorkflowNode{}
	err := it.kubeClient.Get(ctx, request.NamespacedName, &node)
	if err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// TODO 监听 failed 的 workflow node，传递给其子节点
	return reconcile.Result{}, nil
}

func (it *DeadlineReconciler) propagateToChildren(ctx context.Context, parent *v1alpha1.WorkflowNode) error {
	switch parent.Spec.Type {
	case v1alpha1.TypeSerial, v1alpha1.TypeParallel, v1alpha1.TypeTask:
		activeChildNodes, _, err := it.ChildNodesFetcher.fetchChildNodes(ctx, *parent)
		if err != nil {
			return err
		}
		for _, childNode := range activeChildNodes {
			childNode := childNode

			if WorkflowNodeFinished(childNode.Status) {
				it.logger.Info("child node already finished, skip for propagate deadline", "node", fmt.Sprintf("%s/%s", childNode.Namespace, childNode.Name))
				continue
			}

			err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				nodeNeedUpdate := v1alpha1.WorkflowNode{}
				err := it.kubeClient.Get(ctx, types.NamespacedName{
					Namespace: childNode.Namespace,
					Name:      childNode.Name,
				}, &nodeNeedUpdate)
				if err != nil {
					return err
				}
				if ConditionEqualsTo(nodeNeedUpdate.Status, v1alpha1.ConditionDeadlineExceed, corev1.ConditionTrue) {
					it.logger.Info("omit propagate deadline to children, child already in deadline exceed",
						"node", fmt.Sprintf("%s/%s", nodeNeedUpdate.Namespace, nodeNeedUpdate.Name),
						"parent node", fmt.Sprintf("%s/%s", parent.Namespace, parent.Name),
					)
					return nil
				}
				SetCondition(&nodeNeedUpdate.Status, v1alpha1.WorkflowNodeCondition{
					Type:   v1alpha1.ConditionDeadlineExceed,
					Status: corev1.ConditionTrue,
					Reason: v1alpha1.ParentNodeDeadlineExceed,
				})
				it.eventRecorder.Event(&nodeNeedUpdate, recorder.ParentNodeDeadlineExceed{ParentNodeName: parent.Name})
				return it.kubeClient.Status().Update(ctx, &nodeNeedUpdate)
			})
			if err != nil {
				return err
			}
			it.logger.Info("propagate deadline for child node",
				"child node", fmt.Sprintf("%s/%s", childNode.Namespace, childNode.Name),
				"parent node", fmt.Sprintf("%s/%s", parent.Namespace, parent.Name),
			)
		}
		return nil
	default:
		it.logger.V(4).Info("no need to propagate with this type of workflow node", "type", parent.Spec.Type)
		return nil
	}
}
