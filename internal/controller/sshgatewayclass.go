package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	sshv1alpha1 "github.com/wutz/gatessh/api/v1alpha1"
)

const controllerName = "io.gatessh/controller"

type SSHGatewayClassReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SSHGatewayClassReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var gc sshv1alpha1.SSHGatewayClass
	if err := r.Get(ctx, req.NamespacedName, &gc); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if gc.Spec.ControllerName != controllerName {
		return ctrl.Result{}, nil
	}

	log.Info("reconciling SSHGatewayClass", "name", gc.Name)

	meta.SetStatusCondition(&gc.Status.Conditions, metav1.Condition{
		Type:               "Accepted",
		Status:             metav1.ConditionTrue,
		ObservedGeneration: gc.Generation,
		Reason:             "Accepted",
		Message:            "GatewayClass accepted by controller",
	})

	if err := r.Status().Update(ctx, &gc); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SSHGatewayClassReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sshv1alpha1.SSHGatewayClass{}).
		Complete(r)
}
