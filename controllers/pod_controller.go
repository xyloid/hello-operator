/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// PodReconciler reconciles a Pod object
type PodReconciler struct {
	client.Client
	Scheme  *runtime.Scheme
	podInfo map[string]PodInfo
}

type PodInfo struct {
	podname     string
	nodename    string
	cpu_limit   int
	cpu_request int
}

//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	clog := log.FromContext(ctx)

	myLog := clog.WithValues("pod", req.NamespacedName)
	// your logic here
	fmt.Printf("\n\nReconcile function is called: %s\n", req.NamespacedName)

	var pod corev1.Pod
	if err := r.Get(ctx, req.NamespacedName, &pod); err != nil {
		myLog.Error(err, "unable to fetch Pod")
		return ctrl.Result{}, err
	}

	if pod.Spec.NodeName != "" {
		fmt.Printf("Pod node name:%s.\n", pod.Spec.NodeName)
	}
	fmt.Printf("Pod node name nominated:%s.\n", pod.Status.NominatedNodeName)

	cpu_limit, ok := pod.Spec.Containers[0].Resources.Limits["cpu"]

	if ok {
		fmt.Printf("CPU limit %d \n", cpu_limit.Value())
	}

	cpu_request, ok := pod.Spec.Containers[0].Resources.Limits["cpu"]

	if ok {
		fmt.Printf("CPU request %d \n", cpu_request.Value())
	}

	fmt.Printf("Pod scheduler: %s \n", pod.Spec.SchedulerName)
	fmt.Printf("Pod scheduler: %s \n", &pod.Status)
	fmt.Printf("Pod limits:%v\n", pod.Spec.Containers[0].Resources.Limits)
	fmt.Printf("Pod limits:%v\n", pod.Spec.Containers[0].Resources.Limits["cpu"])
	fmt.Printf("Pod limits:%v\n", pod.Spec.Containers[0].Resources.Limits["cpu"].Format)
	fmt.Printf("Pod requests:%v\n", pod.Spec.Containers[0].Resources.Requests)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.podInfo = make(map[string]PodInfo)
	return ctrl.NewControllerManagedBy(mgr).
		// Uncomment the following line adding a pointer to an instance of the controlled resource as an argument
		For(&corev1.Pod{}).
		Complete(r)
}
