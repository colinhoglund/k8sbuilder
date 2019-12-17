package k8s

import (
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
)

type PersistentVolumeClaim struct {
	corev1.PersistentVolumeClaim
}

func NewPersistentVolumeClaim(name string) *PersistentVolumeClaim {
	pvc := &PersistentVolumeClaim{}
	pvc.TypeMeta = getCoreV1TypeMeta("PersistentVolumeClaim")
	pvc.ObjectMeta.Name = name

	return pvc
}

func (pvc *PersistentVolumeClaim) WithNamespace(namespace string) *PersistentVolumeClaim {
	pvc.ObjectMeta.Namespace = namespace
	return pvc
}

func (pvc *PersistentVolumeClaim) WithLabels(labels map[string]string) *PersistentVolumeClaim {
	pvc.ObjectMeta.Labels = labels

	return pvc
}

func (pvc *PersistentVolumeClaim) AppendAccessMode(pvam corev1.PersistentVolumeAccessMode) *PersistentVolumeClaim {
	pvc.Spec.AccessModes = append(pvc.Spec.AccessModes, pvam)
	return pvc
}

func (pvc *PersistentVolumeClaim) WithStorage(size string) *PersistentVolumeClaim {
	if pvc.Spec.Resources.Requests == nil {
		pvc.Spec.Resources.Requests = make(corev1.ResourceList)
	}

	pvc.Spec.Resources.Requests[corev1.ResourceStorage] = resource.MustParse(size)
	return pvc
}
