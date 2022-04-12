# ContextSourceRegistrationFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Information** | Pointer to [**[]RegistrationInfo**](RegistrationInfo.md) |  | [optional] 
**ObservationInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**ManagementInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Location** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**ObservationSpace** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**OperationSpace** | Pointer to [**Geometry**](Geometry.md) |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 

## Methods

### NewContextSourceRegistrationFragment

`func NewContextSourceRegistrationFragment() *ContextSourceRegistrationFragment`

NewContextSourceRegistrationFragment instantiates a new ContextSourceRegistrationFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextSourceRegistrationFragmentWithDefaults

`func NewContextSourceRegistrationFragmentWithDefaults() *ContextSourceRegistrationFragment`

NewContextSourceRegistrationFragmentWithDefaults instantiates a new ContextSourceRegistrationFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ContextSourceRegistrationFragment) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContextSourceRegistrationFragment) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContextSourceRegistrationFragment) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ContextSourceRegistrationFragment) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetInformation

`func (o *ContextSourceRegistrationFragment) GetInformation() []RegistrationInfo`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *ContextSourceRegistrationFragment) GetInformationOk() (*[]RegistrationInfo, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *ContextSourceRegistrationFragment) SetInformation(v []RegistrationInfo)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *ContextSourceRegistrationFragment) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetObservationInterval

`func (o *ContextSourceRegistrationFragment) GetObservationInterval() TimeInterval`

GetObservationInterval returns the ObservationInterval field if non-nil, zero value otherwise.

### GetObservationIntervalOk

`func (o *ContextSourceRegistrationFragment) GetObservationIntervalOk() (*TimeInterval, bool)`

GetObservationIntervalOk returns a tuple with the ObservationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationInterval

`func (o *ContextSourceRegistrationFragment) SetObservationInterval(v TimeInterval)`

SetObservationInterval sets ObservationInterval field to given value.

### HasObservationInterval

`func (o *ContextSourceRegistrationFragment) HasObservationInterval() bool`

HasObservationInterval returns a boolean if a field has been set.

### GetManagementInterval

`func (o *ContextSourceRegistrationFragment) GetManagementInterval() TimeInterval`

GetManagementInterval returns the ManagementInterval field if non-nil, zero value otherwise.

### GetManagementIntervalOk

`func (o *ContextSourceRegistrationFragment) GetManagementIntervalOk() (*TimeInterval, bool)`

GetManagementIntervalOk returns a tuple with the ManagementInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementInterval

`func (o *ContextSourceRegistrationFragment) SetManagementInterval(v TimeInterval)`

SetManagementInterval sets ManagementInterval field to given value.

### HasManagementInterval

`func (o *ContextSourceRegistrationFragment) HasManagementInterval() bool`

HasManagementInterval returns a boolean if a field has been set.

### GetLocation

`func (o *ContextSourceRegistrationFragment) GetLocation() Geometry`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContextSourceRegistrationFragment) GetLocationOk() (*Geometry, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContextSourceRegistrationFragment) SetLocation(v Geometry)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContextSourceRegistrationFragment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *ContextSourceRegistrationFragment) GetObservationSpace() Geometry`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *ContextSourceRegistrationFragment) GetObservationSpaceOk() (*Geometry, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *ContextSourceRegistrationFragment) SetObservationSpace(v Geometry)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *ContextSourceRegistrationFragment) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *ContextSourceRegistrationFragment) GetOperationSpace() Geometry`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *ContextSourceRegistrationFragment) GetOperationSpaceOk() (*Geometry, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *ContextSourceRegistrationFragment) SetOperationSpace(v Geometry)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *ContextSourceRegistrationFragment) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetExpires

`func (o *ContextSourceRegistrationFragment) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ContextSourceRegistrationFragment) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ContextSourceRegistrationFragment) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ContextSourceRegistrationFragment) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *ContextSourceRegistrationFragment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextSourceRegistrationFragment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextSourceRegistrationFragment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContextSourceRegistrationFragment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ContextSourceRegistrationFragment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextSourceRegistrationFragment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextSourceRegistrationFragment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContextSourceRegistrationFragment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *ContextSourceRegistrationFragment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ContextSourceRegistrationFragment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ContextSourceRegistrationFragment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ContextSourceRegistrationFragment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


