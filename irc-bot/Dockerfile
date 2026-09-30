# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-bot":"v5.0.10","image:base":"68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc","image:golang":"034439cb9fb2e8d3d2f43c37f0755580b215e299bcb3545dd3fcbef624ebd653"}

FROM reg.c5h.io/golang@sha256:034439cb9fb2e8d3d2f43c37f0755580b215e299bcb3545dd3fcbef624ebd653 AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.10 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
