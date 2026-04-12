# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.96.5","image:base":"22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66","image:golang":"5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736"}

FROM reg.c5h.io/golang@sha256:5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.96.5 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.96.5"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
