# ContextSourceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Information** | [**[]RegistrationInfo**](RegistrationInfo.md) |  | 
**ObservationInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**ManagementInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Location** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**ObservationSpace** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**OperationSpace** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Endpoint** | **string** |  | 
**Id** | **string** |  | 
**Type** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContextSourceRegistration

`func NewContextSourceRegistration(information []RegistrationInfo, endpoint string, id string, type_ string, ) *ContextSourceRegistration`

NewContextSourceRegistration instantiates a new ContextSourceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextSourceRegistrationWithDefaults

`func NewContextSourceRegistrationWithDefaults() *ContextSourceRegistration`

NewContextSourceRegistrationWithDefaults instantiates a new ContextSourceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ContextSourceRegistration) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContextSourceRegistration) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContextSourceRegistration) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ContextSourceRegistration) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetInformation

`func (o *ContextSourceRegistration) GetInformation() []RegistrationInfo`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *ContextSourceRegistration) GetInformationOk() (*[]RegistrationInfo, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *ContextSourceRegistration) SetInformation(v []RegistrationInfo)`

SetInformation sets Information field to given value.


### GetObservationInterval

`func (o *ContextSourceRegistration) GetObservationInterval() TimeInterval`

GetObservationInterval returns the ObservationInterval field if non-nil, zero value otherwise.

### GetObservationIntervalOk

`func (o *ContextSourceRegistration) GetObservationIntervalOk() (*TimeInterval, bool)`

GetObservationIntervalOk returns a tuple with the ObservationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationInterval

`func (o *ContextSourceRegistration) SetObservationInterval(v TimeInterval)`

SetObservationInterval sets ObservationInterval field to given value.

### HasObservationInterval

`func (o *ContextSourceRegistration) HasObservationInterval() bool`

HasObservationInterval returns a boolean if a field has been set.

### GetManagementInterval

`func (o *ContextSourceRegistration) GetManagementInterval() TimeInterval`

GetManagementInterval returns the ManagementInterval field if non-nil, zero value otherwise.

### GetManagementIntervalOk

`func (o *ContextSourceRegistration) GetManagementIntervalOk() (*TimeInterval, bool)`

GetManagementIntervalOk returns a tuple with the ManagementInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementInterval

`func (o *ContextSourceRegistration) SetManagementInterval(v TimeInterval)`

SetManagementInterval sets ManagementInterval field to given value.

### HasManagementInterval

`func (o *ContextSourceRegistration) HasManagementInterval() bool`

HasManagementInterval returns a boolean if a field has been set.

### GetLocation

`func (o *ContextSourceRegistration) GetLocation() Geometry`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContextSourceRegistration) GetLocationOk() (*Geometry, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContextSourceRegistration) SetLocation(v Geometry)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContextSourceRegistration) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *ContextSourceRegistration) GetObservationSpace() Geometry`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *ContextSourceRegistration) GetObservationSpaceOk() (*Geometry, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *ContextSourceRegistration) SetObservationSpace(v Geometry)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *ContextSourceRegistration) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *ContextSourceRegistration) GetOperationSpace() Geometry`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *ContextSourceRegistration) GetOperationSpaceOk() (*Geometry, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *ContextSourceRegistration) SetOperationSpace(v Geometry)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *ContextSourceRegistration) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetExpires

`func (o *ContextSourceRegistration) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ContextSourceRegistration) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ContextSourceRegistration) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ContextSourceRegistration) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *ContextSourceRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextSourceRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextSourceRegistration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContextSourceRegistration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ContextSourceRegistration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextSourceRegistration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextSourceRegistration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContextSourceRegistration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *ContextSourceRegistration) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ContextSourceRegistration) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ContextSourceRegistration) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetId

`func (o *ContextSourceRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextSourceRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextSourceRegistration) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ContextSourceRegistration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContextSourceRegistration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContextSourceRegistration) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *ContextSourceRegistration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContextSourceRegistration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContextSourceRegistration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContextSourceRegistration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ContextSourceRegistration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ContextSourceRegistration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ContextSourceRegistration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ContextSourceRegistration) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


