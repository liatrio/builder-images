library 'LEAD'

pipeline {
  agent {
    label "lead-toolchain-skaffold"
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
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
  }
}
