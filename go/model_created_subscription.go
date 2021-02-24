/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type CreatedSubscription struct {

	DicEntryId int32 `json:"dicEntryId" bson:"dicEntryId"`

	ConfirmedExpires time.Time `json:"confirmedExpires,omitempty" bson:"confirmedExpires"`
}
