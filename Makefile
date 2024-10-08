ENV        ?= ly
include    env.$(ENV)
print_env:
	@echo $(PROJECT)@$(VERSION)

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


# ly 打包的路径,单独设置
PKG_W64=npm run package-win64
PKG_MAC=npm run package-mac
PKG_W32=npm run package-win32
PKG_LINUX=npm run package-linux
ifeq ($(PROJECT),ThirdpartyAPI)
	QINIU_DIR=~/nk2/ly/
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

web: compile_ui_web win64-web win32-web linux-web mac-web

# 非客户端版本打包，需先运行 make compile_ui_web
win64-web: prepare compile_agent_win64 compile_server_win64 zip_win64_web
win32-web: prepare compile_agent_win32 compile_server_win32 zip_win32_web
linux-web: prepare compile_agent_linux compile_server_linux zip_linux_web
mac-web:   prepare compile_agent_mac   compile_server_mac   zip_mac_web

# 客户端版本打包，需先运行 make compile_ui_<target>
dp-win64-client: prepare build_gui_win64 compile_launcher_win64 compile_server_win64 copy_files_win64 zip_win64_client zip_win64_upgrade
dp-win32-client: prepare build_gui_win32 compile_launcher_win32 compile_server_win32 copy_files_win32 zip_win32_client zip_win32_upgrade
dp-linux-client: prepare build_gui_linux                        compile_server_linux copy_files_linux zip_linux_client zip_linux_upgrade
dp-mac-client:   prepare build_gui_mac        					 compile_server_mac   copy_files_mac   zip_mac_client zip_mac_upgrade
dp-mac-test-client: compile_ui_client_test build_gui_mac

# 第三方 打包
ly-win64-client: prepare compile_ly_ui_client build_gui_win64 compile_ly_launcher_win64 compile_server_win64 copy_files_win64 zip_win64_client zip_win64_upgrade
ly-win32-client: prepare compile_ly_ui_client build_gui_win32 compile_ly_launcher_win32 compile_server_win32 copy_files_win32 zip_win32_client zip_win32_upgrade
ly-linux-client: prepare compile_ly_ui_client build_gui_linux                        compile_server_linux copy_files_linux zip_linux_client zip_linux_upgrade
ly-mac-client:   prepare compile_ly_ui_client build_gui_mac                          compile_server_mac   copy_files_mac   zip_mac_client zip_mac_upgrade

prepare: init_client_project update_version

# 初始化客户端项目
init_client_project:
	@sh ./init.project.sh && yarn config set ignore-engines true

update_version: gen_version_file

gen_version_file:
	@echo 'gen version'
	@mkdir -p ${QINIU_DIR}/${PROJECT}/
	@echo '{"version": "${VERSION}"}' > ${QINIU_DIR}/${PROJECT}/version.json

compile_ui:
	@cd ui && yarn build --dest ../client/ui && cd ..
compile_ui_demo: # DeepTest测试
	@cd ../deeptest-ui && yarn build:demo --dest ../deeptest/client/ui && cd ../deeptest
compile_ui_nancal: # 内网测试
	@cd ../deeptest-ui && yarn build:nancal --dest ../deeptest/client/ui && cd ../deeptest
compile_ui_client:
	@rm -rf client/ui && cd ../thirdpartyapi-frontend && yarn build:client --dest ../thirdpartyapi-backend/client/ui && cd ..
compile_ui_client_test:
	@rm -rf client/ui && cd ../thirdpartyapi-frontend && yarn build:clientTest --dest ../thirdpartyapi-backend/client/ui && cd ..
compile_ly_ui_client:
	@cd ../thirdpartyapi-frontend  && yarn build:client && cd ../thirdpartyapi-backend

compile_ui_web:
	@cd ../thirdpartyapi-frontend && yarn build:web --dest ../thirdpartyapi-backend/bin/deeptest-ui && cd ../thirdpartyapi-backend

# launcher
compile_launcher_win64:
	@echo 'start compile win64 launcher'
	@cd cmd/launcher && \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win64/${PROJECT}.exe && \
		cd ..

# ly-launcher
compile_ly_launcher_win64:
	@echo 'start compile win64 launcher'
	@cd cmd/ly-launcher && \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win64/${PROJECT}.exe && \
		cd ..

# launcher-win32
compile_launcher_win32:
	@echo 'start compile win32 launcher'
	@cd cmd/launcher && \
        CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win32/${PROJECT}.exe && \
        cd ..

# ly-launcher-win32
compile_ly_launcher_win32:
	@echo 'start compile win32 launcher'
	@cd cmd/ly-launcher && \
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

# zip web files
zip_win64_web:
	@echo 'start zip win64'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win64 && rm -rf ${QINIU_DIST_DIR}win64/${PROJECT}-web.zip

	@cp -rf config ${BIN_DIR}win64/
	@cp -rf ${CLIENT_BIN_DIR}win32/deeptest-agent.exe ${BIN_DIR}win64/
	@rm -rf ${BIN_DIR}win64/deeptest-ui
	@cp -rf bin/deeptest-ui ${BIN_DIR}win64

	@cd ${BIN_DIR}win64/ && \
		zip -ry ${QINIU_DIST_DIR}win64/${PROJECT}-web.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win64/${PROJECT}-web.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win64/${PROJECT}-web.zip.md5 && \
        cd ../..; \

zip_win32_web:
	@echo 'start zip win32'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win32 && rm -rf ${QINIU_DIST_DIR}win32/${PROJECT}-web.zip

	@cp -rf config ${BIN_DIR}win32/
	@cp -rf ${CLIENT_BIN_DIR}win32/deeptest-agent.exe ${BIN_DIR}win32/
	@rm -rf ${BIN_DIR}win32/deeptest-ui
	@cp -rf bin/deeptest-ui ${BIN_DIR}win32

	@cd ${BIN_DIR}win32/ && \
		zip -ry ${QINIU_DIST_DIR}win32/${PROJECT}-web.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win32/${PROJECT}-web.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win32/${PROJECT}-web.zip.md5 && \
        cd ../..; \

zip_linux_web:
	@echo 'start zip linux'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}linux && rm -rf ${QINIU_DIST_DIR}linux/${PROJECT}-web.zip

	@cp -rf config ${BIN_DIR}linux/
	@cp -rf ${CLIENT_BIN_DIR}linux/deeptest-agent ${BIN_DIR}linux/
	@rm -rf ${BIN_DIR}linux/deeptest-ui
	@cp -rf bin/deeptest-ui ${BIN_DIR}linux

	@cd ${BIN_DIR}linux/ && \
		zip -ry ${QINIU_DIST_DIR}linux/${PROJECT}-web.zip ./* && \
		md5sum ${QINIU_DIST_DIR}linux/${PROJECT}-web.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}linux/${PROJECT}-web.zip.md5 && \
        cd ../..; \

zip_mac_web:
	@echo 'start zip darwin'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}darwin && rm -rf ${QINIU_DIST_DIR}darwin/${PROJECT}-web.zip

	@cp -rf config ${BIN_DIR}darwin/
	@cp -rf ${CLIENT_BIN_DIR}darwin/deeptest-agent ${BIN_DIR}darwin/
	@rm -rf ${BIN_DIR}darwin/deeptest-ui
	@cp -rf bin/deeptest-ui ${BIN_DIR}darwin

	@cd ${BIN_DIR}darwin/ && \
		zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}-web.zip ./* -x Thirdparty* && \
		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}-web.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}-web.zip.md5 && \
        cd ../..; \

# zip client files
zip_win64_client:
	@echo 'start zip win64'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win64 && rm -rf ${QINIU_DIST_DIR}win64/${PROJECT}-client.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}win64 && \
		zip -ry ${QINIU_DIST_DIR}win64/${PROJECT}-client.zip ./* -x Thirdparty* && \
		md5sum ${QINIU_DIST_DIR}win64/${PROJECT}-client.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win64/${PROJECT}-client.zip.md5 && \
        cd ../../..; \

zip_win32_client:
	@echo 'start zip win32'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win32 && rm -rf ${QINIU_DIST_DIR}win32/${PROJECT}-client.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}win32 && \
		zip -ry ${QINIU_DIST_DIR}win32/${PROJECT}-client.zip ./* -x Thirdparty* && \
		md5sum ${QINIU_DIST_DIR}win32/${PROJECT}-client.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win32/${PROJECT}-client.zip.md5 && \
        cd ../../..; \

zip_linux_client:
	@echo 'start zip linux'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}linux && rm -rf ${QINIU_DIST_DIR}linux/${PROJECT}-client.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}linux && \
		zip -ry ${QINIU_DIST_DIR}linux/${PROJECT}-client.zip ./* -x Thirdparty* && \
		md5sum ${QINIU_DIST_DIR}linux/${PROJECT}-client.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}linux/${PROJECT}-client.zip.md5 && \
        cd ../../..; \

zip_mac_client:
	@echo 'start zip darwin'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}darwin && rm -rf ${QINIU_DIST_DIR}darwin/${PROJECT}-client.zip
	@cd ${CLIENT_OUT_DIR_EXECUTABLE}darwin && \
		zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}-client.zip ./* -x Thirdparty* && \
		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}-client.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}-client.zip.md5 && \
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

zip_server:
	@cp bin/win64/${PROJECT}-server.exe ${QINIU_DIST_DIR}win64/

	@cp bin/win32/${PROJECT}-server.exe ${QINIU_DIST_DIR}win32/
	@cp bin/linux/${PROJECT}-server ${QINIU_DIST_DIR}linux/

	# zip -r /Users/aaron/work/qiniu/deeptest/3.0/darwin/deeptest-server.zip config bin/deeptest-ui bin/darwin/deeptest-server
	@zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}-server.zip \
			bin/darwin/${PROJECT}-server config client/ui && \
		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}-upgrade.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}-upgrade.zip.md5 && \

DEMO_BIN=/home/lighthouse/rd/server/deeptest

demo: checkout compile_server_linux compile_agent_linux compile_ui_demo copy_file copy_file start_service
checkout:
	@git pull
copy_file:
	@cp -f bin/linux/deeptest-server ${DEMO_BIN}
	@cp -f client/bin/linux/agent ${DEMO_BIN}/deeptest-agent
	@rm -rf ${DEMO_BIN}/ui
	@mkdir -p ${DEMO_BIN}/ui
	@cp -fr client/ui ${DEMO_BIN}/ui/dist
start_service:
	@ps -ef | grep 'deeptest-' | grep -v grep | awk '{print $2}' | xargs --no-run-if-empty kill -9
	@RUNNER_TRACKING_ID="" nohup ${DEMO_BIN}/deeptest-server > ${DEMO_BIN}/server.log 2>&1 &
    @RUNNER_TRACKING_ID="" export DemoTestSite=http://111.231.16.35:9000 \
        nohup ${DEMO_BIN}/deeptest-agent > ${DEMO_BIN}/agent.log 2>&1 &