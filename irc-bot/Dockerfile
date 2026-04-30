# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958","image:golang":"3e1c17a1ec056ba6d63607916f131e01d0b1ccdf052536866f1c15bc8567c8ba"}

FROM reg.c5h.io/golang@sha256:3e1c17a1ec056ba6d63607916f131e01d0b1ccdf052536866f1c15bc8567c8ba AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
