FROM centos:centos8

WORKDIR /filestorage/

# RUN yum install golang

COPY ./datanode /filestorage/datanode

# CMD [  ]

ENTRYPOINT  ["./datanode"]

VOLUME ["/filestorage/.datas", "/filestorage/.conf", "/filestorage/datas"]
