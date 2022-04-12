# EntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | **string** |  | 
**IdPattern** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityInfo

`func NewEntityInfo(type_ string, ) *EntityInfo`

NewEntityInfo instantiates a new EntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityInfoWithDefaults

`func NewEntityInfoWithDefaults() *EntityInfo`

NewEntityInfoWithDefaults instantiates a new EntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityInfo) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityInfo) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityInfo) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *EntityInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EntityInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityInfo) SetType(v string)`

SetType sets Type field to given value.


### GetIdPattern

`func (o *EntityInfo) GetIdPattern() string`

GetIdPattern returns the IdPattern field if non-nil, zero value otherwise.

### GetIdPatternOk

`func (o *EntityInfo) GetIdPatternOk() (*string, bool)`

GetIdPatternOk returns a tuple with the IdPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPattern

`func (o *EntityInfo) SetIdPattern(v string)`

SetIdPattern sets IdPattern field to given value.

### HasIdPattern

`func (o *EntityInfo) HasIdPattern() bool`

HasIdPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


