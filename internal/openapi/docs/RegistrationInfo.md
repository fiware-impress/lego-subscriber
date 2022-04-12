# RegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]EntityInfo**](EntityInfo.md) |  | [optional] 
**Properties** | Pointer to **[]string** |  | [optional] 
**Relationships** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRegistrationInfo

`func NewRegistrationInfo() *RegistrationInfo`

NewRegistrationInfo instantiates a new RegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInfoWithDefaults

`func NewRegistrationInfoWithDefaults() *RegistrationInfo`

NewRegistrationInfoWithDefaults instantiates a new RegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RegistrationInfo) GetEntities() []EntityInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RegistrationInfo) GetEntitiesOk() (*[]EntityInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RegistrationInfo) SetEntities(v []EntityInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RegistrationInfo) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetProperties

`func (o *RegistrationInfo) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RegistrationInfo) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RegistrationInfo) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *RegistrationInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRelationships

`func (o *RegistrationInfo) GetRelationships() []string`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RegistrationInfo) GetRelationshipsOk() (*[]string, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RegistrationInfo) SetRelationships(v []string)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RegistrationInfo) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


