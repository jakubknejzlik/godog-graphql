# godog-graphql
Feature context for godog with GraphQL steps

## Running using docker

```
docker run --rm --network="host" -v "${PWD}/features:/godog/features" -e GRAPHQL_URL=http://localhost:3000/graphql jakubknejzlik/godog-graphql
```