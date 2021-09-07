# Generated from https://github.com/csmith/dockerfiles/blob/master/arch/Dockerfile.gotpl
# BOM: {"pacman:acl":"2.3.1-1","pacman:archlinux-keyring":"20210902-1","pacman:attr":"2.5.1-1","pacman:audit":"3.0.4-1","pacman:bash":"5.1.008-1","pacman:bzip2":"1.0.8-4","pacman:ca-certificates":"20210603-1","pacman:ca-certificates-mozilla":"3.69.1-1","pacman:ca-certificates-utils":"20210603-1","pacman:coreutils":"8.32-1","pacman:curl":"7.78.0-1","pacman:e2fsprogs":"1.46.4-1","pacman:expat":"2.4.1-1","pacman:filesystem":"2021.05.31-1","pacman:findutils":"4.8.0-1","pacman:gcc-libs":"11.1.0-1","pacman:glib2":"2.68.4-1","pacman:glibc":"2.33-5","pacman:gmp":"6.2.1-1","pacman:gnupg":"2.2.29-1","pacman:gnutls":"3.7.2-2","pacman:gpgme":"1.16.0-1","pacman:gzip":"1.11-1","pacman:iana-etc":"20210728-1","pacman:keyutils":"1.6.3-1","pacman:krb5":"1.19.1-1","pacman:less":"1:590-1","pacman:libarchive":"3.5.2-1","pacman:libassuan":"2.5.5-1","pacman:libcap":"2.53-1","pacman:libcap-ng":"0.8.2-3","pacman:libffi":"3.3-4","pacman:libgcrypt":"1.9.4-1","pacman:libgpg-error":"1.42-1","pacman:libidn2":"2.3.2-1","pacman:libksba":"1.6.0-1","pacman:libldap":"2.4.59-2","pacman:libnghttp2":"1.44.0-1","pacman:libp11-kit":"0.24.0-1","pacman:libpsl":"0.21.1-1","pacman:libsasl":"2.1.27-3","pacman:libsecret":"0.20.4-1","pacman:libssh2":"1.9.0-3","pacman:libtasn1":"4.17.0-1","pacman:libtirpc":"1.3.2-1","pacman:libunistring":"0.9.10-3","pacman:libxcrypt":"4.4.25-1","pacman:licenses":"20200427-1","pacman:linux-api-headers":"5.12.3-1","pacman:lz4":"1:1.9.3-2","pacman:ncurses":"6.2-2","pacman:nettle":"3.7.3-1","pacman:npth":"1.6-3","pacman:openssl":"1.1.1.l-1","pacman:p11-kit":"0.24.0-1","pacman:pacman":"6.0.1-1","pacman:pacman-mirrorlist":"20210822-1","pacman:pam":"1.5.1-1","pacman:pambase":"20210605-2","pacman:pcre":"8.45-1","pacman:pcre2":"10.37-1","pacman:pinentry":"1.1.1-1","pacman:readline":"8.1.001-1","pacman:sed":"4.8-1","pacman:sqlite":"3.36.0-1","pacman:systemd-libs":"249.4-1","pacman:tzdata":"2021a-2","pacman:util-linux-libs":"2.37.2-1","pacman:xz":"5.2.5-1","pacman:zlib":"1:1.2.11-4","pacman:zstd":"1.5.0-1"}

FROM reg.c5h.io/arch AS build

RUN set -eux; \
    # Install packages we need to bootstrap an image
    pacman --noconfirm --needed -Syu extra/devtools extra/fakechroot core/fakeroot; \
    # Bodge the hooks (?)
    mkdir -vp /build/alpm-hooks/usr/share/libalpm/hooks; \
	find /usr/share/libalpm/hooks -exec ln -sf /dev/null /build/alpm-hooks{} \;; \
	# Create skeleton for pacman
    mkdir -vp /build/var/lib/pacman /build/etc/pacman.d; \
	install -Dm644 /usr/share/devtools/pacman-extra.conf /build/etc/pacman.conf; \
	# Configure pacman to avoid extracting stuff we don't care about
	echo -e '[options]\nNoExtract = \
	    usr/lib/systemd/** \
	    usr/share/audit/sample-rules/** \
	    usr/share/bash-completion/** \
	    usr/share/fish/vendor_completions.d/** \
	    usr/share/gnupg/help*.txt \
	    usr/share/pixmaps/** \
	    usr/share/terminfo/** \
	    usr/share/help/* \
        usr/share/iana-etc/* \
	    usr/share/gtk-doc/html/* \
	    usr/share/doc/* \
	    usr/share/locale/* \
	    !usr/share/locale/locale.alias \
	    usr/share/X11/locale/* \
	    usr/share/i18n/* \
	    !*locale*/en_US/* \
	    !usr/share/i18n/charmaps/UTF-8.gz \
	    !usr/share/*locales/en_GB \
	    !usr/share/*locales/en_US \
	    !usr/share/*locales/i18n \
	    !usr/share/*locales/i18n_ctype \
	    !usr/share/*locales/iso14651_* \
	    !usr/share/*locales/locale.alias \
	    !usr/share/*locales/translit_* \
	    usr/share/man/* \
	    usr/share/info/* \
	    usr/share/vim/vim*/lang/*' >> /build/etc/pacman.conf; \
    # Sync packages
    fakechroot -- fakeroot -- pacman -Sy -r /build \
        --noconfirm --dbpath /build/var/lib/pacman \
        --config /build/etc/pacman.conf \
        --noscriptlet \
        --hookdir /build/alpm-hooks/usr/share/libalpm/hooks/ \
        acl=2.3.1-1 \
        archlinux-keyring=20210902-1 \
        attr=2.5.1-1 \
        audit=3.0.4-1 \
        bash=5.1.008-1 \
        bzip2=1.0.8-4 \
        ca-certificates=20210603-1 \
        ca-certificates-mozilla=3.69.1-1 \
        ca-certificates-utils=20210603-1 \
        coreutils=8.32-1 \
        curl=7.78.0-1 \
        e2fsprogs=1.46.4-1 \
        expat=2.4.1-1 \
        filesystem=2021.05.31-1 \
        findutils=4.8.0-1 \
        gcc-libs=11.1.0-1 \
        glib2=2.68.4-1 \
        glibc=2.33-5 \
        gmp=6.2.1-1 \
        gnupg=2.2.29-1 \
        gnutls=3.7.2-2 \
        gpgme=1.16.0-1 \
        gzip=1.11-1 \
        iana-etc=20210728-1 \
        keyutils=1.6.3-1 \
        krb5=1.19.1-1 \
        less=1:590-1 \
        libarchive=3.5.2-1 \
        libassuan=2.5.5-1 \
        libcap=2.53-1 \
        libcap-ng=0.8.2-3 \
        libffi=3.3-4 \
        libgcrypt=1.9.4-1 \
        libgpg-error=1.42-1 \
        libidn2=2.3.2-1 \
        libksba=1.6.0-1 \
        libldap=2.4.59-2 \
        libnghttp2=1.44.0-1 \
        libp11-kit=0.24.0-1 \
        libpsl=0.21.1-1 \
        libsasl=2.1.27-3 \
        libsecret=0.20.4-1 \
        libssh2=1.9.0-3 \
        libtasn1=4.17.0-1 \
        libtirpc=1.3.2-1 \
        libunistring=0.9.10-3 \
        libxcrypt=4.4.25-1 \
        licenses=20200427-1 \
        linux-api-headers=5.12.3-1 \
        lz4=1:1.9.3-2 \
        ncurses=6.2-2 \
        nettle=3.7.3-1 \
        npth=1.6-3 \
        openssl=1.1.1.l-1 \
        p11-kit=0.24.0-1 \
        pacman=6.0.1-1 \
        pacman-mirrorlist=20210822-1 \
        pam=1.5.1-1 \
        pambase=20210605-2 \
        pcre=8.45-1 \
        pcre2=10.37-1 \
        pinentry=1.1.1-1 \
        readline=8.1.001-1 \
        sed=4.8-1 \
        sqlite=3.36.0-1 \
        systemd-libs=249.4-1 \
        tzdata=2021a-2 \
        util-linux-libs=2.37.2-1 \
        xz=5.2.5-1 \
        zlib=1:1.2.11-4 \
        zstd=1.5.0-1 \
        ; \
    # Drop in some config files
    echo 'en_US.UTF-8 UTF-8' > /build/etc/locale.gen; \
    echo 'LANG=en_US.UTF-8' > /build/etc/locale.conf; \
    echo 'Server = https://mirrors.melbourne.co.uk/archlinux/$repo/os/$arch' > /build/etc/pacman.d/mirrorlist; \
    # Update certificates, locales, etc
    fakechroot -- fakeroot -- chroot /build update-ca-trust; \
	fakechroot -- fakeroot -- chroot /build locale-gen; \
	fakechroot -- fakeroot -- chroot /build pacman-key --init; \
	fakechroot -- fakeroot -- chroot /build pacman-key --populate; \
	# Remove passwordless login for root (see CVE-2019-5021 for reference)
	sed -i -e 's/^root::/root:!:/' /build/etc/shadow; \
	# Remove a bunch of things we don't want to keep in the image
	rm -rf \
	    /build/alpm-hooks \
	    /build/dev \
	    /build/etc/pacman.d/gnupg/openpgp-revocs.d/* \
	    /build/etc/pacman.d/gnupg/private-keys-v1.d/* \
	    /build/etc/pacman.d/gnupg/pubring.gpg~ \
	    /build/etc/pacman.d/gnupg/S.* \
	    /build/etc/resolv.conf \
	    /build/proc \
	    /build/sys \
	    /build/var/lib/pacman/sync \
	    /build/var/log/*; \
	# Clobber all the timestamps to make the build more reproducible
    find /build -exec touch --date=@0 {} \;;

FROM scratch
COPY --from=build /build/ /
CMD ["/bin/bash"]
