package k8s

import (
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// PrintYAMLDocuments uses client-go serializer to print YAML k8s manifests
func PrintYAMLDocuments(objects ...runtime.Object) {
	for _, obj := range objects {
		fmt.Println("---")
		k8sjson.NewYAMLSerializer(k8sjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme).Encode(obj, os.Stdout)
	}
}

// EnvVar is a convenience function for building EnvVars
func EnvVar(name, value string) corev1.EnvVar {
	return corev1.EnvVar{Name: name, Value: value}
}

// EnvVarFromSecret is a convenience function for building an EnvVar from a SecretKeyRef
func EnvVarFromSecret(name, secret, key string) corev1.EnvVar {
	// default to environment variable name if key is empty
	if key == "" {
		key = name
	}

	return corev1.EnvVar{
		Name: name,
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				Key: key,
				LocalObjectReference: corev1.LocalObjectReference{
					Name: secret},
			},
		},
	}
}

func getCoreV1TypeMeta(kind string) metav1.TypeMeta {
	return metav1.TypeMeta{Kind: kind, APIVersion: corev1.SchemeGroupVersion.String()}
}

func getExtensionsV1beta1TypeMeta(kind string) metav1.TypeMeta {
	return metav1.TypeMeta{Kind: kind, APIVersion: extv1beta1.SchemeGroupVersion.String()}
}
