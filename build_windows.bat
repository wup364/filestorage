echo off

echo build namenode
cd ./namenode
go build -o ../bin/namenode.exe

cd ../
echo build datanode 
cd ./datanode
go build -o ../bin/datanode.exe

cd ../
echo build tools
cd ./tools/shelltool
go build -o ../../bin/shelltool.exe

cd  ../uploadtool
go build -o ../../bin/uploadtool.exe

echo build end