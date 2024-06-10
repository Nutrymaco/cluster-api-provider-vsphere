/*
 Copyright 2023 The Kubernetes Authors.
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

package v1alpha3

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	v1beta1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
)

func Convert_v1beta1_VSphereFailureDomainSpec_To_v1alpha3_VSphereFailureDomainSpec(in *v1beta1.VSphereFailureDomainSpec, out *VSphereFailureDomainSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_FailureDomain_To_v1alpha3_FailureDomain(&in.Region, &out.Region, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_FailureDomain_To_v1alpha3_FailureDomain(&in.Zone, &out.Zone, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_Topology_To_v1alpha3_Topology(&in.Topology, &out.Topology, s); err != nil {
		return err
	}
	return nil
}