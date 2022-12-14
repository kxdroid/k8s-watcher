package crd

// source: https://github.com/kubefirst/kubefirst-watcher-operator/blob/main/watcher/api/v1beta1/watcher_types.go

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WatcherSpec defines the desired state of Watcher
type WatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Quantity of instances
	Exit       int32                         `json:"exit"`
	Timeout    int32                         `json:"timeout"`
	ConfigMaps []BasicConfigurationCondition `json:"configmaps,omitempty"`
	Secrets    []BasicConfigurationCondition `json:"secrets,omitempty"`
	Services   []BasicConfigurationCondition `json:"services,omitempty"`
	Pods       []PodCondition                `json:"pods,omitempty"`
	Jobs       []JobCondition                `json:"jobs,omitempty"`
}

// BasicConfigurationCondition general match rules
type BasicConfigurationCondition struct {
	ID         int               `json:"id,omitempty"`
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	APIVersion string            `json:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// PodCondition pod matching rules
type PodCondition struct {
	ID         int               `json:"id,omitempty"`
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Phase      string            `json:"phase,omitempty"`
	APIVersion string            `json:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// JobCondition pod matching rules
type JobCondition struct {
	ID         int               `json:"id,omitempty"`
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Phase      string            `json:"phase,omitempty"`
	APIVersion string            `json:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
	// The number of pending and running pods.
	// +optional
	Active int32 `json:"active,omitempty"`

	// The number of pods which reached phase Succeeded.
	// +optional
	Succeeded int32 `json:"succeeded,omitempty"`

	// The number of pods which reached phase Failed.
	// +optional
	Failed int32 `json:"failed,omitempty"`
}

// WatcherStatus defines the observed state of Watcher
type WatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status    string `json:"status"`
	Instanced bool   `json:"instanced"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// Watcher is the Schema for the watchers API
type Watcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WatcherSpec   `json:"spec,omitempty"`
	Status WatcherStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WatcherList contains a list of Watcher
type WatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Watcher `json:"items"`
}
