# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:ca-certificates":"20220614-r4","apk:libacl":"2.3.1-r1","apk:libc6-compat":"1.2.3-r4","apk:libcrypto3":"3.0.8-r0","apk:libgcc":"12.2.1_git20220924-r4","apk:libstdc++":"12.2.1_git20220924-r4","apk:lz4-libs":"1.9.4-r1","apk:musl":"1.2.3-r4","apk:popt":"1.19-r0","apk:rsync":"3.2.7-r0","apk:tzdata":"2022f-r1","apk:zlib":"1.2.13-r0","apk:zstd-libs":"1.5.2-r9","image:alpine":"3d271830a936728b41e38fcf89a9703ee4a9983af40618722eefead94279b55c"}

FROM reg.c5h.io/alpine@sha256:3d271830a936728b41e38fcf89a9703ee4a9983af40618722eefead94279b55c AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        busybox=1.35.0-r29 \
        busybox-binsh=1.35.0-r29 \
        ca-certificates=20220614-r4 \
        libacl=2.3.1-r1 \
        libcrypto3=3.0.8-r0 \
        lz4-libs=1.9.4-r1 \
        musl=1.2.3-r4 \
        popt=1.19-r0 \
        rsync=3.2.7-r0 \
        zlib=1.2.13-r0 \
        zstd-libs=1.5.2-r9 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache libc6-compat=1.2.3-r4; \
    apk info -qL libc6-compat | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libgcc=12.2.1_git20220924-r4; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=12.2.1_git20220924-r4; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.3-r4; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2022f-r1; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
