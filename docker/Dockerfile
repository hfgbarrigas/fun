FROM debian:stable-20210621-slim

ARG VERSION=0.18.1
ARG LITECOIN_PUB_KEY=FE3348877809386C
ARG LITECOIN_DATA_PATH=/home/litecoin/.litecoin

ENV DATA_PATH=$LITECOIN_DATA_PATH
WORKDIR /tmp

# update and install dependencies
RUN apt-get update \
    && apt-get install -y curl gnupg \
    && apt-get clean

# set permissions to: o+og - full , u - read and execute
RUN chmod -R 775 /usr/local/bin

# import the litecoin public key to gpg
RUN gpg --keyserver hkp://keyserver.ubuntu.com --recv-keys ${LITECOIN_PUB_KEY}

# check the key fingerprint
RUN gpg --fingerprint ${LITECOIN_PUB_KEY}

# download litecoin.tar.gz and its sha256 hashes file
RUN curl -SLO https://download.litecoin.org/litecoin-${VERSION}/linux/litecoin-${VERSION}-x86_64-linux-gnu.tar.gz \
  && curl -SLO https://download.litecoin.org/litecoin-${VERSION}/linux/litecoin-${VERSION}-linux-signatures.asc

# verify the sha256 hashes file signature using the imported public key
RUN gpg --verify litecoin-${VERSION}-linux-signatures.asc

# ensure that our computed sha256 hashs of the downloaded .tar.gz matches one within the sha256 hashes file
RUN grep $(sha256sum litecoin-${VERSION}-x86_64-linux-gnu.tar.gz | awk '{ print $1 }') litecoin-${VERSION}-linux-signatures.asc

# decompress and clean up
RUN mkdir decompressed \
    && tar -xzf *.tar.gz --strip=2 -C decompressed \
    && mv decompressed/* /usr/local/bin

# merely informative
EXPOSE 9332 9333 19332 19335 19444

# create the appropriate user
RUN useradd -m litecoin
USER litecoin
WORKDIR /home/litecoin

RUN mkdir $DATA_PATH
COPY run.sh run.sh

ENTRYPOINT ["./run.sh"]
