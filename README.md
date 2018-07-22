# Next Generation Testing Tools

考虑到部署和DQL（Defect Query Language）设计的方便性，正在迁移到SpringBoot、MyBatis技术栈！

AngularJS, SockJS, SpringBoot, MyBatis, MySQL

### Quick start
```bash
Setup Tomcat8 and Mysql5.x server
Import database from xdoc/ngtesting-dump.sql
Deploy webapp to path /platform in tomcat
Open http://localhost:8080/platform to test backend web and database server works well
Open chrome browser and goto http://localhost:8080/platform/client
If you use different tomcat path, search and replace below service url to your own in main.*.js
- http://localhost:8080/platform/
```

### Test Project
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/project_view.jpg" width="800px" style="margin: 10px auto;">

### Test Case
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/case_edit.jpg" width="800px" style="margin: 10px auto;">

### Test Execution
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/case_exe.jpg" width="800px" style="margin: 10px auto;">

### Test Plan
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_result.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_progress.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_process.jpg" width="800px" style="margin: 10px auto;">
<img src="https://raw.githubusercontent.com/aaronchen2k/ngtesting-platform/master/xdoc/capture/plan_exe_process_by_user.jpg" width="800px" style="margin: 10px auto;">

### Licenses

All source code is licensed under the [GPLv3 License](LICENSE.md).
