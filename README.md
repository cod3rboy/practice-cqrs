# CQRS and Event Sourcing Practice Project

This project creates a JSON API around patient records.
The write API implements event sourcing based on these immutable patient events - admit patient, transfer patient, discharge patient, update patient name and update patient age. These events are aggregated asynchronously to generate a most recent patient record in the database. The read API just queries the database for these generated record(s).

## Directory Structure

| Directory    | Description                                                                    |
| ------------ | ------------------------------------------------------------------------------ |
| activities   | Contains temporal activities                                                   |
| cmd          | Contains program entrypoint in server subpackage                               |
| config       | Configuration settings for db, temporal and api server                         |
| db           | Database client                                                                |
| ent          | Ent ORM schema and generated files                                             |
| eventstore   | Event store implementation using PostgreSQL                                    |
| handlers     | Container HTTP handlers (subpackages seggregate commands and queries handlers) |
| repositories | Contains repositories and models                                               |
| server       | API server setup                                                               |
| workflows    | Contains temporal workflows                                                    |

## Application Configuration (Environment Variables)

| Variable           | Description                              |
| ------------------ | ---------------------------------------- |
| PSQL_HOST          | Host name for PostgreSQL database        |
| PSQL_PORT          | Port for PostgreSQL database             |
| PSQL_USER          | PostgreSQL database user name            |
| PSQL_PASSWORD      | PostgreSQL database user password        |
| PSQL_DB            | PostgreSQL database name                 |
| GIN_MODE           | `debug` or `release` mode for Gin        |
| ENV                | Application mode: `dev` or `prod`        |
| SERVER_PORT        | Server port to listen on (default: 3001) |
| TEMPORAL_HOST      | Temporal host name                       |
| TEMPORAL_PORT      | Temporal port number                     |
| TEMPORAL_NAMESPACE | Temporal namespace to use                |

## Build/Run (Docker Compose)

Run the following steps to build/run the application -

1. `cp .env.example .env`
2. `docker compose up --build -d`

Now, docker containers for postgresql server, temporal server and the application server should be up and running and you can play with the API endpoints.

You can also access temporal web ui at this address: http://localhost:8080 and see which workflows are being executed.

## Build/Run (Manual)

1. Run `cp .env.example .env` command and set the environment variables for postgresql and temporal in the `.env` file.
2. Build command - `go build -o api ./cmd/server`
3. Run command - `./api`

Now, application server should be up and running and you can play with the API endpoints.

>Note: This manual build/run requires pre-running instances of Temporal and PostgreSQL. The application server will connect with these instances.

## Endpoints

Base URL: http://localhost:3001

The API exposes the following endpoints -

### Commands

**Admit New Patient**

`POST /patient/admit`

Request Body
```json
{
    "name": "Foo Name",
    "ward": 10,
    "age": 25
}
```

The response contains `id` field which is the ID for new patient.

**Transfer Patient to Ward**

`POST /patient/transfer`

Request Body
```json
{
    "id": "<patient-id-here>",
    "ward": 11
}
```
**Discharge Patient**

`POST /patient/discharge`

Request Body
```json
{
    "id": "<patient-id-here>",
}
```
**Update Patient Name**

`POST /patient/updateName`

Request Body
```json
{
    "id": "<patient-id-here>",
    "name": "New Foo Name",
}
```
**Update Patient Age**

`POST /patient/updateAge`

Request Body
```json
{
    "id": "<patient-id-here>",
    "age": 30,
}
```

### Queries

**Get All Patients**

`GET /patient/all`

**Get Patient By ID**

`GET /patient/id/<patient-id-here>`

**Get Patient By Name**

`GET /patient/name/<patient-name>`
