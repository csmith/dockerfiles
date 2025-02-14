# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.80.2","image:base":"524e2a72907aa209ab61ecec7762951bd727b2c31592c5d3c1b2acce5f33c83f","image:golang":"b720b8da0a766dcdcaa91f3d90527b7c53789bdab61388e24617c0359845e1ee"}

FROM reg.c5h.io/golang@sha256:b720b8da0a766dcdcaa91f3d90527b7c53789bdab61388e24617c0359845e1ee AS build

RUN set -eux; \
    git clone --depth=1 -b v1.80.2 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.80.2"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:524e2a72907aa209ab61ecec7762951bd727b2c31592c5d3c1b2acce5f33c83f

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
