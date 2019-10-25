# knowstellation

## Client Design

```mermaid
graph TD

Client --- cmd

subgraph knowstellation
cmd --- KnowstellationClient

end

subgraph knowstellationd
KnowstellationClient --- KnowstellationAPI
end
```

## Server Design

```mermaid
graph TB

knowstellationClient[Knowstellation Client]
knowstellationClient --request--> api

subgraph knowstellationAPI
api --- user
api --- star
end
```

