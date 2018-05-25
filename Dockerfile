# VERSION 0.0.1
FROM registry.cn-hangzhou.aliyuncs.com/ngtesting/deploy

MAINTAINER Aaron "462826@qq.com"

ENV PATH /usr/local/bin:$PATH

ENV JAVA_HOME /home/ngt/dev/sdk/jdk1.8.0_172
ENV JRE_HOME $JAVA_HOME/jre
ENV CLASSPATH .:$JAVA_HOME/lib:$JRE_HOME/lib
ENV PATH .:$JAVA_HOME/bin:$PATH

ENV MAVEN_HOME /home/ngt/dev/tool/apache-maven-3.5.3
ENV PATH .:$MAVEN_HOME/bin:$PATH

ENV CATALINA_HOME /home/ngt/dev/server/apache-tomcat-8.5.31

RUN apt-get -y update

WORKDIR /home/ngt/dev/project
RUN rm -rf /home/ngt/dev/project/ngtesting-platform
RUN git clone https://github.com/aaronchen2k/ngtesting-platform.git
WORKDIR /home/ngt/dev/project/ngtesting-platform
RUN mvn -Denv=dev clean package
RUN rm -rf /home/ngt/dev/server/apache-tomcat-8.5.31/webapps/ngtesting-platform-*
RUN cp target/ngtesting-platform-*.war /home/ngt/dev/server/apache-tomcat-8.5.31/webapps

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]
