VERSION=0.8
PROJECT=deeptest

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
    QINIU_DIR=~/work/zentao/qiniu/
else
    QINIU_DIR=~/deeptestZip
endif

QINIU_DIST_DIR=${QINIU_DIR}${PROJECT}/${VERSION}/

SERVER_MAIN_FILE=cmd/server/main.go
AGENT_MAIN_FILE=cmd/agent/main.go

BIN_DIR=bin/
CLIENT_BIN_DIR=client/bin/
CLIENT_OUT_DIR=client/out/

BUILD_TIME=`git show -s --format=%cd`
GO_VERSION=`go version`
GIT_HASH=`git show -s --format=%H`
BUILD_CMD_UNIX=go build -ldflags "-X 'main.AppVersion=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GoVersion=${GO_VERSION}' -X 'main.GitHash=${GIT_HASH}'"
BUILD_CMD_WIN=go build -ldflags "-s -w -X 'main.AppVersion=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GoVersion=${GO_VERSION}' -X 'main.GitHash=${GIT_HASH}'"

default: win64 win32 linux mac

win64: prepare build_gui_win64 compile_launcher_win64 compile_server_win64 copy_files_win64 zip_win64
win32: prepare build_gui_win32 compile_launcher_win32 compile_server_win32 copy_files_win32 zip_win32
linux: prepare build_gui_linux                        compile_server_linux copy_files_linux zip_linux
mac:   prepare build_gui_mac                          compile_server_mac   copy_files_mac   zip_mac

prepare: update_version

update_version: gen_version_file

gen_version_file:
	@echo 'gen version'
	@mkdir -p ${QINIU_DIR}/${PROJECT}/
	@echo ${VERSION} > ${QINIU_DIR}/${PROJECT}/version.txt

compile_ui:
	@cd ui && yarn build --dest ../client/ui && cd ..

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

# launcher
compile_launcher_win64:
	@echo 'start compile win64 launcher'
	@cd cmd/launcher && \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win64/${PROJECT}-gui.exe && \
		cd ..

compile_launcher_win32:
	@echo 'start compile win32 launcher'
	@cd cmd/launcher && \
        CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ../../${BIN_DIR}win32/${PROJECT}-gui.exe && \
        cd ..

# gui
build_gui_win64: compile_gui_win64 package_gui_win64_client
compile_gui_win64:
	@echo 'start compile win64'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 \
		${BUILD_CMD_WIN} -x -v \
		-o ${CLIENT_BIN_DIR}win32/${PROJECT}.exe ${AGENT_MAIN_FILE}
package_gui_win64_client:
	@cd client && npm run package-win64 && cd ..
	@rm -rf ${CLIENT_OUT_DIR}win64 && mkdir ${CLIENT_OUT_DIR}win64 && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-win32-x64 ${CLIENT_OUT_DIR}win64/gui

build_gui_win32: compile_gui_win32 package_gui_win32_client
compile_gui_win32:
	@echo 'start compile win32'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 \
		${BUILD_CMD_WIN} -x -v \
		-o ${CLIENT_BIN_DIR}win32/${PROJECT}.exe ${AGENT_MAIN_FILE}
package_gui_win32_client:
	@cd client && npm run package-win32 && cd ..
	@rm -rf ${CLIENT_OUT_DIR}win32 && mkdir ${CLIENT_OUT_DIR}win32 && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-win32-ia32 ${CLIENT_OUT_DIR}win32/gui

build_gui_linux: compile_gui_linux package_gui_linux_client
compile_gui_linux:
	@echo 'start compile linux'
	@rm -rf ./${CLIENT_BIN_DIR}/*
ifeq ($(PLATFORM),"mac")
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-gcc CXX=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-g++ \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}linux/${PROJECT} ${AGENT_MAIN_FILE}
else
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc CXX=g++ \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}linux/${PROJECT} ${AGENT_MAIN_FILE}
endif
package_gui_linux_client:
	@cd client && npm run package-linux && cd ..
	@rm -rf ${CLIENT_OUT_DIR}linux && mkdir ${CLIENT_OUT_DIR}linux && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-linux-x64 ${CLIENT_OUT_DIR}linux/gui

build_gui_mac: compile_gui_mac package_gui_mac_client
compile_gui_mac:
	@echo 'start compile mac'
	@rm -rf ./${CLIENT_BIN_DIR}/*
	@echo
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
		${BUILD_CMD_UNIX} \
		-o ${CLIENT_BIN_DIR}darwin/${PROJECT} ${AGENT_MAIN_FILE}
package_gui_mac_client:
	@cd client && npm run package-mac && cd ..
	@rm -rf ${CLIENT_OUT_DIR}darwin && mkdir ${CLIENT_OUT_DIR}darwin && \
		mv ${CLIENT_OUT_DIR}${PROJECT}-darwin-x64 ${CLIENT_OUT_DIR}darwin/gui && \
		mv ${CLIENT_OUT_DIR}darwin/gui/${PROJECT}.app ${CLIENT_OUT_DIR}darwin/${PROJECT}.app && rm -rf ${CLIENT_OUT_DIR}darwin/gui

# copy files
copy_files_win64:
	@echo 'start copy files win64'
	cp ${BIN_DIR}win64/${PROJECT}-gui.exe "${CLIENT_OUT_DIR}win64"

copy_files_win32:
	@echo 'start copy files win32'
	cp ${BIN_DIR}win32/${PROJECT}-gui.exe "${CLIENT_OUT_DIR}win32"

copy_files_linux:
	@echo 'start copy files linux'

copy_files_mac:
	@echo 'start copy files darwin'

# zip files
zip_win64:
	@echo 'start zip win64'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win64 && rm -rf ${QINIU_DIST_DIR}win64/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR}win64 && \
		zip -ry ${QINIU_DIST_DIR}win64/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win64/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win64/${PROJECT}.zip.md5 && \
        cd ../..; \

zip_win32:
	@echo 'start zip win32'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}win32 && rm -rf ${QINIU_DIST_DIR}win32/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR}win32 && \
		zip -ry ${QINIU_DIST_DIR}win32/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}win32/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}win32/${PROJECT}.zip.md5 && \
        cd ../..; \

zip_linux:
	@echo 'start zip linux'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}linux && rm -rf ${QINIU_DIST_DIR}linux/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR}linux && \
		zip -ry ${QINIU_DIST_DIR}linux/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}linux/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}linux/${PROJECT}.zip.md5 && \
        cd ../..; \

zip_mac:
	@echo 'start zip darwin'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@mkdir -p ${QINIU_DIST_DIR}darwin && rm -rf ${QINIU_DIST_DIR}darwin/${PROJECT}.zip
	@cd ${CLIENT_OUT_DIR}darwin && \
		zip -ry ${QINIU_DIST_DIR}darwin/${PROJECT}.zip ./* && \
		md5sum ${QINIU_DIST_DIR}darwin/${PROJECT}.zip | awk '{print $$1}' | \
			xargs echo > ${QINIU_DIST_DIR}darwin/${PROJECT}.zip.md5 && \
        cd ../..; \

upload_to:
	@echo 'upload...'
	@find ${QINIU_DIR} -name ".DS_Store" -type f -delete
	@qshell qupload2 --src-dir=${QINIU_DIR} --bucket=download --thread-count=10 --log-file=qshell.log \
					 --skip-path-prefixes=zz,zd,zmanager,driver --rescan-local --overwrite --check-hash
