#!/bin/bash

count=`mysql -ungtesting -pP2ssw0rd -e "show databases" | grep ngtesting`
echo $count

if [ "$count" != "" ]; then
    echo "已初始化"
    exit
else
    echo "未初始化"
fi
