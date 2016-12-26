/*
pREST - Serve a RESTful API from any PostgreSQL database

Install
    go install github.com/nuveo/prest

Run

    Params:

    - PREST\_HTTP_PORT (default 3000)
    - PREST\_PG_HOST (default 127.0.0.1)
    - PREST\_PG_USER
    - PREST\_PG_PASS
    - PREST\_PG_DATABASE
    - PREST\_PG_PORT
    - PREST\_JWT_KEY

    ```
    PREST_PG_USER=postgres PREST_PG_DATABASE=prest PREST_PG_PORT=5432 PREST_HTTP_PORT=3010 prest # Binary installed
    ```
TOML
    Optionally the pREST can be configured by TOML file

    - Set `PREST_CONF` environment variable with file path

    ```toml
    [http]
    port = 6000

    [jwt]
    key = "mysecretkey"

    [pg]
    host = "127.0.0.1"
    user = "postgres"
    pass = "mypass"
    port = 5432
    database = "prest"
    ```

API`s
    HEADER:

    - To start JWT middleware the `PREST_JWT_KEY` environment variable must be set

    ```
    Authorization: JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
    ```

    ### Filter (WHERE) with JSONb field

    ```
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?FIELD->>JSONFIELD:jsonb=VALUE (filter)
    ```

    ### Select - GET

    ```
    http://127.0.0.1:8000/databases (show all databases)
    http://127.0.0.1:8000/schemas (show all schemas)
    http://127.0.0.1:8000/tables (show all tables)
    http://127.0.0.1:8000/DATABASE/SCHEMA (show all tables, find by schema)
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE (show all rows, find by database and table)
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?_page=2&_page_size=10 (pagination, page_size 10 by default)
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?FIELD=VALUE (filter)
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?_renderer=xml (JSON by default)
    ```

    ### Insert - POST

    ```
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE
    ```

    JSON DATA:
    ```
    {
        "data": {
            "FIELD1": "string value",
            "FIELD2": 1234567890
        }
    }
    ```

    ### Update - PATCH/PUT

    Using query string to make filter (WHERE), example:

    ```
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?FIELD1=xyz
    ```

    JSON DATA:
    ```
    {
        "data": {
            "FIELD1": "string value",
            "FIELD2": 1234567890
        }
    }
    ```

    ### Delete - DELETE

    Using query string to make filter (WHERE), example:

    ```
    http://127.0.0.1:8000/DATABASE/SCHEMA/TABLE?FIELD1=xyz
    ```

    ## JOIN

    Using query string to JOIN tables, example:

    ```
    /DATABASE/SCHEMA/TABLE??_join=inner:users:friends.userid:$eq:users.id
    ```

    Parameters:

    1. Join type
    1. Table
    1. Table field 1
    1. Operator (=, <, >, <=, >=)
    1. Table field 2

    Query Operators:

    | Name | Description |
    |-------|-------------|
    | $eq | Matches values that are equal to a specified value.|
    | $gt | Matches values that are greater than a specified value.|
    | $gte | Matches values that are greater than or equal to a specified value.|
    | $lt | Matches values that are less than a specified value.|
    | $lte | Matches values that are less than or equal to a specified value.|
    | $ne | Matches all values that are not equal to a specified value.|
    | $in | Matches any of the values specified in an array.|
    | $nin | Matches none of the values specified in an array.|
*/
package main