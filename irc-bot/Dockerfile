# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-bot":"v5.0.10","image:base":"164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1","image:golang":"0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3"}

FROM reg.c5h.io/golang@sha256:0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3 AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.10 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
