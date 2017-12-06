#!/bin/sh
set -e

# Workaround old docker images with incorrect $HOME
# check https://github.com/docker/docker/issues/2968 for details
if [ "${HOME}" = "/" ]
then
  export HOME=$(getent passwd $(id -un) | cut -d: -f6)
fi

git config --global url."https://github.com".insteadOf "ssh://git@github.com" || true

if [ -e /home/user/work/src/github.com/Nordgedanken/Morpheus/.git ]
then
  cd /home/user/work/src/github.com/Nordgedanken/Morpheus/
  git remote set-url origin "$CIRCLE_REPOSITORY_URL" || true
else
  mkdir -p /home/user/work/src/github.com/Nordgedanken/Morpheus/
  cd /home/user/work/src/github.com/Nordgedanken/Morpheus/
  git clone "$CIRCLE_REPOSITORY_URL" .
fi

if [ -n "$CIRCLE_TAG" ]
then
  git fetch --force origin "refs/tags/${CIRCLE_TAG}"
else
  git fetch --force origin "${CIRCLE_BRANCH}:remotes/origin/${CIRCLE_BRANCH}"
fi


if [ -n "$CIRCLE_TAG" ]
then
  git reset --hard "$CIRCLE_SHA1"
  git checkout -q "$CIRCLE_TAG"
elif [ -n "$CIRCLE_BRANCH" ]
then
  git reset --hard "$CIRCLE_SHA1"
  git checkout -q -B "$CIRCLE_BRANCH"
fi

git reset --hard "$CIRCLE_SHA1"