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

type SSHGatewayReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SSHGatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var gw sshv1alpha1.SSHGateway
	if err := r.Get(ctx, req.NamespacedName, &gw); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var gc sshv1alpha1.SSHGatewayClass
	if err := r.Get(ctx, client.ObjectKey{Name: gw.Spec.GatewayClassName}, &gc); err != nil {
		return ctrl.Result{}, err
	}

	if gc.Spec.ControllerName != controllerName {
		return ctrl.Result{}, nil
	}

	log.Info("reconciling SSHGateway", "name", gw.Name, "namespace", gw.Namespace)

	meta.SetStatusCondition(&gw.Status.Conditions, metav1.Condition{
		Type:               "Accepted",
		Status:             metav1.ConditionTrue,
		ObservedGeneration: gw.Generation,
		Reason:             "Accepted",
		Message:            "Gateway accepted",
	})

	meta.SetStatusCondition(&gw.Status.Conditions, metav1.Condition{
		Type:               "Programmed",
		Status:             metav1.ConditionTrue,
		ObservedGeneration: gw.Generation,
		Reason:             "Programmed",
		Message:            "Gateway programmed",
	})

	gw.Status.Listeners = make([]sshv1alpha1.ListenerStatus, len(gw.Spec.Listeners))
	for i, l := range gw.Spec.Listeners {
		gw.Status.Listeners[i] = sshv1alpha1.ListenerStatus{
			Name: l.Name,
			Conditions: []metav1.Condition{
				{
					Type:               "Ready",
					Status:             metav1.ConditionTrue,
					ObservedGeneration: gw.Generation,
					Reason:             "Ready",
					Message:            "Listener ready",
					LastTransitionTime: metav1.Now(),
				},
			},
		}
	}

	if err := r.Status().Update(ctx, &gw); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SSHGatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sshv1alpha1.SSHGateway{}).
		Complete(r)
}
