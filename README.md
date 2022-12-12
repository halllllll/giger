# what's giger

A HUMAN DOING GIGA-SCHOOL ACTIONS.

# usage

## setting envs
example, filled some `.env` or `secret.json` files. There are some service credential or admin ID/PW, or properties.

## exec docker-compose
`docker compose up -d --build`

## check DB IP Address (docker network)
example

- Login `metabase` container. I think VScode Remote Explorer is super easy way. OR, Docker Desktop is best way, just only go to Terminal tab.
- Command `ping` to `db`. `db` is alias of PostgresSQL container defined by `docker-compose.yml`.
- and you get, like under results.
    ```
    / # ping db
    PING db (172.28.0.2): 56 data bytes
    64 bytes from 172.28.0.2: seq=0 ttl=64 time=0.322 ms
    64 bytes from 172.28.0.2: seq=1 ttl=64 time=0.187 ms
    64 bytes from 172.28.0.2: seq=2 ttl=64 time=0.634 ms
    64 bytes from 172.28.0.2: seq=3 ttl=64 time=0.108 ms
    ```
- Of course, here `(172.28.0.2)` is exactly the IP address you should specify after you login metabase.