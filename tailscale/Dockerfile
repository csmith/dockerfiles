# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:https://github.com/tailscale/tailscale":"v1.102.3","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"030fd08673eb86719f173163c7b6cf26185177f79939c7485e11611f155ba4df"}

FROM reg.c5h.io/golang@sha256:030fd08673eb86719f173163c7b6cf26185177f79939c7485e11611f155ba4df AS build

RUN set -eux; \
    git clone --depth=1 -b v1.102.3 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.102.3"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
