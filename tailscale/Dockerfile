# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.90.5","image:base":"a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024","image:golang":"bcb710a28359e1c23074f451cb23aec940a9a009f76675ba5e049c83acf3f50f"}

FROM reg.c5h.io/golang@sha256:bcb710a28359e1c23074f451cb23aec940a9a009f76675ba5e049c83acf3f50f AS build

RUN set -eux; \
    git clone --depth=1 -b v1.90.5 --single-branch https://github.com/tailscale/tailscale /src

WORKDIR /src


ARG TAG="v1.90.5"

RUN set -eux; \
    # Build
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscale; \
    CGO_ENABLED=0 GOOS=linux go build -tags=xversion -trimpath -ldflags "-extldflags \"-static\" -buildid= -X tailscale.com/version.Long=$(echo $TAG | tr -d 'v')-t$(git rev-parse HEAD | cut -c1-9) -X tailscale.com/version.Short=$(echo $TAG | tr -d 'v') -X tailscale.com/version.GitCommit=$(git rev-parse HEAD)" ./cmd/tailscaled; \
    mkdir -p /out/var/run/tailscale /out/state; \
    chown -R 65532:65532 /out;

FROM reg.c5h.io/base@sha256:a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
