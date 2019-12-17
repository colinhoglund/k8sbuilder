package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	appsv1.Deployment
}

func NewDeployment(name string) *Deployment {
	d := &Deployment{}
	d.TypeMeta = getExtensionsV1beta1TypeMeta("Deployment")
	d.ObjectMeta.Name = name
	return d
}

func (d *Deployment) WithNamespace(namespace string) *Deployment {
	d.ObjectMeta.Namespace = namespace
	return d
}

func (d *Deployment) WithLabels(labels map[string]string) *Deployment {
	d.ObjectMeta.Labels = labels
	d.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
	d.Spec.Template.ObjectMeta.Labels = labels

	return d
}

func (d *Deployment) WithReplicas(replicas int) *Deployment {
	int32Replicas := int32(replicas)
	d.Spec.Replicas = &int32Replicas
	return d
}

func (d *Deployment) AppendContainer(container *Container) *Deployment {
	d.Spec.Template.Spec.Containers = append(d.Spec.Template.Spec.Containers, container.Container)
	return d
}

func (d *Deployment) AppendVolume(volumeName, claimName string) *Deployment {
	d.Spec.Template.Spec.Volumes = append(d.Spec.Template.Spec.Volumes, corev1.Volume{
		Name: volumeName,
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: claimName,
			},
		},
	})
	return d
}

func (d *Deployment) AppendToleration(toleration corev1.Toleration) *Deployment {
	d.Spec.Template.Spec.Tolerations = append(d.Spec.Template.Spec.Tolerations, toleration)
	return d
}

func (d *Deployment) WithNodeSelector(ns map[string]string) *Deployment {
	d.Spec.Template.Spec.NodeSelector = ns
	return d
}

func (d *Deployment) WithStrategy(strategy appsv1.DeploymentStrategyType) *Deployment {
	d.Spec.Strategy.Type = strategy
	return d
}
