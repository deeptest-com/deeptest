#!/bin/bash  
  
# 指定要修改的文件路径  
file_path="/workspace/leyanapi-backend/internal/pkg/consts/consts.go"  
  
# 使用sed命令替换文件中的内容  
sed -i 's/ConfigFileName = "server.yaml"/ConfigFileName = "config\/configMap\/server.yaml"/g' "$file_path"