/*
ETSI ISG CIM / NGSI-LD API

This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders

API version: latest
Contact: NGSI-LD@etsi.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GeoQuery struct for GeoQuery
type GeoQuery struct {
	Georel string `json:"georel"`
	Coordinates []OneOfintegercoordinates `json:"coordinates"`
	Geometry string `json:"geometry"`
	Geoproperty *string `json:"geoproperty,omitempty"`
}

// NewGeoQuery instantiates a new GeoQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoQuery(georel string, coordinates []OneOfintegercoordinates, geometry string) *GeoQuery {
	this := GeoQuery{}
	this.Georel = georel
	this.Coordinates = coordinates
	this.Geometry = geometry
	return &this
}

// NewGeoQueryWithDefaults instantiates a new GeoQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoQueryWithDefaults() *GeoQuery {
	this := GeoQuery{}
	return &this
}

// GetGeorel returns the Georel field value
func (o *GeoQuery) GetGeorel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Georel
}

// GetGeorelOk returns a tuple with the Georel field value
// and a boolean to check if the value has been set.
func (o *GeoQuery) GetGeorelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Georel, true
}

// SetGeorel sets field value
func (o *GeoQuery) SetGeorel(v string) {
	o.Georel = v
}

// GetCoordinates returns the Coordinates field value
func (o *GeoQuery) GetCoordinates() []OneOfintegercoordinates {
	if o == nil {
		var ret []OneOfintegercoordinates
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *GeoQuery) GetCoordinatesOk() ([]OneOfintegercoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// SetCoordinates sets field value
func (o *GeoQuery) SetCoordinates(v []OneOfintegercoordinates) {
	o.Coordinates = v
}

// GetGeometry returns the Geometry field value
func (o *GeoQuery) GetGeometry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Geometry
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
func (o *GeoQuery) GetGeometryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geometry, true
}

// SetGeometry sets field value
func (o *GeoQuery) SetGeometry(v string) {
	o.Geometry = v
}

// GetGeoproperty returns the Geoproperty field value if set, zero value otherwise.
func (o *GeoQuery) GetGeoproperty() string {
	if o == nil || o.Geoproperty == nil {
		var ret string
		return ret
	}
	return *o.Geoproperty
}

// GetGeopropertyOk returns a tuple with the Geoproperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoQuery) GetGeopropertyOk() (*string, bool) {
	if o == nil || o.Geoproperty == nil {
		return nil, false
	}
	return o.Geoproperty, true
}

// HasGeoproperty returns a boolean if a field has been set.
func (o *GeoQuery) HasGeoproperty() bool {
	if o != nil && o.Geoproperty != nil {
		return true
	}

	return false
}

// SetGeoproperty gets a reference to the given string and assigns it to the Geoproperty field.
func (o *GeoQuery) SetGeoproperty(v string) {
	o.Geoproperty = &v
}

func (o GeoQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["georel"] = o.Georel
	}
	if true {
		toSerialize["coordinates"] = o.Coordinates
	}
	if true {
		toSerialize["geometry"] = o.Geometry
	}
	if o.Geoproperty != nil {
		toSerialize["geoproperty"] = o.Geoproperty
	}
	return json.Marshal(toSerialize)
}

type NullableGeoQuery struct {
	value *GeoQuery
	isSet bool
}

func (v NullableGeoQuery) Get() *GeoQuery {
	return v.value
}

func (v *NullableGeoQuery) Set(val *GeoQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoQuery(val *GeoQuery) *NullableGeoQuery {
	return &NullableGeoQuery{value: val, isSet: true}
}

func (v NullableGeoQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

