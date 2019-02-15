#!/bin/bash

echo '1. 检测postgres状态'
echo "localhost:*:*:ngtesting:P2ssw0rd" > $HOME/.pgpass
echo "`chmod 0600 $HOME/.pgpass`"
count=`psql -h localhost -U ngtesting ngtesting-web  -c "select tablename from pg_tables where schemaname='public'" | grep TstProject`
echo "count=$count"

set -e  # Exit the script if an error happens

if [ "$count" != "" ]; then
    echo "数据已初始化，退出"
else
    echo '2. 创建用户....'
    su - postgres -c "psql  -c \"CREATE USER ngtesting WITH PASSWORD 'P2ssw0rd'\""
    echo '2. 成功创建用户'

    echo '3. 创建数据库....'
    su - postgres -c "psql  -c 'CREATE DATABASE \"ngtesting-web\" OWNER ngtesting'"
    echo '3. 成功创建数据库'

    echo '4. 导入数据....'
    su - postgres -c "psql ngtesting-web -U ngtesting < /schema.sql"
    echo '4. 成功导入数据'

    touch ~/init-success
fi
