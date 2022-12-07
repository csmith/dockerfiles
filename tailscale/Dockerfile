# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.34.0","image:base":"d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743","image:golang":"06bd9df2c9f04a24cf605e73052fa1983be5a1d525dcbecdfee4a0c407d318f5"}

FROM reg.c5h.io/golang@sha256:06bd9df2c9f04a24cf605e73052fa1983be5a1d525dcbecdfee4a0c407d318f5 AS build

ARG TAG="v1.34.0"

RUN set -eux; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/tailscale/tailscale /src; \
    cd /src; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
