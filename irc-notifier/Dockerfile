# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.5","image:base":"81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0","image:golang":"954b97f11d387f53f8aede05ec76d617607ceff094bfe9a84d19e208bec43f5e"}

FROM reg.c5h.io/golang@sha256:954b97f11d387f53f8aede05ec76d617607ceff094bfe9a84d19e208bec43f5e AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.5 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
