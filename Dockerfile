FROM dhub.yunpro.cn/zhanglianxiang/jenkins-restapi-base
MAINTAINER zhanglianxiang@goyoo.com

ENV JENKINS_HOST=10.12.1.133 JENKINS_PORT=8080 GOPATH=/root/jenkins_api GOROOT=/usr/lib/go ENVIRONMENT=production
RUN cd /root && git clone -b master http://git.yunpro.cn/zhanglianxiang/jenkins_api.git && go get github.com/bndr/gojenkins && go get github.com/go-martini/martini && go get github.com/Sirupsen/logrus && go get github.com/beevik/etree && go get github.com/bitly/go-simplejson && go get github.com/clbanning/mxj && go get github.com/tsuru/config 
EXPOSE 3000
WORKDIR /root/jenkins_api/src


CMD ["go", "run", "server.go"]
