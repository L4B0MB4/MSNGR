# MSNGR

### Regarding the exercise
I decided to use a synchronous API for the sake of time and complexity as described in the task. Optimally we would add the messages to a queue and responde with a 
hypermedia-api-response including a link to request the status of a request since the calls to the communication-provider can vary in execution time. 

The response could've looked something like this in an async case
```json
{
    "data":{
        "information": "Request was accepted"
    },
    "links":{
        "status":"/messages/{myuniqueId}/status"
    }
}
```

Additionally I built the system enhanceable in the areas where i saw the most potential: "Rules for forwarding the message" and "channels to send it to".

For testing: The integration tests will fail if the .env file is not present or the env variables are not set. I focused on integration tests since the
exercise did not suggest unit tests

I also skipped graceful shutdown of the api server for time's-sake


## Documentation

### Used Packages

Gin -> Webframework. Has to be checked for a proper application since it uses unstable packages

Zerolog -> Logging.

Google's UUID -> UUID Generation for tracing.

### Could've been used
Google's Wire -> Dependency Injection via Code Generation

"swaggo/swag" -> OpenAPI Document Generation

Opentelemetry -> Metrics & Tracing 

### Implementation Details

The application is a synchronous API built on a somewhat "clean" architecture. The request enters through a controller `MessageController` which has a dependency on
a struct that holds business logic `ForwardingProvider` with additional dependencies on `ForwardingRules`. The `ForwardingProvider` executes the rule that was injected
in the "constructor" whenever there is a message that might need forwarding. The default `ForwardingRule` filters for the message type being `Warning` and also filters on
the discord `CommunicationProvider (CP)`. It returns a potentially reduced liste of `CPs` to the `ForwardingProvider` which then commands all of them to send the messages
via their specific channels.

### Enhancements

#### New Channels
Other channels besides discord can be added by simply adding implementations of `CommunicationProvider` and adding them to the DependencyInjection. To have them actively
used the `ForwardingRule` needs to keep them in the filtered return value tho. 

#### Customer-Specific Rules
Adding customer specific rules might be a bit more difficult since one has to introduce a database that contains the rules that a customer defines. It also depends on 
the complexity of the rules. But besides a storage for the custom rules we only need a different implementation of `ForwardingRule`. We could have a factory that creates
a rule per customer or a rule that actively requests the rule-configuration from the database on-request. 


### API Definition

(For a proper application i recommend code generated API definitions)

#### Request

`http://localhost:6616/api/{customerId}/messages` for local development
```json
{ 
    "Type": "String e.g. 'Warning'", 
    "Name": "String", 
    "Description": "String"
} 
```

#### Responses

**Validation Error (400)** 
```json
{
    "Data": null,
    "Errors": [
        {
            "Id": "uuid",
            "Detail": "string"
        }
    ]
}
```

**Unkown Error (500)** 
```json
{
    "Data": null,
    "Errors": [
        {
            "Id": "uuid",
            "Detail": "string"
        }
    ]
}
```

**No Provider found (200)** 
```json
{
    "Data": {
        "Description": "string"
    },
    "Errors": []
}
```