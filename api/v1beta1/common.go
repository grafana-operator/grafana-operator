package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ValueFrom struct {
	TargetPath string          `json:"targetPath"`
	ValueFrom  ValueFromSource `json:"valueFrom"`
}

// +kubebuilder:validation:XValidation:rule="(has(self.configMapKeyRef) && !has(self.secretKeyRef)) || (!has(self.configMapKeyRef) && has(self.secretKeyRef))", message="Either configMapKeyRef or secretKeyRef must be set"
type ValueFromSource struct {
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *v1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a Secret.
	// +optional
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// Common Options that all CRs should embed, excluding GrafanaSpec
// Ensure alignment on handling ResyncPeriod, InstanceSelector, and AllowCrossNamespaceImport
type GrafanaCommonSpec struct {
	// How often the resource is synced, defaults to 10m0s if not set
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=duration
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$"
	// +kubebuilder:default="10m0s"
	ResyncPeriod metav1.Duration `json:"resyncPeriod,omitempty"`

	// Selects Grafana instances for import
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.instanceSelector is immutable"
	InstanceSelector *metav1.LabelSelector `json:"instanceSelector"`

	// Allow the Operator to match this resource with Grafanas outside the current namespace
	// +optional
	// +kubebuilder:default=false
	AllowCrossNamespaceImport bool `json:"allowCrossNamespaceImport,omitempty"`
}

// Common Functions that all CRs should implement, excluding Grafana
// +kubebuilder:object:generate=false
type CommonResource interface {
	MatchLabels() *metav1.LabelSelector
	MatchNamespace() string
	AllowCrossNamespace() bool
	ResyncPeriodHasElapsed() bool
}
