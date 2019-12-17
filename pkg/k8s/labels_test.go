package k8s

import (
	"reflect"
	"testing"
)

func TestCommonLabels(t *testing.T) {
	cl := CommonLabels{
		Name:      "myapp",
		Instance:  "myapp-qwerty",
		Version:   "0.1.0",
		Component: "server",
		PartOf:    "test",
		ManagedBy: "testing",
	}

	got := cl.Map()
	want := map[string]string{
		"app.kubernetes.io/name":       "myapp",
		"app.kubernetes.io/instance":   "myapp-qwerty",
		"app.kubernetes.io/version":    "0.1.0",
		"app.kubernetes.io/component":  "server",
		"app.kubernetes.io/part-of":    "test",
		"app.kubernetes.io/managed-by": "testing",
	}

	assertEqualMapStringString(t, got, want)
}

func TestCommonLabelsEmpty(t *testing.T) {
	cl := CommonLabels{Name: "myapp"}
	got := cl.Map()
	want := map[string]string{"app.kubernetes.io/name": "myapp"}
	assertEqualMapStringString(t, got, want)
}

func assertEqualMapStringString(t *testing.T, got, want map[string]string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\n\tgot: %v\n\twant: %v", got, want)
	}
}
