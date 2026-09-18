# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-goplum":"v4.0.2","image:base":"68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc","image:golang":"424727a3913a09b7b5f3aa8b43cc60a6aded1a6b814d2ca889f7305c2a06c6a3"}

FROM reg.c5h.io/golang@sha256:424727a3913a09b7b5f3aa8b43cc60a6aded1a6b814d2ca889f7305c2a06c6a3 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
