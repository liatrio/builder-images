pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  environment {
    SKAFFOLD_DEFAULT_REPO = 'docker.artifactory.liatr.io/liatrio'
  }
  stages {
    stage('Build & publish container') {
      steps {
        container('skaffold') {
          script {
            docker.withRegistry("https://${SKAFFOLD_DEFAULT_REPO}", 'jenkins-credential-artifactory') {
              sh "make"
            }
          }
        }
      }
    }
  }
}
