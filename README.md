# Discogs v2 RESTful API

A RESTful API for Discogs v2 written in Go

## Endpoints

### Database

* [SCHEME:HOST:PORT]/discogs/v2/artist/?id=
* [SCHEME:HOST:PORT]/discogs/v2/artist/releases/?id=

* [SCHEME:HOST:PORT]/discogs/v2/labels/?id=
* [SCHEME:HOST:PORT]/discogs/v2/labels/releases/?id=

* [SCHEME:HOST:PORT]/discogs/v2/masters/?id=

* [SCHEME:HOST:PORT]/discogs/v2/releases/?id=
* [SCHEME:HOST:PORT]/discogs/v2/releases/rating/?id=

### Search

## Docker

### Build

**Notice**:
Using semantic versioning at the end, which is also tagged at github

```dockerfile
docker build -t deemount/discogs:v0.1.1 .
```

### Run

```dockerfile
docker run --publish 8787:8787 --detach --name dcgs deemount/discogs:v0.1.1  
```

### History

* create achitecture, add functoinalities, first setup
* first upload & initial commit

### To Do's

* github tags
* database connection
* more functionality
* more documentation
