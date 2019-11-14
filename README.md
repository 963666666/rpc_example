#grcp

#安装Protocol Buffer v3

#检查安装
protoc --version

#生成go文件命令
protoc --go_out=plugins=grpc:. *.proto

