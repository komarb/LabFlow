# LabFlow
Service for storing laboratory and practice works for RTU MIREA students

## Configuration

File ```src/api/config.json``` must contain next content:

```js
{
  "DbOptions": {
    "host": "mongo", //host to mongodb server
    "port": "27017", //port to mongodb server
    "dbname": "db", //name of db in mongodb
    "subjectsCollectionName" : "subjects", //names of collections
    "usersCollectionName" : "users",
    "tasksCollectionName" : "tasks",
    "reportsCollectionName" : "reports",
    "groupsCollectionName" : "groups"

  },
  "AppOptions": {
    "appPort": "8080", //port for running an app
    "testMode": true|false //bool option for enabling Tests mode
  }
}
```


