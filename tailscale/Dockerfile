# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.100.0","image:base":"5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38","image:golang":"256f4fbf1a17959c3fddffff306e8dc39900bb07530036adf60e6c81c75eefc3"}

FROM reg.c5h.io/golang@sha256:256f4fbf1a17959c3fddffff306e8dc39900bb07530036adf60e6c81c75eefc3 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.100.0 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.100.0"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
