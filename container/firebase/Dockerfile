FROM openjdk:slim

RUN set -ex \
  && apt-get update \
  && apt-get install -y curl \
  && curl -sL https://deb.nodesource.com/setup_18.x | bash \
  && apt-get install -y nodejs \
  && npm install -g firebase-tools

EXPOSE 9099
EXPOSE 9199
EXPOSE 4000
EXPOSE 4500
