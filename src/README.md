# ITLab Reports API
Created with Golang, Gorilla and MongoDB

## Installation using Docker
Install Docker and in project root directory write this code:
```
docker-compose up -d --build
```
If you’re using Docker natively on Linux, Docker Desktop for Mac, or Docker Desktop for Windows, then the server will be running on
```http://localhost:8080```

If you’re using Docker Machine on a Mac or Windows, use ```docker-machine ip MACHINE_VM``` to get the IP address of your Docker host. Then, open ```http://MACHINE_VM_IP:8080``` in a browser

## Config
In config you can set up database charasteristics, like host, port, DB name and collection name as you want
## Requests
You can get Postman requests collection [here](https://www.getpostman.com/collections/4085657bcce140031d0c)

## DB Backup and Restore
To make a backup of a DB, open root folder where MongoDB is installed, open a command promt and type the command mongodump
To restore the backup, open root folder where MongoDB is installed, open a command promt and type the command mongorestore
(All DB paths are default)

