# BatchEntityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**ProblemDetails**](ProblemDetails.md) |  | [optional] 

## Methods

### NewBatchEntityError

`func NewBatchEntityError() *BatchEntityError`

NewBatchEntityError instantiates a new BatchEntityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchEntityErrorWithDefaults

`func NewBatchEntityErrorWithDefaults() *BatchEntityError`

NewBatchEntityErrorWithDefaults instantiates a new BatchEntityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *BatchEntityError) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BatchEntityError) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BatchEntityError) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *BatchEntityError) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetError

`func (o *BatchEntityError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchEntityError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchEntityError) SetError(v ProblemDetails)`

SetError sets Error field to given value.

### HasError

`func (o *BatchEntityError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


