// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Tag A tag definition that belongs to a specific tag namespace.  "Defined tags" must be set up in your tenancy before
// you can apply them to resources.
// For more information, see Managing Tags and Tag Namespaces (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm).
type Tag struct {

	// The OCID of the compartment that contains the tag definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the namespace that contains the tag definition.
	TagNamespaceId *string `mandatory:"true" json:"tagNamespaceId"`

	// The name of the tag namespace that contains the tag definition.
	TagNamespaceName *string `mandatory:"true" json:"tagNamespaceName"`

	// The OCID of the tag definition.
	Id *string `mandatory:"true" json:"id"`

	// The name of the tag. The name must be unique across all tags in the namespace and can't be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the tag.
	Description *string `mandatory:"true" json:"description"`

	// Indicates whether the tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}``
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`
}

func (m Tag) String() string {
	return common.PointerString(m)
}
