library 'LEAD'

pipeline {
  agent any
  environment {
    VERSION = version()
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
        }
        notifyStageEnd([status: "Published new images: ${VERSION}"])
      }
      post {
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
    stage('GitOps: Update sandbox') {
      when {
        branch 'master'
      }
      agent {
        label "lead-toolchain-gitops"
      }
      environment {
        GITOPS_GIT_URL = 'https://github.com/liatrio/lead-environments.git'
        GITOPS_REPO_FILE = 'aws/liatrio-sandbox/terragrunt.hcl'
        GITOPS_VALUES = "inputs.builder_images_version=${VERSION}:inputs.jenkins_image_version=${VERSION}"
      }
      steps {
        container('gitops') {
          sh "/go/bin/gitops"
        }
      }
    }
  }
}
def version() {
    return sh(script: "git describe --tags --dirty", returnStdout: true).trim()
}
