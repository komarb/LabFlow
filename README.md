# LabFlow
Service for storing laboratory and practice works for RTU MIREA students

## Configuration

File ```src/ITLabReports/api/config.json``` must contain next content:

```js
{
  "DbOptions": {
    "host": "mongo", //host to mongodb server
    "port": "27017", //port to mongodb server
    "dbname": "db", //name of db in mongodb
    "collectionName": "collection" //name of collection in mongodb
  },
  "AppOptions": {
    "testMode": true|false //bool option for enabling Tests mode
  }
}
```

File ```src/ITLabReports/api/auth_config.json``` must contain next content:

```js
{
  "AuthOptions": {
    "keyUrl": "https://examplesite/files/jwks.json", //url to jwks.json
    "audience": "example_audience", //audince for JWT
    "issuer" : "https://exampleissuersite.com", //issuer for JWT
    "scope" : "my_scope" //required scope for JWT
  }
}
```

