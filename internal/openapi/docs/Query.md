# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Entities** | Pointer to [**EntityInfo**](EntityInfo.md) |  | [optional] 
**Attrs** | Pointer to **[]string** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**GeoQ** | Pointer to [**GeoQuery**](GeoQuery.md) |  | [optional] 
**TemporalQ** | Pointer to [**TemporalQuery**](TemporalQuery.md) |  | [optional] 
**Csf** | Pointer to **string** |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Query) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Query) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Query) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Query) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEntities

`func (o *Query) GetEntities() EntityInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Query) GetEntitiesOk() (*EntityInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Query) SetEntities(v EntityInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Query) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetAttrs

`func (o *Query) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *Query) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *Query) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *Query) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetQ

`func (o *Query) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *Query) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *Query) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *Query) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetGeoQ

`func (o *Query) GetGeoQ() GeoQuery`

GetGeoQ returns the GeoQ field if non-nil, zero value otherwise.

### GetGeoQOk

`func (o *Query) GetGeoQOk() (*GeoQuery, bool)`

GetGeoQOk returns a tuple with the GeoQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoQ

`func (o *Query) SetGeoQ(v GeoQuery)`

SetGeoQ sets GeoQ field to given value.

### HasGeoQ

`func (o *Query) HasGeoQ() bool`

HasGeoQ returns a boolean if a field has been set.

### GetTemporalQ

`func (o *Query) GetTemporalQ() TemporalQuery`

GetTemporalQ returns the TemporalQ field if non-nil, zero value otherwise.

### GetTemporalQOk

`func (o *Query) GetTemporalQOk() (*TemporalQuery, bool)`

GetTemporalQOk returns a tuple with the TemporalQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalQ

`func (o *Query) SetTemporalQ(v TemporalQuery)`

SetTemporalQ sets TemporalQ field to given value.

### HasTemporalQ

`func (o *Query) HasTemporalQ() bool`

HasTemporalQ returns a boolean if a field has been set.

### GetCsf

`func (o *Query) GetCsf() string`

GetCsf returns the Csf field if non-nil, zero value otherwise.

### GetCsfOk

`func (o *Query) GetCsfOk() (*string, bool)`

GetCsfOk returns a tuple with the Csf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsf

`func (o *Query) SetCsf(v string)`

SetCsf sets Csf field to given value.

### HasCsf

`func (o *Query) HasCsf() bool`

HasCsf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


