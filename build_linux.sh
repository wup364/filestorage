echo build namenode
cd ./namenode
go build -o ../bin/namenode

cd ../
echo build datanode 
cd ./datanode
go build -o ../bin/datanode

cd ../
echo build tools
cd ./tools/shelltool
go build -o ../../bin/shelltool

cd  ../uploadtool
go build -o ../../bin/uploadtool

echo build end
