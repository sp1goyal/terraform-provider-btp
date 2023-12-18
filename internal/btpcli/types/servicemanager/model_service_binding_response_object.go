/*
 * Service Manager
 *
 * Service Manager provides REST APIs that are responsible for the creation and consumption of service instances in any connected runtime environment.   Use the Service Manager APIs to perform various operations related to your platforms, service brokers, service instances, and service bindings.  Get service plans and service offerings associated with your environment.    #### Platforms   Platforms are OSBAPI-enabled software systems on which applications and services are hosted.   With the Service Manager, you can now register your platform and enable it to consume the SAP BTP services from your native environment.   This registration results in a returned set of credentials that are needed to deploy the Service Manager agent.     #### Service Brokers   Service brokers act as brokers between the Service Manager and a platform’s marketplace to advertise catalogues of service offerings and service plans.  They also receive and process the requests from the marketplace to provision, bind, unbind, and deprovision these offerings and plans.    #### Service Instances   Service instances are instantiations of service plans that make the functionality of those service plans available for consumption.    #### Service Bindings   Service bindings provide access details to existing service instances.  The access details are part of the service bindings' ‘credentials’ property, and typically include access URLs and credentials.    #### Service Plans   Service plans represent sets of capabilities provided by a service offering.  For example, database service offerings provide different plans for different database versions or sizes, while the Service Manager plans offer different data access levels.    #### Service Offerings   Service offerings are advertisements of the services that are supported by a service broker.  For example, software that you can consume in the subaccount.  Service offerings are related to one or more service plans.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package servicemanager

import (
	"encoding/json"
	"time"
)

type ServiceBindingResponseObject struct {
	// The ID of the service binding.
	Id string `json:"id,omitempty"`
	// Whether the service binding is ready.
	Ready         bool                     `json:"ready,omitempty"`
	LastOperation *OperationResponseObject `json:"last_operation,omitempty"`
	// The name of the service binding.
	Name string `json:"name,omitempty"`
	// The ID of the service instance associated with the binding.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	// The ID of the subaccount.
	SubaccountId string `json:"subaccount_id,omitempty"`
	// Contextual data for the resource.
	Context json.RawMessage `json:"context,omitempty"`
	// Credentials to access the binding.
	Credentials json.RawMessage `json:"credentials,omitempty"`
	// The bound resources.
	BindResource map[string]string `json:"bind_resource,omitempty"`
	// The time the binding was created.<br/>In ISO 8601 format:</br> YYYY-MM-DDThh:mm:ssTZD
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The last time the binding was updated.<br/> In ISO 8601 format.
	UpdatedAt time.Time            `json:"updated_at,omitempty"`
	Labels    ServiceManagerLabels `json:"labels,omitempty"`
}
