#grcp

#安装Protocol Buffer v3

#检查安装
protoc --version

#生成go文件命令
protoc --go_out=plugins=grpc:. *.proto

#Server
##生成 Key
openssl genrsa -out ca.key 2048

##生成秘钥
openssl req -new -x509 -days 7200 -key ca.key -out ca.pem

##生成 CSR
openssl req -new -key server.key -out server.csr

##基于 CA 签发
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

#Client
##生成Key
openssl ecparam -genkey -name secp384r1 -out client.key

#生成 CSR
openssl req -new -key client.key -out client.csr

#基于 CA 签发
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem