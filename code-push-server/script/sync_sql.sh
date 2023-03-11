#!/bin/sh

initDir=/var/folders/mysql/init

if [ ! -d $initDir ]; then
    mkdir -p $initDir
else
    echo "Removing $initDir"
    rm -rf $initDir
    echo "Creating $initDir"
    mkdir -p $initDir
fi

echo "Copying sql files to $initDir"
cp -r ../mysql/init/* $initDir