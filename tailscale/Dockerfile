# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.16.1","image:base":"e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022","image:golang":"c8273d3333b2ccddb6fcea5c25e93d7af5b03009dfebee7635bd94b8f9f26b00"}

FROM reg.c5h.io/golang@sha256:c8273d3333b2ccddb6fcea5c25e93d7af5b03009dfebee7635bd94b8f9f26b00 AS build

ARG TAG="v1.16.1"

RUN set -eux; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/tailscale/tailscale /src; \
    cd /src; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /src/tailscale{,d} /out;

FROM reg.c5h.io/base@sha256:e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
