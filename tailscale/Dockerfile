# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.92.0","image:base":"46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7","image:golang":"5f6a4c6c2927f0ab0c0121807a47294754f9d9c4e11c2c0b40a636b8cef50350"}

FROM reg.c5h.io/golang@sha256:5f6a4c6c2927f0ab0c0121807a47294754f9d9c4e11c2c0b40a636b8cef50350 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.92.0 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.92.0"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
