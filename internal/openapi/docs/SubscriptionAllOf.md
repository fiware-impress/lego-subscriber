# SubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Notification** | Pointer to [**NotificationParams**](NotificationParams.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriptionAllOf

`func NewSubscriptionAllOf() *SubscriptionAllOf`

NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAllOfWithDefaults

`func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf`

NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNotification

`func (o *SubscriptionAllOf) GetNotification() NotificationParams`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *SubscriptionAllOf) GetNotificationOk() (*NotificationParams, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *SubscriptionAllOf) SetNotification(v NotificationParams)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *SubscriptionAllOf) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SubscriptionAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SubscriptionAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SubscriptionAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SubscriptionAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


