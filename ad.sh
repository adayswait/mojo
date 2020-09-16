#!/bin/bash
cd `dirname "$0"`
cd ./svr
./do.sh fmt
./do.sh build
mv ./mojo ./../
cd ..
cd ./cli
npm install --unsafe-perm
npm run build
cp -r ./dist ./../
cd ..
rm -rf `ls|egrep -v "(mojo|dist|br.sh|ar.sh)"`
