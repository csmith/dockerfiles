# Generated from https://github.com/csmith/dockerfiles/blob/master/tailscale/Dockerfile.gotpl
# BOM: {"github:tailscale/tailscale":"v1.96.5","image:base":"46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987","image:golang":"d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a"}

FROM reg.c5h.io/golang@sha256:d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a AS build

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

FROM reg.c5h.io/base@sha256:46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987

COPY --from=build /src/tailscaled /tailscaled
COPY --from=build /src/tailscale /tailscale
COPY --from=build /out /

VOLUME /state
ENV PATH=/

ENTRYPOINT ["/tailscaled", "--tun=userspace-networking", "--state=/state/tailscaled.state"]
