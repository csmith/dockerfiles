# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:https://github.com/csmith/irc-distribution":"v1.1.0","image:base":"fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11","image:golang":"0accd527554df271b35255509308ca8723abe68eb6c9776436ab33fb939a0329"}

FROM reg.c5h.io/golang@sha256:0accd527554df271b35255509308ca8723abe68eb6c9776436ab33fb939a0329 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/csmith/irc-distribution /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11

COPY --from=build /work/out /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
