FROM ubuntu

RUN echo deb http://archive.ubuntu.com/ubuntu precise universe > /etc/apt/sources.list.d/universe.list
RUN apt-get update -qq


ADD ./qrserver /opt/qrserver

EXPORT 8080

CMD /opt/qrserver