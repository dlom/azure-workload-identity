//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscriptions

const (
	moduleName    = "armsubscriptions"
	moduleVersion = "v1.1.1"
)

// LocationType - The location type.
type LocationType string

const (
	LocationTypeRegion   LocationType = "Region"
	LocationTypeEdgeZone LocationType = "EdgeZone"
)

// PossibleLocationTypeValues returns the possible values for the LocationType const type.
func PossibleLocationTypeValues() []LocationType {
	return []LocationType{
		LocationTypeRegion,
		LocationTypeEdgeZone,
	}
}

// RegionCategory - The category of the region.
type RegionCategory string

const (
	RegionCategoryExtended    RegionCategory = "Extended"
	RegionCategoryOther       RegionCategory = "Other"
	RegionCategoryRecommended RegionCategory = "Recommended"
)

// PossibleRegionCategoryValues returns the possible values for the RegionCategory const type.
func PossibleRegionCategoryValues() []RegionCategory {
	return []RegionCategory{
		RegionCategoryExtended,
		RegionCategoryOther,
		RegionCategoryRecommended,
	}
}

// RegionType - The type of the region.
type RegionType string

const (
	RegionTypeLogical  RegionType = "Logical"
	RegionTypePhysical RegionType = "Physical"
)

// PossibleRegionTypeValues returns the possible values for the RegionType const type.
func PossibleRegionTypeValues() []RegionType {
	return []RegionType{
		RegionTypeLogical,
		RegionTypePhysical,
	}
}

// ResourceNameStatus - Is the resource name Allowed or Reserved
type ResourceNameStatus string

const (
	ResourceNameStatusAllowed  ResourceNameStatus = "Allowed"
	ResourceNameStatusReserved ResourceNameStatus = "Reserved"
)

// PossibleResourceNameStatusValues returns the possible values for the ResourceNameStatus const type.
func PossibleResourceNameStatusValues() []ResourceNameStatus {
	return []ResourceNameStatus{
		ResourceNameStatusAllowed,
		ResourceNameStatusReserved,
	}
}

// SpendingLimit - The subscription spending limit.
type SpendingLimit string

const (
	SpendingLimitOn               SpendingLimit = "On"
	SpendingLimitOff              SpendingLimit = "Off"
	SpendingLimitCurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
)

// PossibleSpendingLimitValues returns the possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{
		SpendingLimitOn,
		SpendingLimitOff,
		SpendingLimitCurrentPeriodOff,
	}
}

// SubscriptionState - The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
type SubscriptionState string

const (
	SubscriptionStateEnabled  SubscriptionState = "Enabled"
	SubscriptionStateWarned   SubscriptionState = "Warned"
	SubscriptionStatePastDue  SubscriptionState = "PastDue"
	SubscriptionStateDisabled SubscriptionState = "Disabled"
	SubscriptionStateDeleted  SubscriptionState = "Deleted"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateEnabled,
		SubscriptionStateWarned,
		SubscriptionStatePastDue,
		SubscriptionStateDisabled,
		SubscriptionStateDeleted,
	}
}

// TenantCategory - Category of the tenant.
type TenantCategory string

const (
	TenantCategoryHome        TenantCategory = "Home"
	TenantCategoryProjectedBy TenantCategory = "ProjectedBy"
	TenantCategoryManagedBy   TenantCategory = "ManagedBy"
)

// PossibleTenantCategoryValues returns the possible values for the TenantCategory const type.
func PossibleTenantCategoryValues() []TenantCategory {
	return []TenantCategory{
		TenantCategoryHome,
		TenantCategoryProjectedBy,
		TenantCategoryManagedBy,
	}
}