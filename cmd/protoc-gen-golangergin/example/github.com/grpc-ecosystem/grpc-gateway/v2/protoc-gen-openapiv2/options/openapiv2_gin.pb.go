// Code generated. DO NOT EDIT.

package options

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	runtime "github.com/jiandahao/golanger/pkg/generator/gingen/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// `Swagger` is a representation of OpenAPI v2 specification's Swagger object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#swaggerObject
//
// Example:
//
//  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//    info: {
//      title: "Echo API";
//      version: "1.0";
//      description: ";
//      contact: {
//        name: "gRPC-Gateway project";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway";
//        email: "none@example.com";
//      };
//      license: {
//        name: "BSD 3-Clause License";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway/blob/master/LICENSE.txt";
//      };
//    };
//    schemes: HTTPS;
//    consumes: "application/json";
//    produces: "application/json";
//  };
//
type Swagger struct {
	Swagger             string                     `json:"swagger,omitempty"`
	Info                *Info                      `json:"info,omitempty"`
	Host                string                     `json:"host,omitempty"`
	BasePath            string                     `json:"base_path,omitempty"`
	Schemes             []Scheme                   `json:"schemes,omitempty"`
	Consumes            []string                   `json:"consumes,omitempty"`
	Produces            []string                   `json:"produces,omitempty"`
	Responses           map[string]*Response       `json:"responses,omitempty"`
	SecurityDefinitions *SecurityDefinitions       `json:"security_definitions,omitempty"`
	Security            []*SecurityRequirement     `json:"security,omitempty"`
	ExternalDocs        *ExternalDocumentation     `json:"external_docs,omitempty"`
	Extensions          map[string]*structpb.Value `json:"extensions,omitempty"`
}

// `Operation` is a representation of OpenAPI v2 specification's Operation object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#operationObject
//
// Example:
//
//  service EchoService {
//    rpc Echo(SimpleMessage) returns (SimpleMessage) {
//      option (google.api.http) = {
//        get: "/v1/example/echo/{id}"
//      };
//
//      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
//        summary: "Get a message.";
//        operation_id: "getMessage";
//        tags: "echo";
//        responses: {
//          key: "200"
//            value: {
//            description: "OK";
//          }
//        }
//      };
//    }
//  }
type Operation struct {
	Tags         []string                   `json:"tags,omitempty"`
	Summary      string                     `json:"summary,omitempty"`
	Description  string                     `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation     `json:"external_docs,omitempty"`
	OperationId  string                     `json:"operation_id,omitempty"`
	Consumes     []string                   `json:"consumes,omitempty"`
	Produces     []string                   `json:"produces,omitempty"`
	Responses    map[string]*Response       `json:"responses,omitempty"`
	Schemes      []Scheme                   `json:"schemes,omitempty"`
	Deprecated   bool                       `json:"deprecated,omitempty"`
	Security     []*SecurityRequirement     `json:"security,omitempty"`
	Extensions   map[string]*structpb.Value `json:"extensions,omitempty"`
}

// `Header` is a representation of OpenAPI v2 specification's Header object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#headerObject
//
type Header struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Default     string `json:"default,omitempty"`
	Pattern     string `json:"pattern,omitempty"`
}

// `Response` is a representation of OpenAPI v2 specification's Response object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#responseObject
//
type Response struct {
	Description string                     `json:"description,omitempty"`
	Schema      *Schema                    `json:"schema,omitempty"`
	Headers     map[string]*Header         `json:"headers,omitempty"`
	Examples    map[string]string          `json:"examples,omitempty"`
	Extensions  map[string]*structpb.Value `json:"extensions,omitempty"`
}

// `Info` is a representation of OpenAPI v2 specification's Info object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#infoObject
//
// Example:
//
//  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//    info: {
//      title: "Echo API";
//      version: "1.0";
//      description: ";
//      contact: {
//        name: "gRPC-Gateway project";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway";
//        email: "none@example.com";
//      };
//      license: {
//        name: "BSD 3-Clause License";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway/blob/master/LICENSE.txt";
//      };
//    };
//    ...
//  };
//
type Info struct {
	Title          string                     `json:"title,omitempty"`
	Description    string                     `json:"description,omitempty"`
	TermsOfService string                     `json:"terms_of_service,omitempty"`
	Contact        *Contact                   `json:"contact,omitempty"`
	License        *License                   `json:"license,omitempty"`
	Version        string                     `json:"version,omitempty"`
	Extensions     map[string]*structpb.Value `json:"extensions,omitempty"`
}

// `Contact` is a representation of OpenAPI v2 specification's Contact object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#contactObject
//
// Example:
//
//  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//    info: {
//      ...
//      contact: {
//        name: "gRPC-Gateway project";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway";
//        email: "none@example.com";
//      };
//      ...
//    };
//    ...
//  };
//
type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// `License` is a representation of OpenAPI v2 specification's License object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#licenseObject
//
// Example:
//
//  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//    info: {
//      ...
//      license: {
//        name: "BSD 3-Clause License";
//        url: "https://github.com/grpc-ecosystem/grpc-gateway/blob/master/LICENSE.txt";
//      };
//      ...
//    };
//    ...
//  };
//
type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

// `ExternalDocumentation` is a representation of OpenAPI v2 specification's
// ExternalDocumentation object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#externalDocumentationObject
//
// Example:
//
//  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
//    ...
//    external_docs: {
//      description: "More about gRPC-Gateway";
//      url: "https://github.com/grpc-ecosystem/grpc-gateway";
//    }
//    ...
//  };
//
type ExternalDocumentation struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

// `Schema` is a representation of OpenAPI v2 specification's Schema object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject
//
type Schema struct {
	JsonSchema    *JSONSchema            `json:"json_schema,omitempty"`
	Discriminator string                 `json:"discriminator,omitempty"`
	ReadOnly      bool                   `json:"read_only,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"external_docs,omitempty"`
	Example       string                 `json:"example,omitempty"`
}

// `JSONSchema` represents properties from JSON Schema taken, and as used, in
// the OpenAPI v2 spec.
//
// This includes changes made by OpenAPI v2.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject
//
// See also: https://cswr.github.io/JsonSchema/spec/basic_types/,
// https://github.com/json-schema-org/json-schema-spec/blob/master/schema.json
//
// Example:
//
//  message SimpleMessage {
//    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
//      json_schema: {
//        title: "SimpleMessage"
//        description: "A simple message."
//        required: ["id"]
//      }
//    };
//
//    // Id represents the message identifier.
//    string id = 1; [
//        (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
//          description: "The unique identifier of the simple message."
//        }];
//  }
//
type JSONSchema struct {
	Ref                string                             `json:"ref,omitempty"`
	Title              string                             `json:"title,omitempty"`
	Description        string                             `json:"description,omitempty"`
	Default            string                             `json:"default,omitempty"`
	ReadOnly           bool                               `json:"read_only,omitempty"`
	Example            string                             `json:"example,omitempty"`
	MultipleOf         float64                            `json:"multiple_of,omitempty"`
	Maximum            float64                            `json:"maximum,omitempty"`
	ExclusiveMaximum   bool                               `json:"exclusive_maximum,omitempty"`
	Minimum            float64                            `json:"minimum,omitempty"`
	ExclusiveMinimum   bool                               `json:"exclusive_minimum,omitempty"`
	MaxLength          uint64                             `json:"max_length,omitempty"`
	MinLength          uint64                             `json:"min_length,omitempty"`
	Pattern            string                             `json:"pattern,omitempty"`
	MaxItems           uint64                             `json:"max_items,omitempty"`
	MinItems           uint64                             `json:"min_items,omitempty"`
	UniqueItems        bool                               `json:"unique_items,omitempty"`
	MaxProperties      uint64                             `json:"max_properties,omitempty"`
	MinProperties      uint64                             `json:"min_properties,omitempty"`
	Required           []string                           `json:"required,omitempty"`
	Array              []string                           `json:"array,omitempty"`
	Type               []JSONSchema_JSONSchemaSimpleTypes `json:"type,omitempty"`
	Format             string                             `json:"format,omitempty"`
	Enum               []string                           `json:"enum,omitempty"`
	FieldConfiguration *JSONSchema_FieldConfiguration     `json:"field_configuration,omitempty"`
}

// `Tag` is a representation of OpenAPI v2 specification's Tag object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#tagObject
//
type Tag struct {
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"external_docs,omitempty"`
}

// `SecurityDefinitions` is a representation of OpenAPI v2 specification's
// Security Definitions object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityDefinitionsObject
//
// A declaration of the security schemes available to be used in the
// specification. This does not enforce the security schemes on the operations
// and only serves to provide the relevant details for each scheme.
type SecurityDefinitions struct {
	Security map[string]*SecurityScheme `json:"security,omitempty"`
}

// `SecurityScheme` is a representation of OpenAPI v2 specification's
// Security Scheme object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securitySchemeObject
//
// Allows the definition of a security scheme that can be used by the
// operations. Supported schemes are basic authentication, an API key (either as
// a header or as a query parameter) and OAuth2's common flows (implicit,
// password, application and access code).
type SecurityScheme struct {
	Type             SecurityScheme_Type        `json:"type,omitempty"`
	Description      string                     `json:"description,omitempty"`
	Name             string                     `json:"name,omitempty"`
	In               SecurityScheme_In          `json:"in,omitempty"`
	Flow             SecurityScheme_Flow        `json:"flow,omitempty"`
	AuthorizationUrl string                     `json:"authorization_url,omitempty"`
	TokenUrl         string                     `json:"token_url,omitempty"`
	Scopes           *Scopes                    `json:"scopes,omitempty"`
	Extensions       map[string]*structpb.Value `json:"extensions,omitempty"`
}

// `SecurityRequirement` is a representation of OpenAPI v2 specification's
// Security Requirement object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityRequirementObject
//
// Lists the required security schemes to execute this operation. The object can
// have multiple security schemes declared in it which are all required (that
// is, there is a logical AND between the schemes).
//
// The name used for each property MUST correspond to a security scheme
// declared in the Security Definitions.
type SecurityRequirement struct {
	SecurityRequirement map[string]*SecurityRequirement_SecurityRequirementValue `json:"security_requirement,omitempty"`
}

// `Scopes` is a representation of OpenAPI v2 specification's Scopes object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#scopesObject
//
// Lists the available scopes for an OAuth2 security scheme.
type Scopes struct {
	Scope map[string]string `json:"scope,omitempty"`
}

type Swagger_ResponsesEntry struct {
	Key   string    `json:"key,omitempty"`
	Value *Response `json:"value,omitempty"`
}

type Swagger_ExtensionsEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *structpb.Value `json:"value,omitempty"`
}

type Operation_ResponsesEntry struct {
	Key   string    `json:"key,omitempty"`
	Value *Response `json:"value,omitempty"`
}

type Operation_ExtensionsEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *structpb.Value `json:"value,omitempty"`
}

type Response_HeadersEntry struct {
	Key   string  `json:"key,omitempty"`
	Value *Header `json:"value,omitempty"`
}

type Response_ExamplesEntry struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Response_ExtensionsEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *structpb.Value `json:"value,omitempty"`
}

type Info_ExtensionsEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *structpb.Value `json:"value,omitempty"`
}

// 'FieldConfiguration' provides additional field level properties used when generating the OpenAPI v2 file.
// These properties are not defined by OpenAPIv2, but they are used to control the generation.
type JSONSchema_FieldConfiguration struct {
	PathParamName string `json:"path_param_name,omitempty"`
}

type SecurityDefinitions_SecurityEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *SecurityScheme `json:"value,omitempty"`
}

type SecurityScheme_ExtensionsEntry struct {
	Key   string          `json:"key,omitempty"`
	Value *structpb.Value `json:"value,omitempty"`
}

// If the security scheme is of type "oauth2", then the value is a list of
// scope names required for the execution. For other security scheme types,
// the array MUST be empty.
type SecurityRequirement_SecurityRequirementValue struct {
	Scope []string `json:"scope,omitempty"`
}

type SecurityRequirement_SecurityRequirementEntry struct {
	Key   string                                        `json:"key,omitempty"`
	Value *SecurityRequirement_SecurityRequirementValue `json:"value,omitempty"`
}

type Scopes_ScopeEntry struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// Scheme describes the schemes supported by the OpenAPI Swagger
// and Operation objects.
type Scheme int32

const (
	Scheme_UNKNOWN Scheme = 0
	Scheme_HTTP    Scheme = 1
	Scheme_HTTPS   Scheme = 2
	Scheme_WS      Scheme = 3
	Scheme_WSS     Scheme = 4
)

type JSONSchema_JSONSchemaSimpleTypes int32

const (
	JSONSchema_UNKNOWN JSONSchema_JSONSchemaSimpleTypes = 0
	JSONSchema_ARRAY   JSONSchema_JSONSchemaSimpleTypes = 1
	JSONSchema_BOOLEAN JSONSchema_JSONSchemaSimpleTypes = 2
	JSONSchema_INTEGER JSONSchema_JSONSchemaSimpleTypes = 3
	JSONSchema_NULL    JSONSchema_JSONSchemaSimpleTypes = 4
	JSONSchema_NUMBER  JSONSchema_JSONSchemaSimpleTypes = 5
	JSONSchema_OBJECT  JSONSchema_JSONSchemaSimpleTypes = 6
	JSONSchema_STRING  JSONSchema_JSONSchemaSimpleTypes = 7
)

// The type of the security scheme. Valid values are "basic",
// "apiKey" or "oauth2".
type SecurityScheme_Type int32

const (
	SecurityScheme_TYPE_INVALID SecurityScheme_Type = 0
	SecurityScheme_TYPE_BASIC   SecurityScheme_Type = 1
	SecurityScheme_TYPE_API_KEY SecurityScheme_Type = 2
	SecurityScheme_TYPE_OAUTH2  SecurityScheme_Type = 3
)

// The location of the API key. Valid values are "query" or "header".
type SecurityScheme_In int32

const (
	SecurityScheme_IN_INVALID SecurityScheme_In = 0
	SecurityScheme_IN_QUERY   SecurityScheme_In = 1
	SecurityScheme_IN_HEADER  SecurityScheme_In = 2
)

// The flow used by the OAuth2 security scheme. Valid values are
// "implicit", "password", "application" or "accessCode".
type SecurityScheme_Flow int32

const (
	SecurityScheme_FLOW_INVALID     SecurityScheme_Flow = 0
	SecurityScheme_FLOW_IMPLICIT    SecurityScheme_Flow = 1
	SecurityScheme_FLOW_PASSWORD    SecurityScheme_Flow = 2
	SecurityScheme_FLOW_APPLICATION SecurityScheme_Flow = 3
	SecurityScheme_FLOW_ACCESS_CODE SecurityScheme_Flow = 4
)
