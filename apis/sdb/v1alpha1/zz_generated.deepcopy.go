//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabase) DeepCopyInto(out *SQLDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabase.
func (in *SQLDatabase) DeepCopy() *SQLDatabase {
	if in == nil {
		return nil
	}
	out := new(SQLDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseInitParameters) DeepCopyInto(out *SQLDatabaseInitParameters) {
	*out = *in
	if in.MaxCPU != nil {
		in, out := &in.MaxCPU, &out.MaxCPU
		*out = new(float64)
		**out = **in
	}
	if in.MinCPU != nil {
		in, out := &in.MinCPU, &out.MinCPU
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseInitParameters.
func (in *SQLDatabaseInitParameters) DeepCopy() *SQLDatabaseInitParameters {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseList) DeepCopyInto(out *SQLDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseList.
func (in *SQLDatabaseList) DeepCopy() *SQLDatabaseList {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseObservation) DeepCopyInto(out *SQLDatabaseObservation) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxCPU != nil {
		in, out := &in.MaxCPU, &out.MaxCPU
		*out = new(float64)
		**out = **in
	}
	if in.MinCPU != nil {
		in, out := &in.MinCPU, &out.MinCPU
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseObservation.
func (in *SQLDatabaseObservation) DeepCopy() *SQLDatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseParameters) DeepCopyInto(out *SQLDatabaseParameters) {
	*out = *in
	if in.MaxCPU != nil {
		in, out := &in.MaxCPU, &out.MaxCPU
		*out = new(float64)
		**out = **in
	}
	if in.MinCPU != nil {
		in, out := &in.MinCPU, &out.MinCPU
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseParameters.
func (in *SQLDatabaseParameters) DeepCopy() *SQLDatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseSpec) DeepCopyInto(out *SQLDatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseSpec.
func (in *SQLDatabaseSpec) DeepCopy() *SQLDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseStatus) DeepCopyInto(out *SQLDatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseStatus.
func (in *SQLDatabaseStatus) DeepCopy() *SQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}
