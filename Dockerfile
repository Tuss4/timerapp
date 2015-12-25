FROM ubuntu:14.04
MAINTAINER Tomjo Soptame <tj.soptame@gmail.com>
RUN apt-get update && apt-get install -y sudo wget vim
RUN wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir /code
ENV GOPATH /code
WORKDIR /code
ADD . /code/
EXPOSE 5000
CMD ["go", "run", "main.go"]
