//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package grafanauserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthAzuread) DeepCopyInto(out *AuthAzuread) {
	*out = *in
	if in.AllowSignUp != nil {
		in, out := &in.AllowSignUp, &out.AllowSignUp
		*out = new(bool)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedGroups != nil {
		in, out := &in.AllowedGroups, &out.AllowedGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthAzuread.
func (in *AuthAzuread) DeepCopy() *AuthAzuread {
	if in == nil {
		return nil
	}
	out := new(AuthAzuread)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthGenericOauth) DeepCopyInto(out *AuthGenericOauth) {
	*out = *in
	if in.AllowSignUp != nil {
		in, out := &in.AllowSignUp, &out.AllowSignUp
		*out = new(bool)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOrganizations != nil {
		in, out := &in.AllowedOrganizations, &out.AllowedOrganizations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoLogin != nil {
		in, out := &in.AutoLogin, &out.AutoLogin
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthGenericOauth.
func (in *AuthGenericOauth) DeepCopy() *AuthGenericOauth {
	if in == nil {
		return nil
	}
	out := new(AuthGenericOauth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthGithub) DeepCopyInto(out *AuthGithub) {
	*out = *in
	if in.AllowSignUp != nil {
		in, out := &in.AllowSignUp, &out.AllowSignUp
		*out = new(bool)
		**out = **in
	}
	if in.AllowedOrganizations != nil {
		in, out := &in.AllowedOrganizations, &out.AllowedOrganizations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TeamIds != nil {
		in, out := &in.TeamIds, &out.TeamIds
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthGithub.
func (in *AuthGithub) DeepCopy() *AuthGithub {
	if in == nil {
		return nil
	}
	out := new(AuthGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthGitlab) DeepCopyInto(out *AuthGitlab) {
	*out = *in
	if in.AllowSignUp != nil {
		in, out := &in.AllowSignUp, &out.AllowSignUp
		*out = new(bool)
		**out = **in
	}
	if in.AllowedGroups != nil {
		in, out := &in.AllowedGroups, &out.AllowedGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ApiUrl != nil {
		in, out := &in.ApiUrl, &out.ApiUrl
		*out = new(string)
		**out = **in
	}
	if in.AuthUrl != nil {
		in, out := &in.AuthUrl, &out.AuthUrl
		*out = new(string)
		**out = **in
	}
	if in.TokenUrl != nil {
		in, out := &in.TokenUrl, &out.TokenUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthGitlab.
func (in *AuthGitlab) DeepCopy() *AuthGitlab {
	if in == nil {
		return nil
	}
	out := new(AuthGitlab)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthGoogle) DeepCopyInto(out *AuthGoogle) {
	*out = *in
	if in.AllowSignUp != nil {
		in, out := &in.AllowSignUp, &out.AllowSignUp
		*out = new(bool)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthGoogle.
func (in *AuthGoogle) DeepCopy() *AuthGoogle {
	if in == nil {
		return nil
	}
	out := new(AuthGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DateFormats) DeepCopyInto(out *DateFormats) {
	*out = *in
	if in.DefaultTimezone != nil {
		in, out := &in.DefaultTimezone, &out.DefaultTimezone
		*out = new(string)
		**out = **in
	}
	if in.FullDate != nil {
		in, out := &in.FullDate, &out.FullDate
		*out = new(string)
		**out = **in
	}
	if in.IntervalDay != nil {
		in, out := &in.IntervalDay, &out.IntervalDay
		*out = new(string)
		**out = **in
	}
	if in.IntervalHour != nil {
		in, out := &in.IntervalHour, &out.IntervalHour
		*out = new(string)
		**out = **in
	}
	if in.IntervalMinute != nil {
		in, out := &in.IntervalMinute, &out.IntervalMinute
		*out = new(string)
		**out = **in
	}
	if in.IntervalMonth != nil {
		in, out := &in.IntervalMonth, &out.IntervalMonth
		*out = new(string)
		**out = **in
	}
	if in.IntervalSecond != nil {
		in, out := &in.IntervalSecond, &out.IntervalSecond
		*out = new(string)
		**out = **in
	}
	if in.IntervalYear != nil {
		in, out := &in.IntervalYear, &out.IntervalYear
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DateFormats.
func (in *DateFormats) DeepCopy() *DateFormats {
	if in == nil {
		return nil
	}
	out := new(DateFormats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalImageStorage) DeepCopyInto(out *ExternalImageStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalImageStorage.
func (in *ExternalImageStorage) DeepCopy() *ExternalImageStorage {
	if in == nil {
		return nil
	}
	out := new(ExternalImageStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaUserConfig) DeepCopyInto(out *GrafanaUserConfig) {
	*out = *in
	if in.AdditionalBackupRegions != nil {
		in, out := &in.AdditionalBackupRegions, &out.AdditionalBackupRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AlertingEnabled != nil {
		in, out := &in.AlertingEnabled, &out.AlertingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlertingErrorOrTimeout != nil {
		in, out := &in.AlertingErrorOrTimeout, &out.AlertingErrorOrTimeout
		*out = new(string)
		**out = **in
	}
	if in.AlertingMaxAnnotationsToKeep != nil {
		in, out := &in.AlertingMaxAnnotationsToKeep, &out.AlertingMaxAnnotationsToKeep
		*out = new(int)
		**out = **in
	}
	if in.AlertingNodataOrNullvalues != nil {
		in, out := &in.AlertingNodataOrNullvalues, &out.AlertingNodataOrNullvalues
		*out = new(string)
		**out = **in
	}
	if in.AllowEmbedding != nil {
		in, out := &in.AllowEmbedding, &out.AllowEmbedding
		*out = new(bool)
		**out = **in
	}
	if in.AuthAzuread != nil {
		in, out := &in.AuthAzuread, &out.AuthAzuread
		*out = new(AuthAzuread)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthBasicEnabled != nil {
		in, out := &in.AuthBasicEnabled, &out.AuthBasicEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthGenericOauth != nil {
		in, out := &in.AuthGenericOauth, &out.AuthGenericOauth
		*out = new(AuthGenericOauth)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthGithub != nil {
		in, out := &in.AuthGithub, &out.AuthGithub
		*out = new(AuthGithub)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthGitlab != nil {
		in, out := &in.AuthGitlab, &out.AuthGitlab
		*out = new(AuthGitlab)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthGoogle != nil {
		in, out := &in.AuthGoogle, &out.AuthGoogle
		*out = new(AuthGoogle)
		(*in).DeepCopyInto(*out)
	}
	if in.CookieSamesite != nil {
		in, out := &in.CookieSamesite, &out.CookieSamesite
		*out = new(string)
		**out = **in
	}
	if in.CustomDomain != nil {
		in, out := &in.CustomDomain, &out.CustomDomain
		*out = new(string)
		**out = **in
	}
	if in.DashboardPreviewsEnabled != nil {
		in, out := &in.DashboardPreviewsEnabled, &out.DashboardPreviewsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DashboardsMinRefreshInterval != nil {
		in, out := &in.DashboardsMinRefreshInterval, &out.DashboardsMinRefreshInterval
		*out = new(string)
		**out = **in
	}
	if in.DashboardsVersionsToKeep != nil {
		in, out := &in.DashboardsVersionsToKeep, &out.DashboardsVersionsToKeep
		*out = new(int)
		**out = **in
	}
	if in.DataproxySendUserHeader != nil {
		in, out := &in.DataproxySendUserHeader, &out.DataproxySendUserHeader
		*out = new(bool)
		**out = **in
	}
	if in.DataproxyTimeout != nil {
		in, out := &in.DataproxyTimeout, &out.DataproxyTimeout
		*out = new(int)
		**out = **in
	}
	if in.DateFormats != nil {
		in, out := &in.DateFormats, &out.DateFormats
		*out = new(DateFormats)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableGravatar != nil {
		in, out := &in.DisableGravatar, &out.DisableGravatar
		*out = new(bool)
		**out = **in
	}
	if in.EditorsCanAdmin != nil {
		in, out := &in.EditorsCanAdmin, &out.EditorsCanAdmin
		*out = new(bool)
		**out = **in
	}
	if in.ExternalImageStorage != nil {
		in, out := &in.ExternalImageStorage, &out.ExternalImageStorage
		*out = new(ExternalImageStorage)
		**out = **in
	}
	if in.GoogleAnalyticsUaId != nil {
		in, out := &in.GoogleAnalyticsUaId, &out.GoogleAnalyticsUaId
		*out = new(string)
		**out = **in
	}
	if in.IpFilter != nil {
		in, out := &in.IpFilter, &out.IpFilter
		*out = make([]*IpFilter, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IpFilter)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MetricsEnabled != nil {
		in, out := &in.MetricsEnabled, &out.MetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.OauthAllowInsecureEmailLookup != nil {
		in, out := &in.OauthAllowInsecureEmailLookup, &out.OauthAllowInsecureEmailLookup
		*out = new(bool)
		**out = **in
	}
	if in.PrivateAccess != nil {
		in, out := &in.PrivateAccess, &out.PrivateAccess
		*out = new(PrivateAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivatelinkAccess != nil {
		in, out := &in.PrivatelinkAccess, &out.PrivatelinkAccess
		*out = new(PrivatelinkAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectToForkFrom != nil {
		in, out := &in.ProjectToForkFrom, &out.ProjectToForkFrom
		*out = new(string)
		**out = **in
	}
	if in.PublicAccess != nil {
		in, out := &in.PublicAccess, &out.PublicAccess
		*out = new(PublicAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.RecoveryBasebackupName != nil {
		in, out := &in.RecoveryBasebackupName, &out.RecoveryBasebackupName
		*out = new(string)
		**out = **in
	}
	if in.ServiceToForkFrom != nil {
		in, out := &in.ServiceToForkFrom, &out.ServiceToForkFrom
		*out = new(string)
		**out = **in
	}
	if in.SmtpServer != nil {
		in, out := &in.SmtpServer, &out.SmtpServer
		*out = new(SmtpServer)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIps != nil {
		in, out := &in.StaticIps, &out.StaticIps
		*out = new(bool)
		**out = **in
	}
	if in.UserAutoAssignOrg != nil {
		in, out := &in.UserAutoAssignOrg, &out.UserAutoAssignOrg
		*out = new(bool)
		**out = **in
	}
	if in.UserAutoAssignOrgRole != nil {
		in, out := &in.UserAutoAssignOrgRole, &out.UserAutoAssignOrgRole
		*out = new(string)
		**out = **in
	}
	if in.ViewersCanEdit != nil {
		in, out := &in.ViewersCanEdit, &out.ViewersCanEdit
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaUserConfig.
func (in *GrafanaUserConfig) DeepCopy() *GrafanaUserConfig {
	if in == nil {
		return nil
	}
	out := new(GrafanaUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpFilter) DeepCopyInto(out *IpFilter) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpFilter.
func (in *IpFilter) DeepCopy() *IpFilter {
	if in == nil {
		return nil
	}
	out := new(IpFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateAccess) DeepCopyInto(out *PrivateAccess) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateAccess.
func (in *PrivateAccess) DeepCopy() *PrivateAccess {
	if in == nil {
		return nil
	}
	out := new(PrivateAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivatelinkAccess) DeepCopyInto(out *PrivatelinkAccess) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivatelinkAccess.
func (in *PrivatelinkAccess) DeepCopy() *PrivatelinkAccess {
	if in == nil {
		return nil
	}
	out := new(PrivatelinkAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicAccess) DeepCopyInto(out *PublicAccess) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicAccess.
func (in *PublicAccess) DeepCopy() *PublicAccess {
	if in == nil {
		return nil
	}
	out := new(PublicAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmtpServer) DeepCopyInto(out *SmtpServer) {
	*out = *in
	if in.FromName != nil {
		in, out := &in.FromName, &out.FromName
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.SkipVerify != nil {
		in, out := &in.SkipVerify, &out.SkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.StarttlsPolicy != nil {
		in, out := &in.StarttlsPolicy, &out.StarttlsPolicy
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmtpServer.
func (in *SmtpServer) DeepCopy() *SmtpServer {
	if in == nil {
		return nil
	}
	out := new(SmtpServer)
	in.DeepCopyInto(out)
	return out
}
