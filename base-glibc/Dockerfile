# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:acl-libs":"2.3.2-r1","apk:busybox":"1.37.0-r18","apk:busybox-binsh":"1.37.0-r18","apk:ca-certificates":"20250619-r0","apk:gcompat":"1.1.0-r4","apk:libcrypto3":"3.5.1-r0","apk:libgcc":"14.2.0-r6","apk:libstdc++":"14.2.0-r6","apk:libucontext":"1.3.2-r0","apk:libxxhash":"0.8.3-r0","apk:lz4-libs":"1.10.0-r0","apk:musl":"1.2.5-r10","apk:musl-obstack":"1.2.3-r2","apk:popt":"1.19-r4","apk:rsync":"3.4.1-r0","apk:tzdata":"2025b-r0","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.7-r0","image:alpine":"660c3cb1032bcacc15b95b8be733562d35b7c6b66be87fc02e55cfe5bcce3bd8"}

FROM reg.c5h.io/alpine@sha256:660c3cb1032bcacc15b95b8be733562d35b7c6b66be87fc02e55cfe5bcce3bd8 AS build

RUN set -eux; \
    apk add --no-cache \
        busybox=1.37.0-r18 \
        busybox-binsh=1.37.0-r18 \
        ca-certificates=20250619-r0 \
        libcrypto3=3.5.1-r0 \
        musl=1.2.5-r10 \
        acl-libs=2.3.2-r1 \
        libxxhash=0.8.3-r0 \
        lz4-libs=1.10.0-r0 \
        musl=1.2.5-r10 \
        popt=1.19-r4 \
        rsync=3.4.1-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
    ;

RUN set -eux; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache gcompat=1.1.0-r4; \
    apk info -qL gcompat | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libgcc=14.2.0-r6; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=14.2.0-r6; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libucontext=1.3.2-r0; \
    apk info -qL libucontext | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.5-r10; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl-obstack=1.2.3-r2; \
    apk info -qL musl-obstack | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2025b-r0; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
