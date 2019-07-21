# Database Setup

This project use a PostgreSQL with a GraphQL interface, using Hasura
framework. For that, two services must be started: one postgre database,
and a hasura graphql engine.

On this folder, I've separated files to setup the postgre database tables,
populating with graphql mutations. These files are served on a docker image
to be used on diferent services:

- **ddspog/hasura-graphql-engine**: this service uses `schema.up.sql` and
`metadata.up.sql` to start the database with the tables defined.
- **ddspog/graphql-populate**: this service runs a program that loads
GQL files with a query to check if the tables are online (`check-tables.gql`)
and a mutation to populate the database via upsert (`populate.gql`).
