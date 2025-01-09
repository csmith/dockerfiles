# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.78.3","image:base":"a326d0e1775b6525bf9968a89afbd9b957c15d92f33afc33bad4c11e38adb46b","image:golang":"8dec622c0522882adf3ed0d9e6bd7f0fe66ce0153729aec362f9a5d091e973e2"}

FROM reg.c5h.io/golang@sha256:8dec622c0522882adf3ed0d9e6bd7f0fe66ce0153729aec362f9a5d091e973e2 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.78.3 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.78.3"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:a326d0e1775b6525bf9968a89afbd9b957c15d92f33afc33bad4c11e38adb46b

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
