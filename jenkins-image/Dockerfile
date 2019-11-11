FROM jenkins/jenkins:2.192
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt

ENV PIPELINE_PLUGIN_VERSION 1.0.0-0
RUN curl -Lo /usr/share/jenkins/ref/plugins/pipeline-status-plugin.jpi \
     https://artifactory.toolchain.lead.prod.liatr.io/artifactory/maven/io/jenkins/plugins/pipeline-status-plugin/${PIPELINE_PLUGIN_VERSION}/pipeline-status-plugin-${PIPELINE_PLUGIN_VERSION}.hpi
