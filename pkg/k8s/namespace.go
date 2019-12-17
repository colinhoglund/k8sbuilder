package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

type Namespace struct {
	corev1.Namespace
}

func NewNamespace(name string) *Namespace {
	ns := &Namespace{}
	ns.TypeMeta = getCoreV1TypeMeta("Namespace")
	ns.ObjectMeta.Name = name
	return ns
}
