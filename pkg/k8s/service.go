package k8s

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Service struct {
	corev1.Service
}

func NewService(name string) *Service {
	svc := &Service{}
	svc.TypeMeta = getCoreV1TypeMeta("Service")
	svc.ObjectMeta.Name = name
	return svc
}

func (svc *Service) WithNamespace(namespace string) *Service {
	svc.ObjectMeta.Namespace = namespace
	return svc
}

func (svc *Service) AppendNamedPort(name string, port, targetPort int) *Service {
	svc.Spec.Ports = append(svc.Spec.Ports, corev1.ServicePort{Name: name, Port: int32(port), TargetPort: intstr.FromInt(targetPort)})
	return svc
}

func (svc *Service) WithLabels(labels map[string]string) *Service {
	svc.ObjectMeta.Labels = labels
	return svc
}

func (svc *Service) WithLabelSelector(labels map[string]string) *Service {
	svc.Spec.Selector = labels
	return svc
}

func (svc *Service) WithExternalName(name string) *Service {
	svc.Spec.Type = corev1.ServiceTypeExternalName
	svc.Spec.ExternalName = name
	return svc
}
