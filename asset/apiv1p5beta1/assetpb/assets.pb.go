// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/cloud/asset/v1p5beta1/assets.proto

package assetpb

import (
	reflect "reflect"
	sync "sync"

	v1 "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	orgpolicypb "cloud.google.com/go/orgpolicy/apiv1/orgpolicypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An asset in Google Cloud. An asset can be any resource in the Google Cloud
// [resource
// hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy),
// a resource outside the Google Cloud resource hierarchy (such as Google
// Kubernetes Engine clusters and objects), or a policy (e.g. IAM policy).
// See [Supported asset
// types](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
// for more information.
type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full name of the asset. Example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`
	//
	// See [Resource
	// names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the asset. Example: `compute.googleapis.com/Disk`
	//
	// See [Supported asset
	// types](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	// for more information.
	AssetType string `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// A representation of the resource.
	Resource *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// A representation of the IAM policy set on a Google Cloud resource.
	// There can be a maximum of one IAM policy set on any given resource.
	// In addition, IAM policies inherit their granted access scope from any
	// policies set on parent resources in the resource hierarchy. Therefore, the
	// effectively policy is the union of both the policy set on this resource
	// and each policy set on all of the resource's ancestry resource levels in
	// the hierarchy. See
	// [this topic](https://cloud.google.com/iam/help/allow-policies/inheritance)
	// for more information.
	IamPolicy *iampb.Policy `protobuf:"bytes,4,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	// A representation of an [organization
	// policy](https://cloud.google.com/resource-manager/docs/organization-policy/overview#organization_policy).
	// There can be more than one organization policy with different constraints
	// set on a given resource.
	OrgPolicy []*orgpolicypb.Policy `protobuf:"bytes,6,rep,name=org_policy,json=orgPolicy,proto3" json:"org_policy,omitempty"`
	// A representation of an [access
	// policy](https://cloud.google.com/access-context-manager/docs/overview#access-policies).
	//
	// Types that are assignable to AccessContextPolicy:
	//
	//	*Asset_AccessPolicy
	//	*Asset_AccessLevel
	//	*Asset_ServicePerimeter
	AccessContextPolicy isAsset_AccessContextPolicy `protobuf_oneof:"access_context_policy"`
	// The ancestry path of an asset in Google Cloud [resource
	// hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy),
	// represented as a list of relative resource names. An ancestry path starts
	// with the closest ancestor in the hierarchy and ends at root. If the asset
	// is a project, folder, or organization, the ancestry path starts from the
	// asset itself.
	//
	// Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	Ancestors []string `protobuf:"bytes,10,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p5beta1_assets_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *Asset) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Asset) GetIamPolicy() *iampb.Policy {
	if x != nil {
		return x.IamPolicy
	}
	return nil
}

func (x *Asset) GetOrgPolicy() []*orgpolicypb.Policy {
	if x != nil {
		return x.OrgPolicy
	}
	return nil
}

func (m *Asset) GetAccessContextPolicy() isAsset_AccessContextPolicy {
	if m != nil {
		return m.AccessContextPolicy
	}
	return nil
}

func (x *Asset) GetAccessPolicy() *v1.AccessPolicy {
	if x, ok := x.GetAccessContextPolicy().(*Asset_AccessPolicy); ok {
		return x.AccessPolicy
	}
	return nil
}

func (x *Asset) GetAccessLevel() *v1.AccessLevel {
	if x, ok := x.GetAccessContextPolicy().(*Asset_AccessLevel); ok {
		return x.AccessLevel
	}
	return nil
}

func (x *Asset) GetServicePerimeter() *v1.ServicePerimeter {
	if x, ok := x.GetAccessContextPolicy().(*Asset_ServicePerimeter); ok {
		return x.ServicePerimeter
	}
	return nil
}

func (x *Asset) GetAncestors() []string {
	if x != nil {
		return x.Ancestors
	}
	return nil
}

type isAsset_AccessContextPolicy interface {
	isAsset_AccessContextPolicy()
}

type Asset_AccessPolicy struct {
	// Please also refer to the [access policy user
	// guide](https://cloud.google.com/access-context-manager/docs/overview#access-policies).
	AccessPolicy *v1.AccessPolicy `protobuf:"bytes,7,opt,name=access_policy,json=accessPolicy,proto3,oneof"`
}

type Asset_AccessLevel struct {
	// Please also refer to the [access level user
	// guide](https://cloud.google.com/access-context-manager/docs/overview#access-levels).
	AccessLevel *v1.AccessLevel `protobuf:"bytes,8,opt,name=access_level,json=accessLevel,proto3,oneof"`
}

type Asset_ServicePerimeter struct {
	// Please also refer to the [service perimeter user
	// guide](https://cloud.google.com/vpc-service-controls/docs/overview).
	ServicePerimeter *v1.ServicePerimeter `protobuf:"bytes,9,opt,name=service_perimeter,json=servicePerimeter,proto3,oneof"`
}

func (*Asset_AccessPolicy) isAsset_AccessContextPolicy() {}

func (*Asset_AccessLevel) isAsset_AccessContextPolicy() {}

func (*Asset_ServicePerimeter) isAsset_AccessContextPolicy() {}

// A representation of a Google Cloud resource.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version. Example: "v1".
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The URL of the discovery document containing the resource's JSON schema.
	// Example:
	// `https://www.googleapis.com/discovery/v1/apis/compute/v1/rest`
	//
	// This value is unspecified for resources that do not have an API based on a
	// discovery document, such as Cloud Bigtable.
	DiscoveryDocumentUri string `protobuf:"bytes,2,opt,name=discovery_document_uri,json=discoveryDocumentUri,proto3" json:"discovery_document_uri,omitempty"`
	// The JSON schema name listed in the discovery document. Example:
	// `Project`
	//
	// This value is unspecified for resources that do not have an API based on a
	// discovery document, such as Cloud Bigtable.
	DiscoveryName string `protobuf:"bytes,3,opt,name=discovery_name,json=discoveryName,proto3" json:"discovery_name,omitempty"`
	// The REST URL for accessing the resource. An HTTP `GET` request using this
	// URL returns the resource itself. Example:
	// `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123`
	//
	// This value is unspecified for resources without a REST API.
	ResourceUrl string `protobuf:"bytes,4,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The full name of the immediate parent of this resource. See
	// [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	//
	// For Google Cloud assets, this value is the parent resource defined in the
	// [IAM policy
	// hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy).
	// Example:
	// `//cloudresourcemanager.googleapis.com/projects/my_project_123`
	//
	// For third-party assets, this field may be set differently.
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The content of the resource, in which some sensitive fields are removed
	// and may not be present.
	Data *structpb.Struct `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p5beta1_assets_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Resource) GetDiscoveryDocumentUri() string {
	if x != nil {
		return x.DiscoveryDocumentUri
	}
	return ""
}

func (x *Resource) GetDiscoveryName() string {
	if x != nil {
		return x.DiscoveryName
	}
	return ""
}

func (x *Resource) GetResourceUrl() string {
	if x != nil {
		return x.ResourceUrl
	}
	return ""
}

func (x *Resource) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Resource) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_google_cloud_asset_v1p5beta1_assets_proto protoreflect.FileDescriptor

var file_google_cloud_asset_v1p5beta1_assets_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x04, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x09, 0x69, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x6f, 0x72,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72,
	0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x09, 0x6f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5c, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x59, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x68, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x27, 0xea,
	0x41, 0x24, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x01, 0x2a, 0x42, 0x17, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22,
	0xe9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xa9, 0x01, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x70,
	0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x70, 0x62, 0x3b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x56,
	0x31, 0x70, 0x35, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x5c, 0x56, 0x31,
	0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_asset_v1p5beta1_assets_proto_rawDescOnce sync.Once
	file_google_cloud_asset_v1p5beta1_assets_proto_rawDescData = file_google_cloud_asset_v1p5beta1_assets_proto_rawDesc
)

func file_google_cloud_asset_v1p5beta1_assets_proto_rawDescGZIP() []byte {
	file_google_cloud_asset_v1p5beta1_assets_proto_rawDescOnce.Do(func() {
		file_google_cloud_asset_v1p5beta1_assets_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_asset_v1p5beta1_assets_proto_rawDescData)
	})
	return file_google_cloud_asset_v1p5beta1_assets_proto_rawDescData
}

var file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_asset_v1p5beta1_assets_proto_goTypes = []interface{}{
	(*Asset)(nil),               // 0: google.cloud.asset.v1p5beta1.Asset
	(*Resource)(nil),            // 1: google.cloud.asset.v1p5beta1.Resource
	(*iampb.Policy)(nil),        // 2: google.iam.v1.Policy
	(*orgpolicypb.Policy)(nil),  // 3: google.cloud.orgpolicy.v1.Policy
	(*v1.AccessPolicy)(nil),     // 4: google.identity.accesscontextmanager.v1.AccessPolicy
	(*v1.AccessLevel)(nil),      // 5: google.identity.accesscontextmanager.v1.AccessLevel
	(*v1.ServicePerimeter)(nil), // 6: google.identity.accesscontextmanager.v1.ServicePerimeter
	(*structpb.Struct)(nil),     // 7: google.protobuf.Struct
}
var file_google_cloud_asset_v1p5beta1_assets_proto_depIdxs = []int32{
	1, // 0: google.cloud.asset.v1p5beta1.Asset.resource:type_name -> google.cloud.asset.v1p5beta1.Resource
	2, // 1: google.cloud.asset.v1p5beta1.Asset.iam_policy:type_name -> google.iam.v1.Policy
	3, // 2: google.cloud.asset.v1p5beta1.Asset.org_policy:type_name -> google.cloud.orgpolicy.v1.Policy
	4, // 3: google.cloud.asset.v1p5beta1.Asset.access_policy:type_name -> google.identity.accesscontextmanager.v1.AccessPolicy
	5, // 4: google.cloud.asset.v1p5beta1.Asset.access_level:type_name -> google.identity.accesscontextmanager.v1.AccessLevel
	6, // 5: google.cloud.asset.v1p5beta1.Asset.service_perimeter:type_name -> google.identity.accesscontextmanager.v1.ServicePerimeter
	7, // 6: google.cloud.asset.v1p5beta1.Resource.data:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_asset_v1p5beta1_assets_proto_init() }
func file_google_cloud_asset_v1p5beta1_assets_proto_init() {
	if File_google_cloud_asset_v1p5beta1_assets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Asset_AccessPolicy)(nil),
		(*Asset_AccessLevel)(nil),
		(*Asset_ServicePerimeter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_asset_v1p5beta1_assets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_asset_v1p5beta1_assets_proto_goTypes,
		DependencyIndexes: file_google_cloud_asset_v1p5beta1_assets_proto_depIdxs,
		MessageInfos:      file_google_cloud_asset_v1p5beta1_assets_proto_msgTypes,
	}.Build()
	File_google_cloud_asset_v1p5beta1_assets_proto = out.File
	file_google_cloud_asset_v1p5beta1_assets_proto_rawDesc = nil
	file_google_cloud_asset_v1p5beta1_assets_proto_goTypes = nil
	file_google_cloud_asset_v1p5beta1_assets_proto_depIdxs = nil
}
