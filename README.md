# focusio-core

## How run it 

Focusio-core is a Cloud Native application, so for testing purposes or to work with it, is recommendable to use like a container. It's recommendable to use the following compose definition:

```
version: '3'

services:
  focusio-core:
    build:
      context: .
      dockerfile: Containerfile
    ports:
      - "8080:8080"
    depends_on:
      - mongo
    environment:
      MONGO_HOSTNAME: mongodb://mongo:27017
      MONGO_USERNAME: root
      MONGO_PASSWORD: OmTDaWEzppo=
     

  mongo:
    image: mongo:latest
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: OmTDaWEzppo=
      MONGO_INITDB_DATABASE: focusio
```

Just save it in a file and execute the following command:
```
podman compose up --build # If you use Podman
docker-compose up --build # If you use Docker
```

However, if your prefer run the application as a binary you can do it follow the next instructions:

The first step is build the application:

```
go build -o focusio-core
```

Focusio Core need a Mongo database to works properly. You can pass your database credentials by environment variables as follows:
```
MONGO_HOSTNAME=$YOUR_MONGO_HOSTNAME MONGO_USERNAME=$YOUR_MONGO_USERNAME MONGO_PASSWORD=$YOUR_MONGO_PASSWORD ./focusio-core
```
