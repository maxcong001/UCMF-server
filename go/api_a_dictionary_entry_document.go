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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ADictionaryEntryDocumentApiController binds http requests to an api service and writes the service results to the http response
type ADictionaryEntryDocumentApiController struct {
	service ADictionaryEntryDocumentApiServicer
}

// NewADictionaryEntryDocumentApiController creates a default api controller
func NewADictionaryEntryDocumentApiController(s ADictionaryEntryDocumentApiServicer) Router {
	return &ADictionaryEntryDocumentApiController{ service: s }
}

// Routes returns all of the api route for the ADictionaryEntryDocumentApiController
func (c *ADictionaryEntryDocumentApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateDictionaryEntry",
			strings.ToUpper("Post"),
			"/nucmf-uecm/v1/dic-entries",
			c.CreateDictionaryEntry,
		},
	}
}

// CreateDictionaryEntry - Create a dictionary entry in the UCMF
func (c *ADictionaryEntryDocumentApiController) CreateDictionaryEntry(w http.ResponseWriter, r *http.Request) { 
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	jsonData := r.FormValue("jsonData")
	binaryDataUeRadioCapability5GS, err := ReadFormFileToTempFile(r, "binaryDataUeRadioCapability5GS")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	binaryDataUeRadioCapabilityEPS, err := ReadFormFileToTempFile(r, "binaryDataUeRadioCapabilityEPS")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.CreateDictionaryEntry(r.Context(), jsonData, binaryDataUeRadioCapability5GS, binaryDataUeRadioCapabilityEPS)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}