/*
Copyright 2023.

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
	"github.com/ellioben/operator-hook/utils"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	pluginv1beta1 "github.com/ellioben/operator-hook/api/v1beta1"
)

// ExtPowerReconciler reconciles a ExtPower object
type ExtPowerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=plugin.github.com.ellioben,resources=extpowers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=plugin.github.com.ellioben,resources=extpowers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=plugin.github.com.ellioben,resources=extpowers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ExtPower object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ExtPowerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// TODO(user): your logic here
	logger := log.FromContext(ctx)
	extPower := &pluginv1beta1.ExtPower{}
	//从缓存中获取app
	if err := r.Get(ctx, req.NamespacedName, extPower); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	//根据app的配置进行处理
	//1. Deployment的处理
	deployment := utils.NewDeployment(extPower)
	if err := controllerutil.SetControllerReference(extPower, deployment, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}
	//查找同名deployment
	d := &v1.Deployment{}
	if err := r.Get(ctx, req.NamespacedName, d); err != nil {
		if errors.IsNotFound(err) {
			if err := r.Create(ctx, deployment); err != nil {
				logger.Error(err, "create deploy failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if err := r.Update(ctx, deployment); err != nil {
			return ctrl.Result{}, err
		}
	}

	//2. Service的处理
	service := utils.NewService(extPower)
	if err := controllerutil.SetControllerReference(extPower, service, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}
	//查找指定service
	s := &corev1.Service{}
	if err := r.Get(ctx, types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, s); err != nil {
		if errors.IsNotFound(err) && extPower.Spec.EnableService {
			if err := r.Create(ctx, service); err != nil {
				logger.Error(err, "create service failed")
				return ctrl.Result{}, err
			}
		}
		//Fix: 这里还需要修复一下
	} else {
		if extPower.Spec.EnableService {
			//Fix: 当前情况下，不需要更新，结果始终都一样
			if err := r.Update(ctx, service); err != nil {
				return ctrl.Result{}, err
			}
		} else {
			if err := r.Delete(ctx, s); err != nil {
				return ctrl.Result{}, err
			}

		}
	}

	//3. Ingress的处理,ingress配置可能为空
	ingress := utils.NewIngress(extPower)
	if err := controllerutil.SetControllerReference(extPower, ingress, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}
	i := &netv1.Ingress{}
	if err := r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, i); err != nil {
		if errors.IsNotFound(err) && extPower.Spec.EnableIngress {
			if err := r.Create(ctx, ingress); err != nil {
				logger.Error(err, "create ingress failed")
				return ctrl.Result{}, err
			}
		}
		if !errors.IsNotFound(err) && extPower.Spec.EnableIngress {
			return ctrl.Result{}, err
		}
	} else {
		if extPower.Spec.EnableIngress {
			logger.Info("skip update")
		} else {
			if err := r.Delete(ctx, i); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExtPowerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	//owns表示这些内置组件被更新后，也会触发reconcile
	return ctrl.NewControllerManagedBy(mgr).
		For(&pluginv1beta1.ExtPower{}).
		Owns(&v1.Deployment{}).
		Owns(&netv1.Ingress{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
