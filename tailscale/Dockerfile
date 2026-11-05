# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:https://github.com/tailscale/tailscale":"v1.104.5","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"4a8d5555821344f4f81945c99f58174d10913f63c25568befd2c2b58b43f237b"}

FROM reg.c5h.io/golang@sha256:4a8d5555821344f4f81945c99f58174d10913f63c25568befd2c2b58b43f237b AS build

RUN set -eux; \
    git clone --depth=1 -b v1.104.5 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.104.5"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
