# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.4","image:base":"sha256:bba8c74ee389cc37da171979067085b53a8561caa66f49719db79e2fbb5d96cd","image:golang":"sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670"}

FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v5.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-bot/main;

FROM reg.c5h.io/base@sha256:bba8c74ee389cc37da171979067085b53a8561caa66f49719db79e2fbb5d96cd

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
