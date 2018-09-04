#!/bin/bash

echo `service mysql status`

echo '1. 启动mysql....'
service mysql start
sleep 3
echo `service mysql status`

count=`mysql -ungtesting -pP2ssw0rd -e "show databases" | grep ngtesting`
echo "count=$count"

set -e

if [ "$count" != "" ]; then
    echo "数据已初始化，退出"
else
    echo '2. 开始导入数据....'
    mysql < /etc/mysql/database.sql
    mysql ngtesting-web < /etc/mysql/schema.sql
    echo '3. 成功导入数据....'

    sleep 3
    echo `service mysql status`

    echo '4. 开始修改密码....'
    mysql < /etc/mysql/privileges.sql
    echo '5. 修改密码完毕....'

    #sleep 3
    echo `service mysql status`
    echo 'mysql容器启动完毕'

    touch ~/init
fi

tail -f /dev/null
