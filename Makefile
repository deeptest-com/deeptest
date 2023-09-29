#VERSION=1.1.1
#PROJECT=deeptest

# ly 打包配置，开源版可以删除

VERSION=0.0.15
PROJECT=LeyanAPI


ifeq ($(OS),Windows_NT)
    PLATFORM="windows"
else
    ifeq ($(shell uname),Darwin)
        PLATFORM="mac"
    else
        PLATFORM="unix"
    endif
endif

ifeq ($(PLATFORM),"mac")
	QINIU_DIR=/Users/aaron/work/qiniu/
else
    QINIU_DIR=~/work/qiniu/
endif


# ly打包的路径,单独设置
PKG_W64=npm run package-win64
PKG_MAC=npm run package-mac
PKG_W32=npm run package-win32
PKG_LINUX=npm run package-linux
ifeq ($(PROJECT),LeyanAPI)
	QINIU_DIR=~/nk2/ly/
	PKG_W64=npm run ly-package-win64
	PKG_MAC=npm run ly-package-mac
	PKG_W32=npm run ly-package-win32
	PKG_LINUX=npm run ly-package-linux
endif


QINIU_DIST_DIR=${QINIU_DIR}${PROJECT}/${VERSION}/

SERVER_MAIN_FILE=cmd/server/main.go
AGENT_MAIN_FILE=cmd/agent/main.go

BIN_DIR=bin/
CLIENT_UI_DIR=client/ui/
CLIENT_BIN_DIR=client/bin/
CLIENT_OUT_DIR=client/out/
CLIENT_OUT_DIR_EXECUTABLE=${CLIENT_OUT_DIR}executable/
CLIENT_OUT_DIR_UPGRADE=${CLIENT_OUT_DIR}upgrade/

BUILD_TIME=`git show -s --format=%cd`
GO_VERSION=`go version`
GIT_HASH=`git show -s --format=%H`
BUILD_CMD_UNIX=go build -ldflags "-X 'main.AppVersion=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GoVersion=${GO_VERSION}' -X 'main.GitHash=${GIT_HASH}'"
BUILD_CMD_WIN=go build -ldflags "-s -w -X 'main.AppVersion=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GoVersion=${GO_VERSION}' -X 'main.GitHash=${GIT_HASH}'"

default: win64 win32 linux mac

win64: prepare build_gui_win64 compile_launcher_win64 compile_server_win64 copy_files_win64 zip_win64 zip_win64_upgrade
win32: prepare build_gui_win32 compile_launcher_win32 compile_server_win32 copy_files_win32 zip_win32 zip_win32_upgrade
linux: prepare build_gui_linux                        compile_server_linux copy_files_linux zip_linux zip_linux_upgrade
mac:   prepare build_gui_mac                          compile_server_mac   copy_files_mac   zip_mac zip_mac_upgrade

# 乐研 打包
ly-win64: prepare compile_ly_ui_client build_gui_win64 compile_launcher_win64 compile_server_win64 copy_files_win64 zip_win64 zip_win64_upgrade
ly-win32: prepare compile_ly_ui_client build_gui_win32 compile_launcher_win32 compile_server_win32 copy_files_win32 zip_win32 zip_win32_upgrade
ly-linux: prepare compile_ly_ui_client build_gui_linux                        compile_server_linux copy_files_linux zip_linux zip_linux_upgrade
ly-mac:   prepare compile_ly_ui_client build_gui_mac                          compile_server_mac   copy_files_mac   zip_mac zip_mac_upgrade

prepare: update_version

update_version: gen_version_file

gen_version_file:
	@echo 'gen version'
	@mkdir -p ${QINIU_DIR}/${PROJECT}/
	@echo '{"version": "${VERSION}"}' > ${QINIU_DIR}/${PROJECT}/version.json

compile_ui:
	@cd ui && yarn build --mode deeptest-local --dest ../client/ui && cd ..
compile_ui_demo:
	@cd ui && yarn build --mode deeptest-demo --dest ../client/ui && cd ..
compile_ui_client:
	@cd ui && yarn build --mode deeptest-client --dest ../client/ui && cd ..
compile_ly_ui_client:
	@cd ui && yarn build --mode ly-client --dest ../client/ui && cd ..

# launcher
compile_launcher_win64:
	@echo 'start compile win64 launcher'
	@cd cmd/launcher && \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win64/${PROJECT}.exe && \
		cd ..


compile_launcher_win32:
	@echo 'start compile win32 launcher'
	@cd cmd/launcher && \
        CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win32/${PROJECT}.exe && \
        cd ..

# server
compile_server_win64:
	@echo 'start compile win64'
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ${BIN_DIR}win64/${PROJECT}-server.exe ${SERVER_MAIN_FILE}

compile_server_win32:
	@echo 'start compile win32'
	@CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ${BIN_DIR}win32/${PROJECT}-server.exe ${SERVER_MAIN_FILE}

compile_server_linux:
	@echo 'start compile linux'
ifeq ($(PLATFORM),"mac")
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-gcc CXX=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-g++ \
		${BUILD_CMD_UNIX} \
		-o ${BIN_DIR}linux/${PROJECT}-server ${SERVER_MAIN_FILE}
else
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc CXX=g++ \
		${BUILD_CMD_UNIX} \
		-o ${BIN_DIR}linux/${PROJECT}-server ${SERVER_MAIN_FILE}
endif

compile_server_mac:
	@echo 'start compile darwin'
	@echo
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
		${BUILD_CMD_UNIX} \
		-o ${BIN_DIR}darwin/${PROJECT}-server ${SERVER_MAIN_FILE}

# agent then gui
build_gui_win64: compile_agent_win64 package_gui_win64_client
compile_agent_win64:
	@echo 'start compile win64'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ${CLIENT_BIN_DIR}win32/deeptest-agent.exe ${AGENT_MAIN_FILE}
	@rm -rf "${CLIENT_OUT_DIR_UPGRADE}win64" && mkdir -p "${CLIENT_OUT_DIR_UPGRADE}win64" && \
		cp ${CLIENT_BIN_DIR}win32/deeptest-agent.exe "${CLIENT_OUT_DIR_UPGRADE}win64"
package_gui_win64_client:
	@cd client && ${PKG_W64} && cd ..
	@rm -rf ${CLIENT_OUT_DIR_EXECUTABLE}win64 && mkdir -p ${CLIENT_OUT_DIR_EXECUTABLE}win64 && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-win32-x64 ${CLIENT_OUT_DIR_EXECUTABLE}win64/gui

build_gui_win32: compile_agent_win32 package_gui_win32_client
compile_agent_win32:
	@echo 'start compile win32'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ${CLIENT_BIN_DIR}win32/deeptest-agent.exe ${AGENT_MAIN_FILE}
	@rm -rf "${CLIENT_OUT_DIR_UPGRADE}win32" && mkdir -p "${CLIENT_OUT_DIR_UPGRADE}win32" && \
		cp ${CLIENT_BIN_DIR}win32/deeptest-agent.exe "${CLIENT_OUT_DIR_UPGRADE}win32"
package_gui_win32_client:
	@cd client && ${PKG_W32} && cd ..
	@rm -rf ${CLIENT_OUT_DIR_EXECUTABLE}win32 && mkdir -p ${CLIENT_OUT_DIR_EXECUTABLE}win32 && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-win32-ia32 ${CLIENT_OUT_DIR_EXECUTABLE}win32/gui

build_gui_linux: compile_agent_linux package_gui_linux_client
compile_agent_linux:
	@echo 'start compile linux'
	@rm -rf ./${CLIENT_BIN_DIR}/*
ifeq ($(PLATFORM),"mac")
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-gcc CXX=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-g++ \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}linux/deeptest-agent ${AGENT_MAIN_FILE}
else
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc CXX=g++ \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}linux/deeptest-agent ${AGENT_MAIN_FILE}
endif
	@rm -rf "${CLIENT_OUT_DIR_UPGRADE}win32" && mkdir -p "${CLIENT_OUT_DIR_UPGRADE}linux" && \
		cp ${CLIENT_BIN_DIR}linux/deeptest-agent "${CLIENT_OUT_DIR_UPGRADE}linux"
package_gui_linux_client:
	@cd client && ${PKG_LINUX} && cd ..
	@rm -rf ${CLIENT_OUT_DIR_EXECUTABLE}linux && mkdir -p ${CLIENT_OUT_DIR_EXECUTABLE}linux && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-linux-x64 ${CLIENT_OUT_DIR_EXECUTABLE}linux/gui

build_gui_mac: compile_agent_mac package_gui_mac_client
compile_agent_mac:
	@echo 'start compile mac'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@echo
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}darwin/deeptest-agent ${AGENT_MAIN_FILE}
	@rm -rf "${CLIENT_OUT_DIR_UPGRADE}darwin" && mkdir -p "${CLIENT_OUT_DIR_UPGRADE}darwin" && \
		cp ${CLIENT_BIN_DIR}darwin/deeptest-agent "${CLIENT_OUT_DIR_UPGRADE}darwin"
package_gui_mac_client:
	@cd client && ${PKG_MAC} && cd ..
	@rm -rf ${CLIENT_OUT_DIR_EXECUTABLE}darwin && mkdir -p ${CLIENT_OUT_DIR_EXECUTABLE}darwin && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-darwin-x64 ${CLIENT_OUT_DIR_EXECUTABLE}darwin/gui && \
		mv ${CLIENT_OUT_DIR_EXECUTABLE}darwin/gui/${PROJECT}.app ${CLIENT_OUT_DIR_EXECUTABLE}darwin/${PROJECT}.app && rm -rf ${CLIENT_OUT_DIR_EXECUTABLE}darwin/gui

# copy files
copy_files_win64:
	@echo 'start copy files win64'
	cp ${BIN_DIR}win64/${PROJECT}.exe "${CLIENT_OUT_DIR_EXECUTABLE}win64"

copy_files_win32:
	@echo 'start copy files win32'
	cp ${BIN_DIR}win32/${PROJECT}.exe "${CLIENT_OUT_DIR_EXECUTABLE}win32"

copy_files_linux:
	@echo 'start copy files linux'

copy_files_mac:
	@echo 'start copy files darwin'

# zip files
zip_win64:
	@echo 'start zip win64'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win64 && rm -rf ${QINIU_DIST_DIR}win64/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}win64 && \
		zip -ry ${QINIU_DIST_DIR}win64/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win64/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win64/${PROJECT}.zip.md5 && \
        cd ../../..; \

zip_win32:
	@echo 'start zip win32'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win32 && rm -rf ${QINIU_DIST_DIR}win32/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}win32 && \
		zip -ry ${QINIU_DIST_DIR}win32/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win32/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win32/${PROJECT}.zip.md5 && \
        cd ../../..; \

zip_linux:
	@echo 'start zip linux'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}linux && rm -rf ${QINIU_DIST_DIR}linux/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}linux && \
		zip -ry ${QINIU_DIST_DIR}linux/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}linux/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}linux/${PROJECT}.zip.md5 && \
        cd ../../..; \

zip_mac:
	@echo 'start zip darwin'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}darwin && rm -rf ${QINIU_DIST_DIR}darwin/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}darwin && \
		zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}.zip.md5 && \
        cd ../../..; \

# zip upgrade package
zip_win64_upgrade:
	@cp -r ${CLIENT_UI_DIR} ${CLIENT_OUT_DIR_UPGRADE}win64/ui
	@cd ${CLIENT_OUT_DIR_UPGRADE}win64 && \
    		zip -ry ${QINIU_DIST_DIR}win64/${PROJECT}-upgrade.zip ./* && \
    		md5sum ${QINIU_DIST_DIR}win64/${PROJECT}-upgrade.zip | awk '{print $$1}' | \
    			xargs echo > ${QINIU_DIST_DIR}win64/${PROJECT}-upgrade.zip.md5 && \
            cd ../../..; \

zip_win32_upgrade:
	@cp -r ${CLIENT_UI_DIR} ${CLIENT_OUT_DIR_UPGRADE}win64/ui
	@cd ${CLIENT_OUT_DIR_UPGRADE}win32 && \
    		zip -ry ${QINIU_DIST_DIR}win32/${PROJECT}-upgrade.zip ./* && \
    		md5sum ${QINIU_DIST_DIR}win32/${PROJECT}-upgrade.zip | awk '{print $$1}' | \
    			xargs echo > ${QINIU_DIST_DIR}win32/${PROJECT}-upgrade.zip.md5 && \
            cd ../../..; \

zip_linux_upgrade:
	@cp -r ${CLIENT_UI_DIR} ${CLIENT_OUT_DIR_UPGRADE}win64/ui
	@cd ${CLIENT_OUT_DIR_UPGRADE}linux && \
    		zip -ry ${QINIU_DIST_DIR}linux/${PROJECT}-upgrade.zip ./* && \
    		md5sum ${QINIU_DIST_DIR}linux/${PROJECT}-upgrade.zip | awk '{print $$1}' | \
    			xargs echo > ${QINIU_DIST_DIR}linux/${PROJECT}-upgrade.zip.md5 && \
            cd ../../..; \

zip_mac_upgrade:
	@cp -r ${CLIENT_UI_DIR} ${CLIENT_OUT_DIR_UPGRADE}darwin/ui
	@cd ${CLIENT_OUT_DIR_UPGRADE}darwin && \
    		zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}-upgrade.zip ./* && \
    		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}-upgrade.zip | awk '{print $$1}' | \
    			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}-upgrade.zip.md5 && \
            cd ../../..; \

upload_to:
	@echo 'upload...'
	@find ${QINIU_DIR} -name ".DS_Store" -type f -delete
	@qshell qupload2 --src-dir=${QINIU_DIR} --bucket=download --thread-count=10 --log-file=qshell.log \
					 --skip-path-prefixes=ztf,zd,zv,zmanager,driver --rescan-local --overwrite --check-hash

demo: checkout compile_server_linux compile_agent_linux compile_ui_demo copy_file copy_file start_service
checkout:
	@git pull
copy_file:
	@cp -f bin/linux/deeptest-server ~/rd/server/deeptest
	@cp -f client/bin/linux/agent ~/rd/server/deeptest/deeptest-agent
	@rm -rf ~/rd/server/deeptest/ui
	@mkdir -p ~/rd/server/deeptest/ui
	@cp -fr client/ui ~/rd/server/deeptest/ui/dist
start_service:
	@ps -ef | grep 'deeptest-' | grep -v grep | awk '{print $2}' | xargs --no-run-if-empty kill -9
	@RUNNER_TRACKING_ID="" nohup ~/rd/server/deeptest/deeptest-server > server.log 2>&1 &
    @RUNNER_TRACKING_ID="" export DemoTestSite=http://111.231.16.35:9000 \
        nohup ~/rd/server/deeptest/deeptest-agent > agent.log 2>&1 &