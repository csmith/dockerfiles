# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff","image:golang":"a65c74423f2c96b1a1d676105009332bb38adf96f8aed5e1e51e2bc6d9def27d"}

FROM reg.c5h.io/golang@sha256:a65c74423f2c96b1a1d676105009332bb38adf96f8aed5e1e51e2bc6d9def27d AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
