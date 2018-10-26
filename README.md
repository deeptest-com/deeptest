# Next Generation Testing Tools
有兴趣请加微信, Wechat: 462826

AngularJS, SockJS, SpringBoot, MyBatis, MySQL

Demo: http://47.99.102.138:8080

## Quick start
### Use Docker

1. Install docker and docker-compose;
2. Download file from https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/docker-compose.yml ;
3. In same dir, enter "docker-compose up" to launch.

   First time, you may get an error caused by MySQL service not ready for web server connectting, just run Step 3 again to fix.
   
   **Once you create and start containers by using "docker-compose up" command, next time you should use "docker-compose start|stop" to avoid to docker-compose re-creating the container after the remote image changed, which will cause data-losing.**
4. Open http://localhost:58080/test.html using Chrome.

### Use Java Jar
1. Create dir "/work/ngtesting-data/" for file upload;
2. Create a MySQL database named "ngtesting-web";
3. Add MySQL user "ngtesting" with password "P2ssw0rd";
4. Run "git clone https://github.com/aaronchen2k/ngtesting-platform.git";
5. Import "src/main/docker/mysql/schema.sql" to database;
6. In project dir, enter command "xdoc/ngtesting-web-1.*.jar" to run;

   You may change the config params in application.yml file like this: "xdoc/ngtesting-web-1.*.jar --config.mysql.host=10.0.0.10 --config.mysql.port=3306".
7. Open http://localhost:8080/test.html using Chrome.

## Test Project
![project_view](xdoc/capture/project_view.jpg)

## Test Case
![project_view](xdoc/capture/case_edit.jpg)

## Test Execution
![project_view](xdoc/capture/case_exe.jpg)

## Test Plan
![project_view](xdoc/capture/plan_exe_result.jpg)
![project_view](xdoc/capture/plan_exe_progress.jpg)
![project_view](xdoc/capture/plan_exe_process.jpg)
![project_view](xdoc/capture/plan_exe_process_by_user.jpg)

## Test Plan
![issue_query](xdoc/capture/issue_query.jpg)

## Licenses

All source code is licensed under the [GPLv3 License](LICENSE.md).
