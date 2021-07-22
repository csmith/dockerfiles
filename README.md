# Dockerfiles from the ground up

## What? Why?

This is a collection of Dockerfiles for various software projects I want to run, built from the ground up.

Most projects have either official docker images or third-party contributions, but they're always a bit
hit-or-miss on how they work, what base images are used, etc. Third-party images often lag behind releases
or just die off without warning, too. Building everything out from scratch ensures the images are standard,
and can follow upstream updates as quickly or slowly as required.

I'm intending to use these for production services, but won't vouch for their stability or usability for
anyone else's purposes. Feel free to use them, and report any issues you do find, but at your own risk!

## Structure

Each image has its own folder, containing at least:

* `Dockerfile` - the actual dockerfile the latest version should be built from
* `Dockerfile.gotpl` - a template file for generating the Dockerfile
* `update.go` - a tiny go application to grab the latest upstream version and render the template

My eventual plan is to have CI infrastructure to run the update checks periodically and commit the
updated dockerfiles back to the repository, triggering a build and push of the image.

## Images

| Name              | Description                                             |
|-------------------|---------------------------------------------------------|
| alpine            | Latest stable version of Alpine Linux                   |
| base              | Minimal "distroless" base image                         |
| distribution      | Docker registry                                         |
| golang            | Latest stable golang in Alpine, with musl/etc for cgo   |
