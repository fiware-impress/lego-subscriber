# Subscription

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
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Notification** | Pointer to [**NotificationParams**](NotificationParams.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Subscription) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Subscription) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Subscription) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Subscription) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEntities

`func (o *Subscription) GetEntities() []EntityInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Subscription) GetEntitiesOk() (*[]EntityInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Subscription) SetEntities(v []EntityInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Subscription) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetName

`func (o *Subscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWatchedAttributes

`func (o *Subscription) GetWatchedAttributes() []string`

GetWatchedAttributes returns the WatchedAttributes field if non-nil, zero value otherwise.

### GetWatchedAttributesOk

`func (o *Subscription) GetWatchedAttributesOk() (*[]string, bool)`

GetWatchedAttributesOk returns a tuple with the WatchedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchedAttributes

`func (o *Subscription) SetWatchedAttributes(v []string)`

SetWatchedAttributes sets WatchedAttributes field to given value.

### HasWatchedAttributes

`func (o *Subscription) HasWatchedAttributes() bool`

HasWatchedAttributes returns a boolean if a field has been set.

### GetTimeInterval

`func (o *Subscription) GetTimeInterval() float32`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *Subscription) GetTimeIntervalOk() (*float32, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *Subscription) SetTimeInterval(v float32)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *Subscription) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetExpires

`func (o *Subscription) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Subscription) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Subscription) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Subscription) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsActive

`func (o *Subscription) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Subscription) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Subscription) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Subscription) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetThrottling

`func (o *Subscription) GetThrottling() float32`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *Subscription) GetThrottlingOk() (*float32, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *Subscription) SetThrottling(v float32)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *Subscription) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetQ

`func (o *Subscription) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *Subscription) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *Subscription) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *Subscription) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetGeoQ

`func (o *Subscription) GetGeoQ() GeoQuery`

GetGeoQ returns the GeoQ field if non-nil, zero value otherwise.

### GetGeoQOk

`func (o *Subscription) GetGeoQOk() (*GeoQuery, bool)`

GetGeoQOk returns a tuple with the GeoQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoQ

`func (o *Subscription) SetGeoQ(v GeoQuery)`

SetGeoQ sets GeoQ field to given value.

### HasGeoQ

`func (o *Subscription) HasGeoQ() bool`

HasGeoQ returns a boolean if a field has been set.

### GetCsf

`func (o *Subscription) GetCsf() string`

GetCsf returns the Csf field if non-nil, zero value otherwise.

### GetCsfOk

`func (o *Subscription) GetCsfOk() (*string, bool)`

GetCsfOk returns a tuple with the Csf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsf

`func (o *Subscription) SetCsf(v string)`

SetCsf sets Csf field to given value.

### HasCsf

`func (o *Subscription) HasCsf() bool`

HasCsf returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Subscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Subscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNotification

`func (o *Subscription) GetNotification() NotificationParams`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *Subscription) GetNotificationOk() (*NotificationParams, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *Subscription) SetNotification(v NotificationParams)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *Subscription) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Subscription) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Subscription) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Subscription) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Subscription) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


