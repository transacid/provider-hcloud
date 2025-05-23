// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BalancerServiceHTTPInitParameters struct {

	// List of IDs from certificates which the Load Balancer has.
	// +listType=set
	Certificates []*float64 `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Lifetime of the cookie for sticky session (in seconds). Default: 300
	CookieLifetime *float64 `json:"cookieLifetime,omitempty" tf:"cookie_lifetime,omitempty"`

	// Name of the cookie for sticky session. Default: HCLBSTICKY
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Redirect HTTP to HTTPS traffic. Only supported for services with protocol https using the default HTTP port 80.
	RedirectHTTP *bool `json:"redirectHttp,omitempty" tf:"redirect_http,omitempty"`

	// Enable sticky sessions
	StickySessions *bool `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`
}

type BalancerServiceHTTPObservation struct {

	// List of IDs from certificates which the Load Balancer has.
	// +listType=set
	Certificates []*float64 `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Lifetime of the cookie for sticky session (in seconds). Default: 300
	CookieLifetime *float64 `json:"cookieLifetime,omitempty" tf:"cookie_lifetime,omitempty"`

	// Name of the cookie for sticky session. Default: HCLBSTICKY
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Redirect HTTP to HTTPS traffic. Only supported for services with protocol https using the default HTTP port 80.
	RedirectHTTP *bool `json:"redirectHttp,omitempty" tf:"redirect_http,omitempty"`

	// Enable sticky sessions
	StickySessions *bool `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`
}

type BalancerServiceHTTPParameters struct {

	// List of IDs from certificates which the Load Balancer has.
	// +kubebuilder:validation:Optional
	// +listType=set
	Certificates []*float64 `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Lifetime of the cookie for sticky session (in seconds). Default: 300
	// +kubebuilder:validation:Optional
	CookieLifetime *float64 `json:"cookieLifetime,omitempty" tf:"cookie_lifetime,omitempty"`

	// Name of the cookie for sticky session. Default: HCLBSTICKY
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Redirect HTTP to HTTPS traffic. Only supported for services with protocol https using the default HTTP port 80.
	// +kubebuilder:validation:Optional
	RedirectHTTP *bool `json:"redirectHttp,omitempty" tf:"redirect_http,omitempty"`

	// Enable sticky sessions
	// +kubebuilder:validation:Optional
	StickySessions *bool `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`
}

type BalancerServiceInitParameters struct {

	// Port the service connects to the targets on, required if protocol is tcp. Can be everything between 1 and 65535.
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// HTTP configuration when protocol is http or https.
	HTTP []BalancerServiceHTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Health Check configuration when protocol is http or https.
	HealthCheck []HealthCheckInitParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Port the service listen on, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// Id of the load balancer this service belongs to.
	// +crossplane:generate:reference:type=github.com/transacid/provider-hcloud/apis/loadbalancer/v1alpha1.Balancer
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a Balancer in loadbalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a Balancer in loadbalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// Protocol of the service. http, https or tcp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Enable proxyprotocol.
	Proxyprotocol *bool `json:"proxyprotocol,omitempty" tf:"proxyprotocol,omitempty"`
}

type BalancerServiceObservation struct {

	// Port the service connects to the targets on, required if protocol is tcp. Can be everything between 1 and 65535.
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// HTTP configuration when protocol is http or https.
	HTTP []BalancerServiceHTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// Health Check configuration when protocol is http or https.
	HealthCheck []HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Port the service listen on, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// Id of the load balancer this service belongs to.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Protocol of the service. http, https or tcp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Enable proxyprotocol.
	Proxyprotocol *bool `json:"proxyprotocol,omitempty" tf:"proxyprotocol,omitempty"`
}

type BalancerServiceParameters struct {

	// Port the service connects to the targets on, required if protocol is tcp. Can be everything between 1 and 65535.
	// +kubebuilder:validation:Optional
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// HTTP configuration when protocol is http or https.
	// +kubebuilder:validation:Optional
	HTTP []BalancerServiceHTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Health Check configuration when protocol is http or https.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Port the service listen on, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	// +kubebuilder:validation:Optional
	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// Id of the load balancer this service belongs to.
	// +crossplane:generate:reference:type=github.com/transacid/provider-hcloud/apis/loadbalancer/v1alpha1.Balancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a Balancer in loadbalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a Balancer in loadbalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// Protocol of the service. http, https or tcp
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Enable proxyprotocol.
	// +kubebuilder:validation:Optional
	Proxyprotocol *bool `json:"proxyprotocol,omitempty" tf:"proxyprotocol,omitempty"`
}

type HTTPInitParameters struct {

	// Domain we try to access when performing the Health Check.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Path we try to access when performing the Health Check.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Response we expect to be included in the Target response when a Health Check was performed.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`

	// We expect that the target answers with these status codes. If not the target is marked as unhealthy.
	StatusCodes []*string `json:"statusCodes,omitempty" tf:"status_codes,omitempty"`

	// Enable TLS certificate checking.
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type HTTPObservation struct {

	// Domain we try to access when performing the Health Check.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Path we try to access when performing the Health Check.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Response we expect to be included in the Target response when a Health Check was performed.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`

	// We expect that the target answers with these status codes. If not the target is marked as unhealthy.
	StatusCodes []*string `json:"statusCodes,omitempty" tf:"status_codes,omitempty"`

	// Enable TLS certificate checking.
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type HTTPParameters struct {

	// Domain we try to access when performing the Health Check.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Path we try to access when performing the Health Check.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Response we expect to be included in the Target response when a Health Check was performed.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`

	// We expect that the target answers with these status codes. If not the target is marked as unhealthy.
	// +kubebuilder:validation:Optional
	StatusCodes []*string `json:"statusCodes,omitempty" tf:"status_codes,omitempty"`

	// Enable TLS certificate checking.
	// +kubebuilder:validation:Optional
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type HealthCheckInitParameters struct {

	// HTTP configuration when protocol is http or https.
	HTTP []HTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Interval how often the health check will be performed, in seconds.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Port the health check tries to connect to, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol of the service. http, https or tcp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Number of tries a health check will be performed until a target will be listed as unhealthy.
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// Timeout when a health check try will be canceled if there is no response, in seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type HealthCheckObservation struct {

	// HTTP configuration when protocol is http or https.
	HTTP []HTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// Interval how often the health check will be performed, in seconds.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Port the health check tries to connect to, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol of the service. http, https or tcp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Number of tries a health check will be performed until a target will be listed as unhealthy.
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// Timeout when a health check try will be canceled if there is no response, in seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type HealthCheckParameters struct {

	// HTTP configuration when protocol is http or https.
	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Interval how often the health check will be performed, in seconds.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// Port the health check tries to connect to, required if protocol is tcp. Can be everything between 1 and 65535. Must be unique per Load Balancer.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol of the service. http, https or tcp
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Number of tries a health check will be performed until a target will be listed as unhealthy.
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// Timeout when a health check try will be canceled if there is no response, in seconds.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`
}

// BalancerServiceSpec defines the desired state of BalancerService
type BalancerServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerServiceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BalancerServiceInitParameters `json:"initProvider,omitempty"`
}

// BalancerServiceStatus defines the observed state of BalancerService.
type BalancerServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BalancerService is the Schema for the BalancerServices API. Define services for Hetzner Cloud Load Balancers.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type BalancerService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   BalancerServiceSpec   `json:"spec"`
	Status BalancerServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerServiceList contains a list of BalancerServices
type BalancerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerService `json:"items"`
}

// Repository type metadata.
var (
	BalancerService_Kind             = "BalancerService"
	BalancerService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerService_Kind}.String()
	BalancerService_KindAPIVersion   = BalancerService_Kind + "." + CRDGroupVersion.String()
	BalancerService_GroupVersionKind = CRDGroupVersion.WithKind(BalancerService_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerService{}, &BalancerServiceList{})
}
