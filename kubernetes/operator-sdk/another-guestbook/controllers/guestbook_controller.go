/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "github.com/rh/for-sake-of-testing/kubernetes/operator-sdk/another-guestbook/api/v1"
)

// GuestBookReconciler reconciles a GuestBook object
type GuestBookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.mifomm.eu,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.mifomm.eu,resources=guestbooks/status,verbs=get;update;patch

// Reconcile defines ...
func (r *GuestBookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("guestbook", req.NamespacedName)

	log.Info("reconciling guestbook")

	var book webappv1.GuestBook
	if err := r.Get(ctx, req.NamespacedName, &book); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var redis webappv1.Redis
	redisName := client.ObjectKey{Name: book.Spec.RedisName, Namespace: req.Namespace}
	if err := r.Get(ctx, redisName, &redis); err != nil {
		log.Error(err, "didn't get redis")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("got redis", "redis", redis.Name)

	deployment, err := r.desiredDeployment(book, redis)
	if err != nil {
		return ctrl.Result{}, err
	}
	svc, err := r.desiredService(book)
	if err != nil {
		return ctrl.Result{}, err
	}

	applyOpts := []client.PatchOption{client.ForceOwnership, client.FieldOwner("guestbook-controller")}

	err = r.Patch(ctx, &deployment, client.Apply, applyOpts...)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.Patch(ctx, &svc, client.Apply, applyOpts...)
	if err != nil {
		return ctrl.Result{}, err
	}

	book.Status.URL = urlForService(svc, book.Spec.Frontend.ServingPort)
	log.Info("Status URL", "Status URL", book.Status.URL)

	err = r.Status().Update(ctx, &book)
	if err != nil {
		return ctrl.Result{}, err
	}

	log.Info("reconciled guestbook")
	return ctrl.Result{}, nil
}

// SetupWithManager defines ...
func (r *GuestBookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.GuestBook{}).
		Complete(r)
}
