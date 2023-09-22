
## 打包时，关于环境的注意事项

1. node版本请使用 v16.x.x

2. 使用 npm 安装依赖，否则会出现依赖版本不一致的问题

3. 提示electron 安装慢，可以.npmrc修改镜像源， 或者直接执行一下命令

```bash
ELECTRON_MIRROR=https://npmmirror.com/mirrors/electron/ npm  install --registry=https://registry.npmmirror.com

sudo ELECTRON_MIRROR=https://npmmirror.com/mirrors/electron/ npm install electron -g --unsafe-perm=true --allow-root
```

4. 关于如何设置客户端应用的图标，参考
   https://www.electronforge.io/guides/create-and-add-icons

5. 不同客户端对应的格式

| Platform | Icon Format | Size           |
|----------|-------------|----------------|
| macOS    | .icns       | 512x512 pixels  (1024x1024 for Retina displays) |
| Windows  | .ico        | 256x256 pixels               |
| Linux    | .png        |      512x512 pixels          |


6. 报错如下，解决方案

```bash
An unhandled rejection has occurred inside Forge:
Error: EACCES: permission denied, rmdir '/tmp/electron-packager'

手动删除 /tmp/electron-packager 目录即可
```



## 具体的打包打包步骤如下

1. 修改版本号
   1. Makefile 文件中的 VERSION 变量 
   2. client/package.json 中的 version 字段
   3. package-ly.json 中的 version 字段 

2. 执行打包命令

```bash
sudo make ly-mac # mac
sudo make ly-win # windows
```      



