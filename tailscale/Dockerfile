# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.100.0","image:base":"932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929","image:golang":"bb6a061c510ddf8aaf4d480c1eabe858d1a996bf52ba05108d4da3da16a91d68"}

FROM reg.c5h.io/golang@sha256:bb6a061c510ddf8aaf4d480c1eabe858d1a996bf52ba05108d4da3da16a91d68 AS build

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

FROM reg.c5h.io/base@sha256:932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
