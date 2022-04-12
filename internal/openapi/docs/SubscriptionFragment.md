# SubscriptionFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Entities** | Pointer to [**[]EntityInfo**](EntityInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WatchedAttributes** | Pointer to **[]string** |  | [optional] 
**TimeInterval** | Pointer to **float32** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Throttling** | Pointer to **float32** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**GeoQ** | Pointer to [**GeoQuery**](GeoQuery.md) |  | [optional] 
**Csf** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionFragment

`func NewSubscriptionFragment() *SubscriptionFragment`

NewSubscriptionFragment instantiates a new SubscriptionFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionFragmentWithDefaults

`func NewSubscriptionFragmentWithDefaults() *SubscriptionFragment`

NewSubscriptionFragmentWithDefaults instantiates a new SubscriptionFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SubscriptionFragment) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SubscriptionFragment) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SubscriptionFragment) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *SubscriptionFragment) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEntities

`func (o *SubscriptionFragment) GetEntities() []EntityInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SubscriptionFragment) GetEntitiesOk() (*[]EntityInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SubscriptionFragment) SetEntities(v []EntityInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SubscriptionFragment) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionFragment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionFragment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionFragment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionFragment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SubscriptionFragment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionFragment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionFragment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionFragment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWatchedAttributes

`func (o *SubscriptionFragment) GetWatchedAttributes() []string`

GetWatchedAttributes returns the WatchedAttributes field if non-nil, zero value otherwise.

### GetWatchedAttributesOk

`func (o *SubscriptionFragment) GetWatchedAttributesOk() (*[]string, bool)`

GetWatchedAttributesOk returns a tuple with the WatchedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchedAttributes

`func (o *SubscriptionFragment) SetWatchedAttributes(v []string)`

SetWatchedAttributes sets WatchedAttributes field to given value.

### HasWatchedAttributes

`func (o *SubscriptionFragment) HasWatchedAttributes() bool`

HasWatchedAttributes returns a boolean if a field has been set.

### GetTimeInterval

`func (o *SubscriptionFragment) GetTimeInterval() float32`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *SubscriptionFragment) GetTimeIntervalOk() (*float32, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *SubscriptionFragment) SetTimeInterval(v float32)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *SubscriptionFragment) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetExpires

`func (o *SubscriptionFragment) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SubscriptionFragment) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SubscriptionFragment) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SubscriptionFragment) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsActive

`func (o *SubscriptionFragment) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SubscriptionFragment) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SubscriptionFragment) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SubscriptionFragment) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetThrottling

`func (o *SubscriptionFragment) GetThrottling() float32`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *SubscriptionFragment) GetThrottlingOk() (*float32, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *SubscriptionFragment) SetThrottling(v float32)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *SubscriptionFragment) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetQ

`func (o *SubscriptionFragment) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *SubscriptionFragment) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *SubscriptionFragment) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *SubscriptionFragment) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetGeoQ

`func (o *SubscriptionFragment) GetGeoQ() GeoQuery`

GetGeoQ returns the GeoQ field if non-nil, zero value otherwise.

### GetGeoQOk

`func (o *SubscriptionFragment) GetGeoQOk() (*GeoQuery, bool)`

GetGeoQOk returns a tuple with the GeoQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoQ

`func (o *SubscriptionFragment) SetGeoQ(v GeoQuery)`

SetGeoQ sets GeoQ field to given value.

### HasGeoQ

`func (o *SubscriptionFragment) HasGeoQ() bool`

HasGeoQ returns a boolean if a field has been set.

### GetCsf

`func (o *SubscriptionFragment) GetCsf() string`

GetCsf returns the Csf field if non-nil, zero value otherwise.

### GetCsfOk

`func (o *SubscriptionFragment) GetCsfOk() (*string, bool)`

GetCsfOk returns a tuple with the Csf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsf

`func (o *SubscriptionFragment) SetCsf(v string)`

SetCsf sets Csf field to given value.

### HasCsf

`func (o *SubscriptionFragment) HasCsf() bool`

HasCsf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


