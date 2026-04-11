# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66","image:golang":"e077549c644e3d9f04836265464d0b49f0eecef44b7a777274a66cda8830225b"}

FROM reg.c5h.io/golang@sha256:e077549c644e3d9f04836265464d0b49f0eecef44b7a777274a66cda8830225b AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
