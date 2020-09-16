FROM ubuntu:focal

WORKDIR /blockchain
EXPOSE 12345 1317 26657 8080 26656
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH ${GOPATH}/bin:${GOROOT}/bin:${PATH}

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y make gcc vim curl wget git net-tools iputils-ping nmap netcat && \
    wget https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz && \
    tar -xvf go1.15.1.linux-amd64.tar.gz && \
    mv go /usr/local && \
    touch ~/.bashrc && \
    echo ${GOROOT} && \
    echo ${GOPATH} && \
    echo ${PATH} && \
    # echo 'export GOROOT=/usr/local/go' > ~/.bashrc && \
    # echo 'export GOPATH=${HOME}/go' >> ~/.bashrc && \
    # echo 'export PATH=${GOPATH}/bin:${GOROOT}/bin:${PATH}' >> ~/.bashrc && \
    git clone https://github.com/aofiee/amulet.git && \
    cd amulet && \
    git checkout commit-transaction  && \
    make install && \
    rm -r /var/lib/apt/lists/*
