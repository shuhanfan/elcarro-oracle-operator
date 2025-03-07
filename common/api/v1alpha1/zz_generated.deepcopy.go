// +build !ignore_autogenerated

/*
Copyright 2021 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]DiskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostAntiAffinityNamespaces != nil {
		in, out := &in.HostAntiAffinityNamespaces, &out.HostAntiAffinityNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSpec) DeepCopyInto(out *CredentialSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.GsmSecretRef != nil {
		in, out := &in.GsmSecretRef, &out.GsmSecretRef
		*out = new(GsmSecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSpec.
func (in *CredentialSpec) DeepCopy() *CredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronAnythingSpec) DeepCopyInto(out *CronAnythingSpec) {
	*out = *in
	if in.TriggerDeadlineSeconds != nil {
		in, out := &in.TriggerDeadlineSeconds, &out.TriggerDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.FinishableStrategy != nil {
		in, out := &in.FinishableStrategy, &out.FinishableStrategy
		*out = new(FinishableStrategy)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.TotalResourceLimit != nil {
		in, out := &in.TotalResourceLimit, &out.TotalResourceLimit
		*out = new(int32)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(ResourceRetention)
		(*in).DeepCopyInto(*out)
	}
	if in.CascadeDelete != nil {
		in, out := &in.CascadeDelete, &out.CascadeDelete
		*out = new(bool)
		**out = **in
	}
	if in.ResourceBaseName != nil {
		in, out := &in.ResourceBaseName, &out.ResourceBaseName
		*out = new(string)
		**out = **in
	}
	if in.ResourceTimestampFormat != nil {
		in, out := &in.ResourceTimestampFormat, &out.ResourceTimestampFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronAnythingSpec.
func (in *CronAnythingSpec) DeepCopy() *CronAnythingSpec {
	if in == nil {
		return nil
	}
	out := new(CronAnythingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronAnythingStatus) DeepCopyInto(out *CronAnythingStatus) {
	*out = *in
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.TriggerHistory != nil {
		in, out := &in.TriggerHistory, &out.TriggerHistory
		*out = make([]TriggerHistoryRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PendingTrigger != nil {
		in, out := &in.PendingTrigger, &out.PendingTrigger
		*out = new(PendingTrigger)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronAnythingStatus.
func (in *CronAnythingStatus) DeepCopy() *CronAnythingStatus {
	if in == nil {
		return nil
	}
	out := new(CronAnythingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldResourceTimestampStrategy) DeepCopyInto(out *FieldResourceTimestampStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldResourceTimestampStrategy.
func (in *FieldResourceTimestampStrategy) DeepCopy() *FieldResourceTimestampStrategy {
	if in == nil {
		return nil
	}
	out := new(FieldResourceTimestampStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinishableStrategy) DeepCopyInto(out *FinishableStrategy) {
	*out = *in
	if in.TimestampField != nil {
		in, out := &in.TimestampField, &out.TimestampField
		*out = new(TimestampFieldStrategy)
		**out = **in
	}
	if in.StringField != nil {
		in, out := &in.StringField, &out.StringField
		*out = new(StringFieldStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinishableStrategy.
func (in *FinishableStrategy) DeepCopy() *FinishableStrategy {
	if in == nil {
		return nil
	}
	out := new(FinishableStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GsmSecretReference) DeepCopyInto(out *GsmSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GsmSecretReference.
func (in *GsmSecretReference) DeepCopy() *GsmSecretReference {
	if in == nil {
		return nil
	}
	out := new(GsmSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]DiskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceCidrRanges != nil {
		in, out := &in.SourceCidrRanges, &out.SourceCidrRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Patching != nil {
		in, out := &in.Patching, &out.Patching
		*out = new(PatchingSpec)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[Service]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.DatabaseResources.DeepCopyInto(&out.DatabaseResources)
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowSpec) DeepCopyInto(out *MaintenanceWindowSpec) {
	*out = *in
	if in.TimeRanges != nil {
		in, out := &in.TimeRanges, &out.TimeRanges
		*out = make([]TimeRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowSpec.
func (in *MaintenanceWindowSpec) DeepCopy() *MaintenanceWindowSpec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchingSpec) DeepCopyInto(out *PatchingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchingSpec.
func (in *PatchingSpec) DeepCopy() *PatchingSpec {
	if in == nil {
		return nil
	}
	out := new(PatchingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingTrigger) DeepCopyInto(out *PendingTrigger) {
	*out = *in
	in.ScheduleTime.DeepCopyInto(&out.ScheduleTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingTrigger.
func (in *PendingTrigger) DeepCopy() *PendingTrigger {
	if in == nil {
		return nil
	}
	out := new(PendingTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRetention) DeepCopyInto(out *ResourceRetention) {
	*out = *in
	if in.HistoryCountLimit != nil {
		in, out := &in.HistoryCountLimit, &out.HistoryCountLimit
		*out = new(int32)
		**out = **in
	}
	if in.HistoryTimeLimitSeconds != nil {
		in, out := &in.HistoryTimeLimitSeconds, &out.HistoryTimeLimitSeconds
		*out = new(uint64)
		**out = **in
	}
	in.ResourceTimestampStrategy.DeepCopyInto(&out.ResourceTimestampStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRetention.
func (in *ResourceRetention) DeepCopy() *ResourceRetention {
	if in == nil {
		return nil
	}
	out := new(ResourceRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTimestampStrategy) DeepCopyInto(out *ResourceTimestampStrategy) {
	*out = *in
	if in.FieldResourceTimestampStrategy != nil {
		in, out := &in.FieldResourceTimestampStrategy, &out.FieldResourceTimestampStrategy
		*out = new(FieldResourceTimestampStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTimestampStrategy.
func (in *ResourceTimestampStrategy) DeepCopy() *ResourceTimestampStrategy {
	if in == nil {
		return nil
	}
	out := new(ResourceTimestampStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringFieldStrategy) DeepCopyInto(out *StringFieldStrategy) {
	*out = *in
	if in.FinishedValues != nil {
		in, out := &in.FinishedValues, &out.FinishedValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringFieldStrategy.
func (in *StringFieldStrategy) DeepCopy() *StringFieldStrategy {
	if in == nil {
		return nil
	}
	out := new(StringFieldStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = (*in).DeepCopy()
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestampFieldStrategy) DeepCopyInto(out *TimestampFieldStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestampFieldStrategy.
func (in *TimestampFieldStrategy) DeepCopy() *TimestampFieldStrategy {
	if in == nil {
		return nil
	}
	out := new(TimestampFieldStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerHistoryRecord) DeepCopyInto(out *TriggerHistoryRecord) {
	*out = *in
	in.ScheduleTime.DeepCopyInto(&out.ScheduleTime)
	in.CreationTimestamp.DeepCopyInto(&out.CreationTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerHistoryRecord.
func (in *TriggerHistoryRecord) DeepCopy() *TriggerHistoryRecord {
	if in == nil {
		return nil
	}
	out := new(TriggerHistoryRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.CredentialSpec.DeepCopyInto(&out.CredentialSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}
