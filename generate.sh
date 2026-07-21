#!/bin/bash
echo "Compiling Go stubs..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    *.proto

echo "Compiling C# stubs..."
cd ../artemis_link
dotnet build
