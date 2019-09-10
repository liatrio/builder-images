library 'LEAD'

pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  stages {
    stage('Build & publish container') {
      steps {
        notifyPipelineStart()
        notifyStageStart()
        container('skaffold') {
          sh "make all"
            script {
              def version = sh ( script: "make version", returnStdout: true).trim()
              notifyStageEnd([status: "Published new charts: ${version}"])
            }
        }
      }
      post {
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
  }
}
