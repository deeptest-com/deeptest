# Next Generation Testing Tools

AngularJS, SockJS, SpringBoot, MyBatis, MySQL

Demo: http://47.99.102.138:8080

## Quick start
### Use Docker

1. Install docker and docker-compose;
2. Run "git clone -b branch-1.0.x https://github.com/aaronchen2k/ngtesting-platform.git" to clone or download codes from https://github.com/aaronchen2k/ngtesting-platform/archive/branch-1.0.x.zip;
3. Goto project dir, Enter "docker-compose up" to launch.

   First time, you may get an error caused by MySQL service not ready for web server connectting, just run Step 2 again to fix.
   
   **Once you create and start containers by using "docker-compose up" command, next time you should use "docker-compose start|stop" to avoid to docker-compose re-creating the container after the remote image changed, which will cause data-losing.**
4. Open http://localhost:58080/test.html using Chrome.

### Use Java Jar
1. Create dir "/work/ngtesting-data/" for file upload;
2. Create a MySQL database named "ngtesting-web";
3. Add MySQL user "ngtesting" with password "P2ssw0rd";
4. Run "git clone -b branch-1.0.x https://github.com/aaronchen2k/ngtesting-platform.git" to clone or download codes from https://github.com/aaronchen2k/ngtesting-platform/archive/branch-1.0.x.zip;
5. Import "src/main/docker/mysql/schema.sql" to database;
6. In project dir, enter command "xdoc/ngtesting-web-1.0.jar“ to run this executable jar;

   You may change the config params in application.yml file like this: "xdoc/ngtesting-web-1.0.jar --config.mysql.host=10.0.0.10 --config.mysql.port=3306".
7. Open http://localhost:8080/test.html using Chrome.

## Test Project
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/project_view.jpg" width="800px" style="margin: 10px auto;">

## Test Case
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/case_edit.jpg" width="800px" style="margin: 10px auto;">

## Test Execution
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/case_exe.jpg" width="800px" style="margin: 10px auto;">

## Test Plan
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_result.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_progress.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_process.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_process_by_user.jpg" width="800px" style="margin: 10px auto;">

## Licenses

All source code is licensed under the [GPLv3 License](LICENSE.md).
