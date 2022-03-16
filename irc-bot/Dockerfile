# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.5","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968","image:golang":"ad20796ad97b44998b5a22b02f0b65ad8f650f12b403b96d02dad3dbfff44ceb"}

FROM reg.c5h.io/golang@sha256:ad20796ad97b44998b5a22b02f0b65ad8f650f12b403b96d02dad3dbfff44ceb AS build

ARG TAG="v5.0.5"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
