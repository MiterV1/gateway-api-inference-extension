/*
Copyright 2024.

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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	apiv1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/api/v1alpha1"
	applyconfigurationapiv1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/client-go/applyconfiguration/api/v1alpha1"
	scheme "inference.networking.x-k8s.io/llm-instance-gateway/client-go/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// LLMServicesGetter has a method to return a LLMServiceInterface.
// A group's client should implement this interface.
type LLMServicesGetter interface {
	LLMServices(namespace string) LLMServiceInterface
}

// LLMServiceInterface has methods to work with LLMService resources.
type LLMServiceInterface interface {
	Create(ctx context.Context, lLMService *apiv1alpha1.LLMService, opts v1.CreateOptions) (*apiv1alpha1.LLMService, error)
	Update(ctx context.Context, lLMService *apiv1alpha1.LLMService, opts v1.UpdateOptions) (*apiv1alpha1.LLMService, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, lLMService *apiv1alpha1.LLMService, opts v1.UpdateOptions) (*apiv1alpha1.LLMService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apiv1alpha1.LLMService, error)
	List(ctx context.Context, opts v1.ListOptions) (*apiv1alpha1.LLMServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1alpha1.LLMService, err error)
	Apply(ctx context.Context, lLMService *applyconfigurationapiv1alpha1.LLMServiceApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.LLMService, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, lLMService *applyconfigurationapiv1alpha1.LLMServiceApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.LLMService, err error)
	LLMServiceExpansion
}

// lLMServices implements LLMServiceInterface
type lLMServices struct {
	*gentype.ClientWithListAndApply[*apiv1alpha1.LLMService, *apiv1alpha1.LLMServiceList, *applyconfigurationapiv1alpha1.LLMServiceApplyConfiguration]
}

// newLLMServices returns a LLMServices
func newLLMServices(c *ApiV1alpha1Client, namespace string) *lLMServices {
	return &lLMServices{
		gentype.NewClientWithListAndApply[*apiv1alpha1.LLMService, *apiv1alpha1.LLMServiceList, *applyconfigurationapiv1alpha1.LLMServiceApplyConfiguration](
			"llmservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apiv1alpha1.LLMService { return &apiv1alpha1.LLMService{} },
			func() *apiv1alpha1.LLMServiceList { return &apiv1alpha1.LLMServiceList{} }),
	}
}
