# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | 
**Accept** | Pointer to **string** |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint(uri string, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *Endpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Endpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Endpoint) SetUri(v string)`

SetUri sets Uri field to given value.


### GetAccept

`func (o *Endpoint) GetAccept() string`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *Endpoint) GetAcceptOk() (*string, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *Endpoint) SetAccept(v string)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *Endpoint) HasAccept() bool`

HasAccept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


