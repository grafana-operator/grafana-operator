// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grafana) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfig) DeepCopyInto(out *GrafanaConfig) {
	*out = *in
	out.Paths = in.Paths
	out.Server = in.Server
	out.Database = in.Database
	out.RemoteCache = in.RemoteCache
	out.Security = in.Security
	out.Users = in.Users
	out.Auth = in.Auth
	out.AuthBasic = in.AuthBasic
	out.AuthAnonymous = in.AuthAnonymous
	out.AuthGoogle = in.AuthGoogle
	out.AuthGithub = in.AuthGithub
	out.AuthGenericOauth = in.AuthGenericOauth
	out.AuthLdap = in.AuthLdap
	out.AuthProxy = in.AuthProxy
	out.DataProxy = in.DataProxy
	out.Analytics = in.Analytics
	out.Dashboards = in.Dashboards
	out.Smtp = in.Smtp
	out.Log = in.Log
	out.Metrics = in.Metrics
	out.MetricsGraphite = in.MetricsGraphite
	out.Snapshots = in.Snapshots
	out.ExternalImageStorage = in.ExternalImageStorage
	out.ExternalImageStorageS3 = in.ExternalImageStorageS3
	out.ExternalImageStorageWebdav = in.ExternalImageStorageWebdav
	out.ExternalImageStorageGcs = in.ExternalImageStorageGcs
	out.ExternalImageStorageAzureBlob = in.ExternalImageStorageAzureBlob
	out.Alerting = in.Alerting
	out.Panels = in.Panels
	out.Plugins = in.Plugins
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfig.
func (in *GrafanaConfig) DeepCopy() *GrafanaConfig {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAlerting) DeepCopyInto(out *GrafanaConfigAlerting) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAlerting.
func (in *GrafanaConfigAlerting) DeepCopy() *GrafanaConfigAlerting {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAlerting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAnalytics) DeepCopyInto(out *GrafanaConfigAnalytics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAnalytics.
func (in *GrafanaConfigAnalytics) DeepCopy() *GrafanaConfigAnalytics {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAnalytics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuth) DeepCopyInto(out *GrafanaConfigAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuth.
func (in *GrafanaConfigAuth) DeepCopy() *GrafanaConfigAuth {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthAnonymous) DeepCopyInto(out *GrafanaConfigAuthAnonymous) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthAnonymous.
func (in *GrafanaConfigAuthAnonymous) DeepCopy() *GrafanaConfigAuthAnonymous {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthAnonymous)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthBasic) DeepCopyInto(out *GrafanaConfigAuthBasic) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthBasic.
func (in *GrafanaConfigAuthBasic) DeepCopy() *GrafanaConfigAuthBasic {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthBasic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthGenericOauth) DeepCopyInto(out *GrafanaConfigAuthGenericOauth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthGenericOauth.
func (in *GrafanaConfigAuthGenericOauth) DeepCopy() *GrafanaConfigAuthGenericOauth {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthGenericOauth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthGithub) DeepCopyInto(out *GrafanaConfigAuthGithub) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthGithub.
func (in *GrafanaConfigAuthGithub) DeepCopy() *GrafanaConfigAuthGithub {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthGitlab) DeepCopyInto(out *GrafanaConfigAuthGitlab) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthGitlab.
func (in *GrafanaConfigAuthGitlab) DeepCopy() *GrafanaConfigAuthGitlab {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthGitlab)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthGoogle) DeepCopyInto(out *GrafanaConfigAuthGoogle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthGoogle.
func (in *GrafanaConfigAuthGoogle) DeepCopy() *GrafanaConfigAuthGoogle {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthLdap) DeepCopyInto(out *GrafanaConfigAuthLdap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthLdap.
func (in *GrafanaConfigAuthLdap) DeepCopy() *GrafanaConfigAuthLdap {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthLdap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthProxy) DeepCopyInto(out *GrafanaConfigAuthProxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthProxy.
func (in *GrafanaConfigAuthProxy) DeepCopy() *GrafanaConfigAuthProxy {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigDashboards) DeepCopyInto(out *GrafanaConfigDashboards) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigDashboards.
func (in *GrafanaConfigDashboards) DeepCopy() *GrafanaConfigDashboards {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigDashboards)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigDataProxy) DeepCopyInto(out *GrafanaConfigDataProxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigDataProxy.
func (in *GrafanaConfigDataProxy) DeepCopy() *GrafanaConfigDataProxy {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigDataProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigDatabase) DeepCopyInto(out *GrafanaConfigDatabase) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigDatabase.
func (in *GrafanaConfigDatabase) DeepCopy() *GrafanaConfigDatabase {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigExternalImageStorage) DeepCopyInto(out *GrafanaConfigExternalImageStorage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigExternalImageStorage.
func (in *GrafanaConfigExternalImageStorage) DeepCopy() *GrafanaConfigExternalImageStorage {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigExternalImageStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigExternalImageStorageAzureBlob) DeepCopyInto(out *GrafanaConfigExternalImageStorageAzureBlob) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigExternalImageStorageAzureBlob.
func (in *GrafanaConfigExternalImageStorageAzureBlob) DeepCopy() *GrafanaConfigExternalImageStorageAzureBlob {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigExternalImageStorageAzureBlob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigExternalImageStorageGcs) DeepCopyInto(out *GrafanaConfigExternalImageStorageGcs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigExternalImageStorageGcs.
func (in *GrafanaConfigExternalImageStorageGcs) DeepCopy() *GrafanaConfigExternalImageStorageGcs {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigExternalImageStorageGcs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigExternalImageStorageS3) DeepCopyInto(out *GrafanaConfigExternalImageStorageS3) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigExternalImageStorageS3.
func (in *GrafanaConfigExternalImageStorageS3) DeepCopy() *GrafanaConfigExternalImageStorageS3 {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigExternalImageStorageS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigExternalImageStorageWebdav) DeepCopyInto(out *GrafanaConfigExternalImageStorageWebdav) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigExternalImageStorageWebdav.
func (in *GrafanaConfigExternalImageStorageWebdav) DeepCopy() *GrafanaConfigExternalImageStorageWebdav {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigExternalImageStorageWebdav)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigLog) DeepCopyInto(out *GrafanaConfigLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigLog.
func (in *GrafanaConfigLog) DeepCopy() *GrafanaConfigLog {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigMetrics) DeepCopyInto(out *GrafanaConfigMetrics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigMetrics.
func (in *GrafanaConfigMetrics) DeepCopy() *GrafanaConfigMetrics {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigMetricsGraphite) DeepCopyInto(out *GrafanaConfigMetricsGraphite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigMetricsGraphite.
func (in *GrafanaConfigMetricsGraphite) DeepCopy() *GrafanaConfigMetricsGraphite {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigMetricsGraphite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigPanels) DeepCopyInto(out *GrafanaConfigPanels) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigPanels.
func (in *GrafanaConfigPanels) DeepCopy() *GrafanaConfigPanels {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigPanels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigPaths) DeepCopyInto(out *GrafanaConfigPaths) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigPaths.
func (in *GrafanaConfigPaths) DeepCopy() *GrafanaConfigPaths {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigPaths)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigPlugins) DeepCopyInto(out *GrafanaConfigPlugins) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigPlugins.
func (in *GrafanaConfigPlugins) DeepCopy() *GrafanaConfigPlugins {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigPlugins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigRemoteCache) DeepCopyInto(out *GrafanaConfigRemoteCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigRemoteCache.
func (in *GrafanaConfigRemoteCache) DeepCopy() *GrafanaConfigRemoteCache {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigRemoteCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigSecurity) DeepCopyInto(out *GrafanaConfigSecurity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigSecurity.
func (in *GrafanaConfigSecurity) DeepCopy() *GrafanaConfigSecurity {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigServer) DeepCopyInto(out *GrafanaConfigServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigServer.
func (in *GrafanaConfigServer) DeepCopy() *GrafanaConfigServer {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigSmtp) DeepCopyInto(out *GrafanaConfigSmtp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigSmtp.
func (in *GrafanaConfigSmtp) DeepCopy() *GrafanaConfigSmtp {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigSmtp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigSnapshots) DeepCopyInto(out *GrafanaConfigSnapshots) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigSnapshots.
func (in *GrafanaConfigSnapshots) DeepCopy() *GrafanaConfigSnapshots {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigSnapshots)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigUsers) DeepCopyInto(out *GrafanaConfigUsers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigUsers.
func (in *GrafanaConfigUsers) DeepCopy() *GrafanaConfigUsers {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigUsers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboard) DeepCopyInto(out *GrafanaDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboard.
func (in *GrafanaDashboard) DeepCopy() *GrafanaDashboard {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardList) DeepCopyInto(out *GrafanaDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardList.
func (in *GrafanaDashboardList) DeepCopy() *GrafanaDashboardList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardSpec) DeepCopyInto(out *GrafanaDashboardSpec) {
	*out = *in
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardSpec.
func (in *GrafanaDashboardSpec) DeepCopy() *GrafanaDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardSpecFields) DeepCopyInto(out *GrafanaDashboardSpecFields) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardSpecFields.
func (in *GrafanaDashboardSpecFields) DeepCopy() *GrafanaDashboardSpecFields {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardSpecFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardStatus) DeepCopyInto(out *GrafanaDashboardStatus) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]GrafanaDashboardStatusMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardStatus.
func (in *GrafanaDashboardStatus) DeepCopy() *GrafanaDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardStatusMessage) DeepCopyInto(out *GrafanaDashboardStatusMessage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardStatusMessage.
func (in *GrafanaDashboardStatusMessage) DeepCopy() *GrafanaDashboardStatusMessage {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardStatusMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSource) DeepCopyInto(out *GrafanaDataSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSource.
func (in *GrafanaDataSource) DeepCopy() *GrafanaDataSource {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDataSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceJsonData) DeepCopyInto(out *GrafanaDataSourceJsonData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceJsonData.
func (in *GrafanaDataSourceJsonData) DeepCopy() *GrafanaDataSourceJsonData {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceJsonData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceList) DeepCopyInto(out *GrafanaDataSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDataSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceList.
func (in *GrafanaDataSourceList) DeepCopy() *GrafanaDataSourceList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDataSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceSecureJsonData) DeepCopyInto(out *GrafanaDataSourceSecureJsonData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceSecureJsonData.
func (in *GrafanaDataSourceSecureJsonData) DeepCopy() *GrafanaDataSourceSecureJsonData {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceSecureJsonData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceSpec) DeepCopyInto(out *GrafanaDataSourceSpec) {
	*out = *in
	out.DataSource = in.DataSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceSpec.
func (in *GrafanaDataSourceSpec) DeepCopy() *GrafanaDataSourceSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceSpecFields) DeepCopyInto(out *GrafanaDataSourceSpecFields) {
	*out = *in
	out.JsonData = in.JsonData
	out.SecureJsonData = in.SecureJsonData
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceSpecFields.
func (in *GrafanaDataSourceSpecFields) DeepCopy() *GrafanaDataSourceSpecFields {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceSpecFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceStatus) DeepCopyInto(out *GrafanaDataSourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceStatus.
func (in *GrafanaDataSourceStatus) DeepCopy() *GrafanaDataSourceStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaIngress) DeepCopyInto(out *GrafanaIngress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaIngress.
func (in *GrafanaIngress) DeepCopy() *GrafanaIngress {
	if in == nil {
		return nil
	}
	out := new(GrafanaIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaList) DeepCopyInto(out *GrafanaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grafana, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaList.
func (in *GrafanaList) DeepCopy() *GrafanaList {
	if in == nil {
		return nil
	}
	out := new(GrafanaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaPlugin) DeepCopyInto(out *GrafanaPlugin) {
	*out = *in
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(GrafanaDashboard)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaPlugin.
func (in *GrafanaPlugin) DeepCopy() *GrafanaPlugin {
	if in == nil {
		return nil
	}
	out := new(GrafanaPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaService) DeepCopyInto(out *GrafanaService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaService.
func (in *GrafanaService) DeepCopy() *GrafanaService {
	if in == nil {
		return nil
	}
	out := new(GrafanaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaSpec) DeepCopyInto(out *GrafanaSpec) {
	*out = *in
	out.Config = in.Config
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DashboardLabelSelector != nil {
		in, out := &in.DashboardLabelSelector, &out.DashboardLabelSelector
		*out = make([]*metav1.LabelSelector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(metav1.LabelSelector)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Service.DeepCopyInto(&out.Service)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaSpec.
func (in *GrafanaSpec) DeepCopy() *GrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaStatus) DeepCopyInto(out *GrafanaStatus) {
	*out = *in
	if in.InstalledPlugins != nil {
		in, out := &in.InstalledPlugins, &out.InstalledPlugins
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailedPlugins != nil {
		in, out := &in.FailedPlugins, &out.FailedPlugins
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaStatus.
func (in *GrafanaStatus) DeepCopy() *GrafanaStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PluginList) DeepCopyInto(out *PluginList) {
	{
		in := &in
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginList.
func (in PluginList) DeepCopy() PluginList {
	if in == nil {
		return nil
	}
	out := new(PluginList)
	in.DeepCopyInto(out)
	return *out
}
