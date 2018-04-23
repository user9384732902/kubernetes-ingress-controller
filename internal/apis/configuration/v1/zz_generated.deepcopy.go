// +build !ignore_autogenerated

/*
Copyright 2018 The Kong Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheck) DeepCopyInto(out *ActiveHealthCheck) {
	*out = *in
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		if *in == nil {
			*out = nil
		} else {
			*out = new(Healthy)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		if *in == nil {
			*out = nil
		} else {
			*out = new(Unhealthy)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheck.
func (in *ActiveHealthCheck) DeepCopy() *ActiveHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Healthchecks) DeepCopyInto(out *Healthchecks) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		if *in == nil {
			*out = nil
		} else {
			*out = new(ActiveHealthCheck)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Passive != nil {
		in, out := &in.Passive, &out.Passive
		if *in == nil {
			*out = nil
		} else {
			*out = new(Passive)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Healthchecks.
func (in *Healthchecks) DeepCopy() *Healthchecks {
	if in == nil {
		return nil
	}
	out := new(Healthchecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Healthy) DeepCopyInto(out *Healthy) {
	*out = *in
	if in.HTTPStatuses != nil {
		in, out := &in.HTTPStatuses, &out.HTTPStatuses
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Healthy.
func (in *Healthy) DeepCopy() *Healthy {
	if in == nil {
		return nil
	}
	out := new(Healthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngress) DeepCopyInto(out *KongIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		if *in == nil {
			*out = nil
		} else {
			*out = new(Upstream)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		if *in == nil {
			*out = nil
		} else {
			*out = new(Proxy)
			**out = **in
		}
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		if *in == nil {
			*out = nil
		} else {
			*out = new(Route)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngress.
func (in *KongIngress) DeepCopy() *KongIngress {
	if in == nil {
		return nil
	}
	out := new(KongIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressList) DeepCopyInto(out *KongIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressList.
func (in *KongIngressList) DeepCopy() *KongIngressList {
	if in == nil {
		return nil
	}
	out := new(KongIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Passive) DeepCopyInto(out *Passive) {
	*out = *in
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		if *in == nil {
			*out = nil
		} else {
			*out = new(Healthy)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		if *in == nil {
			*out = nil
		} else {
			*out = new(Unhealthy)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Passive.
func (in *Passive) DeepCopy() *Passive {
	if in == nil {
		return nil
	}
	out := new(Passive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Unhealthy) DeepCopyInto(out *Unhealthy) {
	*out = *in
	if in.HTTPStatuses != nil {
		in, out := &in.HTTPStatuses, &out.HTTPStatuses
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Unhealthy.
func (in *Unhealthy) DeepCopy() *Unhealthy {
	if in == nil {
		return nil
	}
	out := new(Unhealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upstream) DeepCopyInto(out *Upstream) {
	*out = *in
	if in.Healthchecks != nil {
		in, out := &in.Healthchecks, &out.Healthchecks
		if *in == nil {
			*out = nil
		} else {
			*out = new(Healthchecks)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upstream.
func (in *Upstream) DeepCopy() *Upstream {
	if in == nil {
		return nil
	}
	out := new(Upstream)
	in.DeepCopyInto(out)
	return out
}
