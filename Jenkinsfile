library 'LEAD'

pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  environment {
    VERSION = version()
    GITOPS_GIT_URL = 'git@github.com:liatrio/lead-environments.git'
    GITOPS_REPO_FILE = 'aws/liatrio-sandbox/terragrunt.hcl'
    GITOPS_VALUES = 'inputs.builder_images_version=${VERSION}:inputs.jenkins_image_version${VERSION}'
    GITOPS_GIT_USERNAME = ''
    GITOPS_GIT_PASSWORD = ''

  }
  stages {
    stage('Build & publish images') {
      steps {
        notifyPipelineStart()
        notifyStageStart()
        container('skaffold') {
          sh "make all"
            script {
              def version = sh ( script: "make version", returnStdout: true).trim()
              notifyStageEnd([status: "Published new images: ${version}"])
            }
        }
      }
      post {
        success {
          agent {
            label "lead-toolchain-gitops"
          }
        }
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
  }
}
def version() {
    retrun sh(script: "git describe --tags --dirty", returnStdout: true).substring(1);
}
