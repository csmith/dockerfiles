# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.1.0","image:base":"5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38","image:golang":"256f4fbf1a17959c3fddffff306e8dc39900bb07530036adf60e6c81c75eefc3"}

FROM reg.c5h.io/golang@sha256:256f4fbf1a17959c3fddffff306e8dc39900bb07530036adf60e6c81c75eefc3 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/csmith/irc-distribution /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38

COPY --from=build /work/out /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
