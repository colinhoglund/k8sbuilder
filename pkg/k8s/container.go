package k8s

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
)

type Container struct {
	corev1.Container
}

func NewContainer(name string) *Container {
	c := &Container{}
	c.Name = name
	return c
}

func (c *Container) WithImage(repo, tag string) *Container {
	c.Image = fmt.Sprintf("%s:%s", repo, tag)
	return c
}

func (c *Container) WithImagePullPolicy(policy corev1.PullPolicy) *Container {
	c.ImagePullPolicy = policy
	return c
}

func (c *Container) AppendEnvVar(name, value string) *Container {
	c.Env = append(c.Env, corev1.EnvVar{Name: name, Value: value})
	return c
}

func (c *Container) AppendEnvVarFromSecret(name, secret, key string) *Container {
	// default to environment variable name if key is empty
	if key == "" {
		key = name
	}

	c.Env = append(c.Env, corev1.EnvVar{
		Name: name,
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				Key: key,
				LocalObjectReference: corev1.LocalObjectReference{
					Name: secret},
			},
		},
	})
	return c
}

func (c *Container) WithLivenessProbe(path, port string) *Container {
	c.LivenessProbe = &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{Path: path, Port: intstr.FromString(port)},
		},
	}
	return c
}

func (c *Container) AppendNamedPort(name string, port int) *Container {
	c.Ports = append(c.Ports, corev1.ContainerPort{
		Name: name, ContainerPort: int32(port),
	})
	return c
}

func (c *Container) AppendVolumeMount(name, path string) *Container {
	c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{
		Name: name, MountPath: path,
	})
	return c
}

func (c *Container) WithArgs(command []string) *Container {
	c.Args = command
	return c
}

func (c *Container) WithResources(resources corev1.ResourceRequirements) *Container {
	c.Resources = resources
	return c
}
