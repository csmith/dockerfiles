# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02","image:golang":"sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670"}

FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/irc-news/main;

FROM reg.c5h.io/base@sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
