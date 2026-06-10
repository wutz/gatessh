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

type SSHRouteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SSHRouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var route sshv1alpha1.SSHRoute
	if err := r.Get(ctx, req.NamespacedName, &route); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("reconciling SSHRoute", "name", route.Name, "namespace", route.Namespace)

	for i, ref := range route.Spec.ParentRefs {
		ns := route.Namespace
		if ref.Namespace != "" {
			ns = ref.Namespace
		}

		var gw sshv1alpha1.SSHGateway
		if err := r.Get(ctx, client.ObjectKey{Name: ref.Name, Namespace: ns}, &gw); err != nil {
			if i < len(route.Status.Parents) {
				meta.SetStatusCondition(&route.Status.Parents[i].Conditions, metav1.Condition{
					Type:               "Accepted",
					Status:             metav1.ConditionFalse,
					ObservedGeneration: route.Generation,
					Reason:             "GatewayNotFound",
					Message:            "Referenced gateway not found",
				})
			}
			continue
		}

		if len(route.Status.Parents) <= i {
			route.Status.Parents = append(route.Status.Parents, sshv1alpha1.RouteParentStatus{
				ParentRef: ref,
			})
		}

		meta.SetStatusCondition(&route.Status.Parents[i].Conditions, metav1.Condition{
			Type:               "Accepted",
			Status:             metav1.ConditionTrue,
			ObservedGeneration: route.Generation,
			Reason:             "Accepted",
			Message:            "Route accepted by gateway",
		})

		meta.SetStatusCondition(&route.Status.Parents[i].Conditions, metav1.Condition{
			Type:               "ResolvedRefs",
			Status:             metav1.ConditionTrue,
			ObservedGeneration: route.Generation,
			Reason:             "ResolvedRefs",
			Message:            "All references resolved",
		})
	}

	if err := r.Status().Update(ctx, &route); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SSHRouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sshv1alpha1.SSHRoute{}).
		Complete(r)
}
