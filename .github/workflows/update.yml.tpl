on:
  workflow_dispatch:
  schedule:
    - cron: '33 3 * * *'
  push:
    paths-ignore:
      - '**/Containerfile'
      - '**/Dockerfile'
      - '.github/**'
      - 'README.md'
name: update
concurrency: dockerfiles
jobs:
{% for target in targets %}
  {{ target.name }}:
    name: Build {{ target.name }}
    runs-on: ubuntu-latest
{% if target.needed.size > 0 %}
    needs:
{%- for dep in target.needed -%}
        - {{ dep }}
{%- endfor %}
{% endif %}
{%- raw -%}
    steps:
      - name: Checkout source
        uses: actions/checkout@master
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: '1.19'
      - name: Update
        env:
          REGISTRY: ${{ secrets.REGISTRY }}
          REGISTRY_USER: ${{ secrets.REGISTRY_USER }}
          REGISTRY_PASS: ${{ secrets.REGISTRY_PASS }}
          SOURCE_LINK: https://github.com/csmith/dockerfiles/blob/master/
        run:  |
          go install github.com/csmith/contempt/cmd/contempt@latest
          git config user.name "${{ secrets.GIT_USERNAME }}"
          git config user.email "${{ secrets.GIT_EMAIL }}"
          buildah login -u "$REGISTRY_USER" -p "$REGISTRY_PASS" $REGISTRY
{%- endraw %}
          contempt --commit --build --push --project {{ target.name }} . .
          git push
{% endfor %}