# Generated from https://github.com/csmith/dockerfiles/blob/master/vault/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.33.1-r3","apk:ca-certificates":"20191127-r5","apk:libacl":"2.2.53-r0","apk:libcap":"2.50-r0","apk:libcrypto1.1":"1.1.1k-r0","apk:musl":"1.2.2-r3","apk:popt":"1.18-r0","apk:rsync":"3.2.3-r4","apk:tzdata":"2021a-r0","apk:zlib":"1.2.11-r3","apk:zstd-libs":"1.4.9-r1","image:reg.c5h.io/alpine":"sha256:1b32fef9ac1c1fce7d7fcc352d0408dc0e8a62a60d380572b8f2e638d435da5f","image:reg.c5h.io/golang":"sha256:3a3534d57448c2109cdc0e08d2fdb9bc759207508ca9af0b166584736d5e034a"}

FROM reg.c5h.io/golang@sha256:3a3534d57448c2109cdc0e08d2fdb9bc759207508ca9af0b166584736d5e034a AS build

ARG TAG="v1.8.0"

RUN set -eux; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/hashicorp/vault /src; \
    cd /src; \
    go build -tags release -ldflags "-X github.com/hashicorp/vault/sdk/version.GitCommit=$(git rev-parse HEAD)" -o /vault  github.com/hashicorp/vault; \
    # Make some directories with the correct ownership to copy into the final image, and a listener config snippet
    mkdir -p /out/logs /out/file /out/config; \
    echo -e 'listener "tcp" {\n  address = "0.0.0.0:8080"\n  tls_disable = 1\n}' > /out/config/listener.hcl; \
    chown -R 65532:65532 /out; \
    # Clobber all the timestamps to make the build more reproducible
    find /out -exec touch --date=@0 {} \;; \
    touch --date=@0 /vault;

FROM reg.c5h.io/alpine@sha256:1b32fef9ac1c1fce7d7fcc352d0408dc0e8a62a60d380572b8f2e638d435da5f AS base
# This is basically the same thing as reg.c5h.io/base, but we add libcap and don't drop root until after we've
# had a chance to call setcap on the binary.

RUN set -eux; \
    # Sort out our dependencies (and sources)
    apk add --no-cache \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        libacl=2.2.53-r0 \
        libcap=2.50-r0 \
        libcrypto1.1=1.1.1k-r0 \
        musl=1.2.2-r3 \
        popt=1.18-r0 \
        rsync=3.2.3-r4 \
        tzdata=2021a-r0 \
        zlib=1.2.11-r3 \
        zstd-libs=1.4.9-r1 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy package contents
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL libcap | rsync -aq --files-from=- / /rootfs/; \
    # Generate thw CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Clobber all the timestamps to make the build more reproducible
    find /rootfs -exec touch --date=@0 {} \;;

FROM scratch

COPY --from=base /rootfs/ /
COPY --from=build /vault /bin/
COPY --from=build /out/ /vault/

RUN ["/usr/sbin/setcap","CAP_IPC_LOCK=+ep","/bin/vault"]

ENV VAULT_ADDR=http://127.0.0.1:8080 \
    SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

USER nonroot
EXPOSE 8080
VOLUME /vault/logs /vault/file
ENTRYPOINT ["/bin/vault"]

CMD ["server", "-config=/vault/config/"]
