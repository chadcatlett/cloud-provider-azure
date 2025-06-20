/*
Copyright 2021 The Kubernetes Authors.

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

package options

import (
	"github.com/spf13/pflag"

	app "sigs.k8s.io/cloud-provider-azure/cmd/cloud-controller-manager/app/config"
)

// LoadBalancerOptions holds the configurations of the service-lb-controller
type LoadBalancerOptions struct {
	LoadBalancerClass string
}

// AddFlags adds flags related to dynamic reloading for controller manager to the specified FlagSet
func (o *LoadBalancerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.StringVar(&o.LoadBalancerClass, "loadbalancerclass", "", "The loadbalancerclass name to use for filtering services to process")
}

// ApplyTo fills up dynamic reloading config with options
func (o *LoadBalancerOptions) ApplyTo(cfg *app.LoadBalancerConfig) error {
	if o == nil {
		return nil
	}

	cfg.LoadBalancerClass = o.LoadBalancerClass

	return nil
}

// Validate checks validation of LoadBalancerOptions
func (o *LoadBalancerOptions) Validate() []error {
	return nil
}

func defaultLoadBalancerOptions() *LoadBalancerOptions {
	return &LoadBalancerOptions{
		LoadBalancerClass: "",
	}
}
