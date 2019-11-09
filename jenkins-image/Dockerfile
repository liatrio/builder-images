FROM jenkins/jenkins:2.192
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt
RUN curl https://artifactory.toolchain.lead.prod.liatr.io/artifactory/maven/pipeline-status-0.1.1-17-g62d3666.hpi --output /usr/share/jenkins/ref/plugins/pipeline-status-plugin.hpi
