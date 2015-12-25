FROM ubuntu:14.04
MAINTAINER Tomjo Soptame <tuss4dbfn@gmail.com>
RUN apt-get update && apt-get install -y sudo wget vim git
RUN wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir /godir
ENV GOPATH /godir
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p $GOPATH/src/code
WORKDIR $GOPATH/src/code
ADD . $GOPATH/src/code/
EXPOSE 5000
RUN . $GOPATH/src/code/dependencies.sh
RUN go install
CMD ["code"]
