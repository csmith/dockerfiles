# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.5","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"ada7e34fe68a90eb6495bb95c3e19297ee91630356b1d4c50b496db8c3f9ad26"}

FROM reg.c5h.io/golang@sha256:ada7e34fe68a90eb6495bb95c3e19297ee91630356b1d4c50b496db8c3f9ad26 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.5 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
