# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce","image:golang":"dabaf8f643cc9e4176cae12dce4e174222a7285410bf0df42baee2324fd10f82"}

FROM reg.c5h.io/golang@sha256:dabaf8f643cc9e4176cae12dce4e174222a7285410bf0df42baee2324fd10f82 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
