package k8s

import (
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type IngressRule struct {
	extv1beta1.IngressRule
}

func NewIngressRule(host string) *IngressRule {
	r := &IngressRule{}
	r.Host = host
	return r
}

func (r *IngressRule) AppendHTTPPath(path, service, port string) *IngressRule {
	if r.IngressRuleValue.HTTP == nil {
		r.IngressRuleValue.HTTP = &extv1beta1.HTTPIngressRuleValue{}
	}

	r.IngressRuleValue.HTTP.Paths = append(
		r.IngressRuleValue.HTTP.Paths,
		extv1beta1.HTTPIngressPath{
			Path: path,
			Backend: extv1beta1.IngressBackend{
				ServiceName: service,
				ServicePort: intstr.FromString(port),
			},
		},
	)

	return r
}

type Ingress struct {
	extv1beta1.Ingress
}

func NewIngress(name string) *Ingress {
	ing := &Ingress{}
	ing.TypeMeta = getExtensionsV1beta1TypeMeta("Ingress")
	ing.ObjectMeta.Name = name
	return ing
}

func (ing *Ingress) WithNamespace(namespace string) *Ingress {
	ing.ObjectMeta.Namespace = namespace
	return ing
}

func (ing *Ingress) AppendIngressRule(rule *IngressRule) *Ingress {
	ing.Spec.Rules = append(ing.Spec.Rules, rule.IngressRule)
	return ing
}

func (ing *Ingress) WithLabels(labels map[string]string) *Ingress {
	ing.ObjectMeta.Labels = labels
	return ing
}
