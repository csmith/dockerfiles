# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.6","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"04c2815074cae37712e2ccaa1bc24d08b3ef27776d72b05986a503413c3d1764"}

FROM reg.c5h.io/golang@sha256:04c2815074cae37712e2ccaa1bc24d08b3ef27776d72b05986a503413c3d1764 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.6 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
