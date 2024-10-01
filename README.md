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




## Documentation