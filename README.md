# keyghost
A Keycloak Mock Server



# Sample Server Profile Routes:

* Wellknown: http://localhost:8080/auth/realms/demo/.well-known/openid-configuration 
* Authorization: http://localhost:8080/auth/realms/demo/protocol/openid-connect/auth?state=my-state-123

### Exchange Authorization Code to JWT Tokens:

```shell
$ curl http://localhost:8080/auth/realms/demo/protocol/openid-connect/token -d "grant_type=authorization_code" -d "code=546c9b63-a5f6-41d4-ae5e-1d4b44914ef7"
```
