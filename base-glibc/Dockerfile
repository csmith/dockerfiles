# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:acl-libs":"2.3.2-r1","apk:busybox":"1.37.0-r31","apk:busybox-binsh":"1.37.0-r31","apk:ca-certificates":"20260413-r0","apk:gcompat":"1.1.0-r4","apk:libcrypto3":"3.5.6-r0","apk:libgcc":"15.2.0-r5","apk:libstdc++":"15.2.0-r5","apk:libucontext":"1.5.1-r0","apk:libxxhash":"0.8.3-r1","apk:lz4-libs":"1.10.0-r1","apk:musl":"1.2.6-r2","apk:musl-obstack":"1.2.3-r2","apk:popt":"1.19-r4","apk:rsync":"3.4.3-r1","apk:tzdata":"2026b-r0","apk:zlib":"1.3.2-r0","apk:zstd-libs":"1.5.7-r2","image:alpine":"129ff9b715c38d5c8bb17ef16f2490821ab651ed31b3d024d81ad45a7ffbe516"}

FROM reg.c5h.io/alpine@sha256:129ff9b715c38d5c8bb17ef16f2490821ab651ed31b3d024d81ad45a7ffbe516 AS build

RUN set -eux; \
    apk add --no-cache \
        busybox=1.37.0-r31 \
        busybox-binsh=1.37.0-r31 \
        ca-certificates=20260413-r0 \
        libcrypto3=3.5.6-r0 \
        musl=1.2.6-r2 \
        acl-libs=2.3.2-r1 \
        libxxhash=0.8.3-r1 \
        lz4-libs=1.10.0-r1 \
        musl=1.2.6-r2 \
        popt=1.19-r4 \
        rsync=3.4.3-r1 \
        zlib=1.3.2-r0 \
        zstd-libs=1.5.7-r2 \
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
    apk add --no-cache libgcc=15.2.0-r5; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=15.2.0-r5; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libucontext=1.5.1-r0; \
    apk info -qL libucontext | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.6-r2; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl-obstack=1.2.3-r2; \
    apk info -qL musl-obstack | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2026b-r0; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
