#!/bin/bash

GIT_USER_ID="fnt-eve"
GIT_REPO_ID="esi/esiclient"
openapi-generator-cli generate \
  --git-user-id $GIT_USER_ID \
  --git-repo-id $GIT_REPO_ID \
  -g go \
  -o esiclient \
  -i https://esi.evetech.net/meta/openapi-3.0.json \
  -p packageName=esiclient \
  -p withGoMod=false


gofmt -s -w .
goimports -w .