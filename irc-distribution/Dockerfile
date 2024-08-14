# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff","image:golang":"73e57c51a73ef102f8d9ce0034cfe2bdfd1ab577ec2cce5335a6a8163486bc4a"}

FROM reg.c5h.io/golang@sha256:73e57c51a73ef102f8d9ce0034cfe2bdfd1ab577ec2cce5335a6a8163486bc4a AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
