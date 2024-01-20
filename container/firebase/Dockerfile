FROM openjdk:slim

RUN set -ex \
  && apt-get update \
  && apt-get install -y --no-install-recommends curl \
  && curl -sL https://deb.nodesource.com/setup_20.x | bash - \
  && apt-get install -y --no-install-recommends nodejs \
  && npm install -g firebase-tools@latest

RUN  apt-get clean && rm -rf /var/lib/apt/lists/*

EXPOSE 9099 9199 4000 4500
