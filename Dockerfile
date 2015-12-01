FROM dhub.yunpro.cn/zhanglianxiang/jenkins-restapi-base
MAINTAINER zhanglianxiang@goyoo.com

ENV JENKINS_HOST=127.0.0.1 JENKINS_PORT=8080 GOPATH=/root/jenkins_api GOROOT=/usr/lib/go JENKINS_RESTAPI_HOST=127.0.0.1 JENKINS_RESTAPI_PORT=3000 ENVIRONMENT=production
RUN cd /root && git clone https://github.com/zhanglianx111/jenkins_api.git && go get github.com/drone/routes && go get github.com/bndr/gojenkins
EXPOSE 3000
CMD ["go", "run", "/root/jenkins_api/src/server.go"]
