# Dockerfiles from the ground up

## What? Why?

This is a collection of Dockerfiles for various software projects I want to run, built from the ground up.

Most projects have either official docker images or third-party contributions, but they're always a bit
hit-or-miss on how they work, what base images are used, etc. Third-party images often lag behind releases
or just die off without warning, too. Building everything out from scratch ensures the images are standard,
and can follow upstream updates as quickly or slowly as required.

I'm intending to use these for production services, but won't vouch for their stability or usability for
anyone else's purposes. Feel free to use them, and report any issues you do find, but at your own risk!

### Aims

#### Reproducibility

It should be possible to take a Dockerfile from this repository, build it on any machine, and get an image with the
same digest.  This requires using `buildah` to build instead of `docker` as it supports setting the timestamp
for layers.

As most of the images are based on Alpine they won't continue to be reproducible forever as Alpine do not keep
old packages indefinitely. While this might be nice, for now reproducibility during the lifetime of the image is
sufficient.

#### Minimal and non-root

Almost all the images here use a minimal base image, and only ship resources actually required to run. In particular,
there is generally no shell, no scripted entrypoints, no build tools, etc, shipped in the final images.

Where possible the images all run as a non-root user (uid/gid `65532`), with no direct way to gain root access within
the container.

#### Up-to-date and transparent

The images all track their upstream projects, and the Dockerfiles are automatically updated when new releases are
made. The update procedure is performed using GitHub actions, and the revised Dockerfiles are immediately committed
to the repository. This makes it easy to know what an image contains, and when and how it was built.

## Structure

Each image has its own folder, containing:

* `Dockerfile` - the actual dockerfile the latest version should be built from
* `Dockerfile.gotpl` - a template file for generating the Dockerfile

The templates are interpreted using the [contempt](https://github.com/csmith/contempt) tool, which handles generating
the Dockerfile from the template as well as optionally building and pushing the images and committing the results.

## Images

All images are available at `reg.c5h.io/<name>`. Only the latest tag is built.

Images with the `-certwrapper` suffix have [csmith/certwrapper](https://github.com/csmith/certwrapper)
included as the entrypoint, to facilitate obtaining SSL certificates. To keep image size down, certwrapper
is built with the `httpreq` build tag (so can only use DNS providers that support the HTTPREQ protocol).
Pull requests for more general builds are welcome, as these variations are obviously quite specific to
my setup.
