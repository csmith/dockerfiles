# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.78.3","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"f11bab806d7b9d164d3d52afa3ac5c0f976cb5b78a09c1986ed6d25c971c2a45"}

FROM reg.c5h.io/golang@sha256:f11bab806d7b9d164d3d52afa3ac5c0f976cb5b78a09c1986ed6d25c971c2a45 AS build

ARG TAG="v1.78.3"

RUN set -eux; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/tailscale/tailscale /src; \
    cd /src; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
