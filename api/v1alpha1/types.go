package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

type SSHGatewayClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SSHGatewayClassSpec   `json:"spec"`
	Status            SSHGatewayClassStatus `json:"status,omitempty"`
}

type SSHGatewayClassSpec struct {
	ControllerName string                `json:"controllerName"`
	Description    *string               `json:"description,omitempty"`
	ParametersRef  *ParametersReference  `json:"parametersRef,omitempty"`
}

type ParametersReference struct {
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

type SSHGatewayClassStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

type SSHGatewayClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHGatewayClass `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type SSHGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SSHGatewaySpec   `json:"spec"`
	Status            SSHGatewayStatus `json:"status,omitempty"`
}

type SSHGatewaySpec struct {
	GatewayClassName string        `json:"gatewayClassName"`
	Listeners        []SSHListener `json:"listeners"`
}

type SSHListener struct {
	Name     string          `json:"name"`
	Port     int32           `json:"port"`
	Protocol SSHProtocol     `json:"protocol,omitempty"`
	HostKey  *SecretRef      `json:"hostKey,omitempty"`
	Auth     SSHAuthConfig   `json:"auth"`
}

// +kubebuilder:validation:Enum=SSH
type SSHProtocol string

const (
	SSHProtocolSSH SSHProtocol = "SSH"
)

type SSHAuthConfig struct {
	Methods    []SSHAuthMethod `json:"methods"`
	MaxRetries *int32          `json:"maxRetries,omitempty"`
}

// +kubebuilder:validation:Enum=Password;PublicKey;Keyboard
type SSHAuthMethod string

const (
	SSHAuthPassword SSHAuthMethod = "Password"
	SSHAuthPublicKey SSHAuthMethod = "PublicKey"
	SSHAuthKeyboard SSHAuthMethod = "Keyboard"
)

type SecretRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

type SSHGatewayStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	Listeners  []ListenerStatus   `json:"listeners,omitempty"`
	Addresses  []GatewayAddress   `json:"addresses,omitempty"`
}

type ListenerStatus struct {
	Name       string             `json:"name"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	AttachedRoutes int32          `json:"attachedRoutes"`
}

type GatewayAddress struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value"`
}

// +kubebuilder:object:root=true

type SSHGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHGateway `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type SSHRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SSHRouteSpec   `json:"spec"`
	Status            SSHRouteStatus `json:"status,omitempty"`
}

type SSHRouteSpec struct {
	ParentRefs []ParentReference `json:"parentRefs"`
	Rules      []SSHRouteRule    `json:"rules"`
}

type ParentReference struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	SectionName string `json:"sectionName,omitempty"`
	Port        *int32 `json:"port,omitempty"`
}

type SSHRouteRule struct {
	Matches    []SSHRouteMatch   `json:"matches,omitempty"`
	BackendRefs []SSHBackendRef  `json:"backendRefs,omitempty"`
	Filters    []SSHRouteFilter  `json:"filters,omitempty"`
}

type SSHRouteMatch struct {
	Users    []string          `json:"users,omitempty"`
	Commands []CommandMatch    `json:"commands,omitempty"`
}

type CommandMatch struct {
	Type  CommandMatchType `json:"type,omitempty"`
	Value string           `json:"value"`
}

// +kubebuilder:validation:Enum=Exact;Prefix;Regex
type CommandMatchType string

const (
	CommandMatchExact  CommandMatchType = "Exact"
	CommandMatchPrefix CommandMatchType = "Prefix"
	CommandMatchRegex  CommandMatchType = "Regex"
)

type SSHBackendRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Port      *int32 `json:"port,omitempty"`
	Weight    *int32 `json:"weight,omitempty"`
}

type SSHRouteFilter struct {
	Type             SSHRouteFilterType    `json:"type"`
	CommandTemplate  *CommandTemplateFilter `json:"commandTemplate,omitempty"`
	RecordSession    *RecordSessionFilter   `json:"recordSession,omitempty"`
}

// +kubebuilder:validation:Enum=CommandTemplate;RecordSession
type SSHRouteFilterType string

const (
	SSHRouteFilterCommandTemplate SSHRouteFilterType = "CommandTemplate"
	SSHRouteFilterRecordSession   SSHRouteFilterType = "RecordSession"
)

type CommandTemplateFilter struct {
	AllowedCommands []string `json:"allowedCommands"`
	DeniedCommands  []string `json:"deniedCommands,omitempty"`
}

type RecordSessionFilter struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path,omitempty"`
}

type SSHRouteStatus struct {
	Parents []RouteParentStatus `json:"parents,omitempty"`
}

type RouteParentStatus struct {
	ParentRef  ParentReference    `json:"parentRef"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

type SSHRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHRoute `json:"items"`
}
