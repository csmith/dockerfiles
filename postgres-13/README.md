# Postgresql

* Upstream: https://www.postgresql.org/
* Dependencies: alpine

Provides the latest stable release of Postgres 13. Because upgrades between major postgres
versions require manual intervention, builds from v14 will be released under a different
image.

This is built using the Dockerfile and associated scripts from the
[official image](https://github.com/docker-library/postgres) used by DockerHub, as manually
compiling Postgres is an _involved_  process. It thus runs in a full alpine container instead
of a more minimal base image.
