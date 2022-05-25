# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"93252001f0ac448e4117c7203496cd54f7e8a73d779b9aa8b8f285d6acf31811","image:golang":"a56c892f582c7e3f2de33296f300c2546d1ad6562f2fbff30cfc9a427412d876"}

FROM reg.c5h.io/golang@sha256:a56c892f582c7e3f2de33296f300c2546d1ad6562f2fbff30cfc9a427412d876 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:93252001f0ac448e4117c7203496cd54f7e8a73d779b9aa8b8f285d6acf31811

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
