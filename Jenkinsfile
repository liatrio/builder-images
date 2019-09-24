library 'LEAD'

pipeline {
  agent any
  environment {
    VERSION = version()
    GITOPS_GIT_URL = 'git@github.com:liatrio/lead-environments.git'
    GITOPS_REPO_FILE = 'aws/liatrio-sandbox/terragrunt.hcl'
    GITOPS_VALUES = 'inputs.builder_images_version=${VERSION}:inputs.jenkins_image_version=${VERSION}'
    GITOPS_GIT_USERNAME = ''
    GITOPS_GIT_PASSWORD = ''

  }
  stages {
    stage('Build & publish images') {
      agent {
        label "lead-toolchain-skaffold"
      }
      steps {
        notifyPipelineStart()
        notifyStageStart()
        container('skaffold') {
          sh "make all"
          script {
            // def version = sh ( script: "make version", returnStdout: true).trim()
            notifyStageEnd([status: "Published new images: ${VERSION}"])
          }
        }
      }
      post {
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
    stage('GitOps: Update sandbox') {
      when {
        branch 'ENG-1183'
      }
      agent {
        label "lead-toolchain-gitops"
      }
      environment {
        GITOPS_GIT_URL = 'git@github.com:liatrio/lead-environments.git'
        GITOPS_REPO_FILE = 'aws/liatrio-sandbox/terragrunt.hcl'
        GITOPS_VALUES = 'inputs.builder_images_version=${VERSION}:inputs.jenkins_image_version=${VERSION}'
        GITOPS_GIT_USERNAME = ''
        GITOPS_GIT_PASSWORD = ''
      }
      steps {
        echo "====++++something++++===="
      }
    }
  }
}
def version() {
    return sh(script: "git describe --tags --dirty", returnStdout: true).substring(1);
}
