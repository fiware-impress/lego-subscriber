# Go API client for openapi

This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: latest
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BatchOperationsApi* | [**BatchEntityCreation**](docs/BatchOperationsApi.md#batchentitycreation) | **Post** /entityOperations/create | 
*BatchOperationsApi* | [**BatchEntityDelete**](docs/BatchOperationsApi.md#batchentitydelete) | **Post** /entityOperations/delete | 
*BatchOperationsApi* | [**BatchEntityUpdate**](docs/BatchOperationsApi.md#batchentityupdate) | **Post** /entityOperations/update | 
*BatchOperationsApi* | [**BatchEntityUpsert**](docs/BatchOperationsApi.md#batchentityupsert) | **Post** /entityOperations/upsert | 
*CSourceRegistrationsApi* | [**QueryCsources**](docs/CSourceRegistrationsApi.md#querycsources) | **Get** /csourceRegistrations/ | 
*CSourceRegistrationsApi* | [**RegisterCsource**](docs/CSourceRegistrationsApi.md#registercsource) | **Post** /csourceRegistrations/ | 
*CSourceRegistrationsApi* | [**RemoveCsource**](docs/CSourceRegistrationsApi.md#removecsource) | **Delete** /csourceRegistrations/{registrationId} | 
*CSourceRegistrationsApi* | [**RetrieveCsource**](docs/CSourceRegistrationsApi.md#retrievecsource) | **Get** /csourceRegistrations/{registrationId} | 
*CSourceSubscriptionsApi* | [**CreateCSourceSubscription**](docs/CSourceSubscriptionsApi.md#createcsourcesubscription) | **Post** /csourceSubscriptions/ | 
*CSourceSubscriptionsApi* | [**RemoveCSourceSubscription**](docs/CSourceSubscriptionsApi.md#removecsourcesubscription) | **Delete** /csourceSubscriptions/{subscriptionId} | 
*CSourceSubscriptionsApi* | [**RetrieveCSourceSubscriptions**](docs/CSourceSubscriptionsApi.md#retrievecsourcesubscriptions) | **Get** /csourceSubscriptions/ | 
*CSourceSubscriptionsApi* | [**RetrieveCSourceSubscriptionsById**](docs/CSourceSubscriptionsApi.md#retrievecsourcesubscriptionsbyid) | **Get** /csourceSubscriptions/{subscriptionId} | 
*CSourceSubscriptionsApi* | [**UpdateCSourceSubscription**](docs/CSourceSubscriptionsApi.md#updatecsourcesubscription) | **Patch** /csourceSubscriptions/{subscriptionId} | 
*ContextInformationApi* | [**AppendEntityAttrs**](docs/ContextInformationApi.md#appendentityattrs) | **Post** /entities/{entityId}/attrs/ | 
*ContextInformationApi* | [**CreateEntity**](docs/ContextInformationApi.md#createentity) | **Post** /entities/ | 
*ContextInformationApi* | [**PartialAttrUpdate**](docs/ContextInformationApi.md#partialattrupdate) | **Patch** /entities/{entityId}/attrs/{attrId} | 
*ContextInformationApi* | [**QueryEntities**](docs/ContextInformationApi.md#queryentities) | **Get** /entities/ | 
*ContextInformationApi* | [**RemoveEntityAttr**](docs/ContextInformationApi.md#removeentityattr) | **Delete** /entities/{entityId}/attrs/{attrId} | 
*ContextInformationApi* | [**RemoveEntityById**](docs/ContextInformationApi.md#removeentitybyid) | **Delete** /entities/{entityId} | 
*ContextInformationApi* | [**RetrieveEntityById**](docs/ContextInformationApi.md#retrieveentitybyid) | **Get** /entities/{entityId} | 
*ContextInformationApi* | [**UpdateEntityAttrs**](docs/ContextInformationApi.md#updateentityattrs) | **Patch** /entities/{entityId}/attrs/ | 
*ContextSourcesApi* | [**CreateCSourceSubscription**](docs/ContextSourcesApi.md#createcsourcesubscription) | **Post** /csourceSubscriptions/ | 
*ContextSourcesApi* | [**QueryCsources**](docs/ContextSourcesApi.md#querycsources) | **Get** /csourceRegistrations/ | 
*ContextSourcesApi* | [**RegisterCsource**](docs/ContextSourcesApi.md#registercsource) | **Post** /csourceRegistrations/ | 
*ContextSourcesApi* | [**RemoveCSourceSubscription**](docs/ContextSourcesApi.md#removecsourcesubscription) | **Delete** /csourceSubscriptions/{subscriptionId} | 
*ContextSourcesApi* | [**RemoveCsource**](docs/ContextSourcesApi.md#removecsource) | **Delete** /csourceRegistrations/{registrationId} | 
*ContextSourcesApi* | [**RetrieveCSourceSubscriptions**](docs/ContextSourcesApi.md#retrievecsourcesubscriptions) | **Get** /csourceSubscriptions/ | 
*ContextSourcesApi* | [**RetrieveCSourceSubscriptionsById**](docs/ContextSourcesApi.md#retrievecsourcesubscriptionsbyid) | **Get** /csourceSubscriptions/{subscriptionId} | 
*ContextSourcesApi* | [**RetrieveCsource**](docs/ContextSourcesApi.md#retrievecsource) | **Get** /csourceRegistrations/{registrationId} | 
*ContextSourcesApi* | [**UpdateCSourceSubscription**](docs/ContextSourcesApi.md#updatecsourcesubscription) | **Patch** /csourceSubscriptions/{subscriptionId} | 
*ContextSubscriptionApi* | [**CreateSubscription**](docs/ContextSubscriptionApi.md#createsubscription) | **Post** /subscriptions/ | 
*ContextSubscriptionApi* | [**RemoveSubscription**](docs/ContextSubscriptionApi.md#removesubscription) | **Delete** /subscriptions/{subscriptionId} | 
*ContextSubscriptionApi* | [**RetrieveSubscriptionById**](docs/ContextSubscriptionApi.md#retrievesubscriptionbyid) | **Get** /subscriptions/{subscriptionId} | 
*ContextSubscriptionApi* | [**RetrieveSubscriptions**](docs/ContextSubscriptionApi.md#retrievesubscriptions) | **Get** /subscriptions/ | 
*ContextSubscriptionApi* | [**UpdateSubscription**](docs/ContextSubscriptionApi.md#updatesubscription) | **Patch** /subscriptions/{subscriptionId} | 
*EntitiesApi* | [**AppendEntityAttrs**](docs/EntitiesApi.md#appendentityattrs) | **Post** /entities/{entityId}/attrs/ | 
*EntitiesApi* | [**BatchEntityCreation**](docs/EntitiesApi.md#batchentitycreation) | **Post** /entityOperations/create | 
*EntitiesApi* | [**BatchEntityDelete**](docs/EntitiesApi.md#batchentitydelete) | **Post** /entityOperations/delete | 
*EntitiesApi* | [**BatchEntityUpdate**](docs/EntitiesApi.md#batchentityupdate) | **Post** /entityOperations/update | 
*EntitiesApi* | [**BatchEntityUpsert**](docs/EntitiesApi.md#batchentityupsert) | **Post** /entityOperations/upsert | 
*EntitiesApi* | [**CreateEntity**](docs/EntitiesApi.md#createentity) | **Post** /entities/ | 
*EntitiesApi* | [**PartialAttrUpdate**](docs/EntitiesApi.md#partialattrupdate) | **Patch** /entities/{entityId}/attrs/{attrId} | 
*EntitiesApi* | [**QueryEntities**](docs/EntitiesApi.md#queryentities) | **Get** /entities/ | 
*EntitiesApi* | [**RemoveEntityAttr**](docs/EntitiesApi.md#removeentityattr) | **Delete** /entities/{entityId}/attrs/{attrId} | 
*EntitiesApi* | [**RemoveEntityById**](docs/EntitiesApi.md#removeentitybyid) | **Delete** /entities/{entityId} | 
*EntitiesApi* | [**RetrieveEntityById**](docs/EntitiesApi.md#retrieveentitybyid) | **Get** /entities/{entityId} | 
*EntitiesApi* | [**UpdateEntityAttrs**](docs/EntitiesApi.md#updateentityattrs) | **Patch** /entities/{entityId}/attrs/ | 
*SubscriptionsApi* | [**CreateSubscription**](docs/SubscriptionsApi.md#createsubscription) | **Post** /subscriptions/ | 
*SubscriptionsApi* | [**RemoveSubscription**](docs/SubscriptionsApi.md#removesubscription) | **Delete** /subscriptions/{subscriptionId} | 
*SubscriptionsApi* | [**RetrieveSubscriptionById**](docs/SubscriptionsApi.md#retrievesubscriptionbyid) | **Get** /subscriptions/{subscriptionId} | 
*SubscriptionsApi* | [**RetrieveSubscriptions**](docs/SubscriptionsApi.md#retrievesubscriptions) | **Get** /subscriptions/ | 
*SubscriptionsApi* | [**UpdateSubscription**](docs/SubscriptionsApi.md#updatesubscription) | **Patch** /subscriptions/{subscriptionId} | 
*TemporalApi* | [**AddTemporalEntityAttrs**](docs/TemporalApi.md#addtemporalentityattrs) | **Post** /temporal/entities/{entityId}/attrs/ | 
*TemporalApi* | [**CreateUpdateEntityTemporal**](docs/TemporalApi.md#createupdateentitytemporal) | **Post** /temporal/entities/ | 
*TemporalApi* | [**ModifyEntityTemporalAttrInstance**](docs/TemporalApi.md#modifyentitytemporalattrinstance) | **Patch** /temporal/entities/{entityId}/attrs/{attrId}/{instanceId} | 
*TemporalApi* | [**QueryTemporalEntities**](docs/TemporalApi.md#querytemporalentities) | **Get** /temporal/entities/ | 
*TemporalApi* | [**QueryTemporalEntitiesOnPost**](docs/TemporalApi.md#querytemporalentitiesonpost) | **Post** /temporal/entityOperations/query | 
*TemporalApi* | [**RemoveEntityTemporalAttr**](docs/TemporalApi.md#removeentitytemporalattr) | **Delete** /temporal/entities/{entityId}/attrs/{attrId} | 
*TemporalApi* | [**RemoveEntityTemporalAttrInstance**](docs/TemporalApi.md#removeentitytemporalattrinstance) | **Delete** /temporal/entities/{entityId}/attrs/{attrId}/{instanceId} | 
*TemporalApi* | [**RemoveEntityTemporalById**](docs/TemporalApi.md#removeentitytemporalbyid) | **Delete** /temporal/entities/{entityId} | 
*TemporalApi* | [**RetrieveEntityTemporalById**](docs/TemporalApi.md#retrieveentitytemporalbyid) | **Get** /temporal/entities/{entityId} | 
*TemporalEvolutionApi* | [**AddTemporalEntityAttrs**](docs/TemporalEvolutionApi.md#addtemporalentityattrs) | **Post** /temporal/entities/{entityId}/attrs/ | 
*TemporalEvolutionApi* | [**CreateUpdateEntityTemporal**](docs/TemporalEvolutionApi.md#createupdateentitytemporal) | **Post** /temporal/entities/ | 
*TemporalEvolutionApi* | [**ModifyEntityTemporalAttrInstance**](docs/TemporalEvolutionApi.md#modifyentitytemporalattrinstance) | **Patch** /temporal/entities/{entityId}/attrs/{attrId}/{instanceId} | 
*TemporalEvolutionApi* | [**QueryTemporalEntities**](docs/TemporalEvolutionApi.md#querytemporalentities) | **Get** /temporal/entities/ | 
*TemporalEvolutionApi* | [**QueryTemporalEntitiesOnPost**](docs/TemporalEvolutionApi.md#querytemporalentitiesonpost) | **Post** /temporal/entityOperations/query | 
*TemporalEvolutionApi* | [**RemoveEntityTemporalAttr**](docs/TemporalEvolutionApi.md#removeentitytemporalattr) | **Delete** /temporal/entities/{entityId}/attrs/{attrId} | 
*TemporalEvolutionApi* | [**RemoveEntityTemporalAttrInstance**](docs/TemporalEvolutionApi.md#removeentitytemporalattrinstance) | **Delete** /temporal/entities/{entityId}/attrs/{attrId}/{instanceId} | 
*TemporalEvolutionApi* | [**RemoveEntityTemporalById**](docs/TemporalEvolutionApi.md#removeentitytemporalbyid) | **Delete** /temporal/entities/{entityId} | 
*TemporalEvolutionApi* | [**RetrieveEntityTemporalById**](docs/TemporalEvolutionApi.md#retrieveentitytemporalbyid) | **Get** /temporal/entities/{entityId} | 
*TemporalRetrievalApi* | [**QueryTemporalEntities**](docs/TemporalRetrievalApi.md#querytemporalentities) | **Get** /temporal/entities/ | 
*TemporalRetrievalApi* | [**QueryTemporalEntitiesOnPost**](docs/TemporalRetrievalApi.md#querytemporalentitiesonpost) | **Post** /temporal/entityOperations/query | 
*TemporalRetrievalApi* | [**RetrieveEntityTemporalById**](docs/TemporalRetrievalApi.md#retrieveentitytemporalbyid) | **Get** /temporal/entities/{entityId} | 


## Documentation For Models

 - [BatchEntityError](docs/BatchEntityError.md)
 - [BatchOperationResult](docs/BatchOperationResult.md)
 - [ContextSourceRegistration](docs/ContextSourceRegistration.md)
 - [ContextSourceRegistrationAllOf](docs/ContextSourceRegistrationAllOf.md)
 - [ContextSourceRegistrationFragment](docs/ContextSourceRegistrationFragment.md)
 - [Coordinates](docs/Coordinates.md)
 - [Endpoint](docs/Endpoint.md)
 - [Entity](docs/Entity.md)
 - [EntityFragment](docs/EntityFragment.md)
 - [EntityInfo](docs/EntityInfo.md)
 - [EntityTemporal](docs/EntityTemporal.md)
 - [EntityTemporalFragment](docs/EntityTemporalFragment.md)
 - [GeoProperty](docs/GeoProperty.md)
 - [GeoQuery](docs/GeoQuery.md)
 - [Geometry](docs/Geometry.md)
 - [GeorelEnum](docs/GeorelEnum.md)
 - [LineString](docs/LineString.md)
 - [MultiLineString](docs/MultiLineString.md)
 - [MultiPoint](docs/MultiPoint.md)
 - [MultiPolygon](docs/MultiPolygon.md)
 - [NotUpdatedDetails](docs/NotUpdatedDetails.md)
 - [NotificationParams](docs/NotificationParams.md)
 - [Point](docs/Point.md)
 - [Polygon](docs/Polygon.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [Property](docs/Property.md)
 - [Query](docs/Query.md)
 - [RegistrationInfo](docs/RegistrationInfo.md)
 - [Relationship](docs/Relationship.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionAllOf](docs/SubscriptionAllOf.md)
 - [SubscriptionFragment](docs/SubscriptionFragment.md)
 - [TemporalQuery](docs/TemporalQuery.md)
 - [TimeInterval](docs/TimeInterval.md)
 - [Timerel](docs/Timerel.md)
 - [UpdateResult](docs/UpdateResult.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

NGSI-LD@etsi.org
