# go-common-plugin

Collection of helper function needed when building scalable backend application using go languange.

#### Response Data
```
A standard model to return API Response. Consist of :
- Code -> indicate the http status of the response using int variable
- Status -> display a http status / error status using string. you can custom an error status using this field.
- Errors -> details of errors returned by the API. usually if you have multiple field validation,
            the map key will be a field name, and the value will be list of errors
- Data -> A generic type to store the response value returned by the API.
```

#### Exception
```
- Wrap the error status and message. This useful for handler/controller 
for mapping errors by custom http status.
```
