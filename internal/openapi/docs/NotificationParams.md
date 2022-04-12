# NotificationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **[]string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Endpoint** | [**Endpoint**](Endpoint.md) |  | 
**Status** | Pointer to **string** |  | [optional] 
**TimesSent** | Pointer to **float32** |  | [optional] 
**LastNotification** | Pointer to **time.Time** |  | [optional] 
**LastFailure** | Pointer to **time.Time** |  | [optional] 
**LastSuccess** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNotificationParams

`func NewNotificationParams(endpoint Endpoint, ) *NotificationParams`

NewNotificationParams instantiates a new NotificationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationParamsWithDefaults

`func NewNotificationParamsWithDefaults() *NotificationParams`

NewNotificationParamsWithDefaults instantiates a new NotificationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NotificationParams) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotificationParams) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotificationParams) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NotificationParams) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetFormat

`func (o *NotificationParams) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *NotificationParams) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *NotificationParams) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *NotificationParams) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetEndpoint

`func (o *NotificationParams) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NotificationParams) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NotificationParams) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.


### GetStatus

`func (o *NotificationParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimesSent

`func (o *NotificationParams) GetTimesSent() float32`

GetTimesSent returns the TimesSent field if non-nil, zero value otherwise.

### GetTimesSentOk

`func (o *NotificationParams) GetTimesSentOk() (*float32, bool)`

GetTimesSentOk returns a tuple with the TimesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesSent

`func (o *NotificationParams) SetTimesSent(v float32)`

SetTimesSent sets TimesSent field to given value.

### HasTimesSent

`func (o *NotificationParams) HasTimesSent() bool`

HasTimesSent returns a boolean if a field has been set.

### GetLastNotification

`func (o *NotificationParams) GetLastNotification() time.Time`

GetLastNotification returns the LastNotification field if non-nil, zero value otherwise.

### GetLastNotificationOk

`func (o *NotificationParams) GetLastNotificationOk() (*time.Time, bool)`

GetLastNotificationOk returns a tuple with the LastNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotification

`func (o *NotificationParams) SetLastNotification(v time.Time)`

SetLastNotification sets LastNotification field to given value.

### HasLastNotification

`func (o *NotificationParams) HasLastNotification() bool`

HasLastNotification returns a boolean if a field has been set.

### GetLastFailure

`func (o *NotificationParams) GetLastFailure() time.Time`

GetLastFailure returns the LastFailure field if non-nil, zero value otherwise.

### GetLastFailureOk

`func (o *NotificationParams) GetLastFailureOk() (*time.Time, bool)`

GetLastFailureOk returns a tuple with the LastFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailure

`func (o *NotificationParams) SetLastFailure(v time.Time)`

SetLastFailure sets LastFailure field to given value.

### HasLastFailure

`func (o *NotificationParams) HasLastFailure() bool`

HasLastFailure returns a boolean if a field has been set.

### GetLastSuccess

`func (o *NotificationParams) GetLastSuccess() time.Time`

GetLastSuccess returns the LastSuccess field if non-nil, zero value otherwise.

### GetLastSuccessOk

`func (o *NotificationParams) GetLastSuccessOk() (*time.Time, bool)`

GetLastSuccessOk returns a tuple with the LastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccess

`func (o *NotificationParams) SetLastSuccess(v time.Time)`

SetLastSuccess sets LastSuccess field to given value.

### HasLastSuccess

`func (o *NotificationParams) HasLastSuccess() bool`

HasLastSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


