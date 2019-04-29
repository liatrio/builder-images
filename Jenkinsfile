pipeline {
    agent {
        label "jenkins-jx-base"
    }
    environment {
      SKAFFOLD_DEFAULT_REPO = 'docker.artifactory.liatr.io/liatrio'
    }
    stages {
        stage('Build') {
            steps {
                container('jx-base') {
                    script {
                      docker.withRegistry("https://${SKAFFOLD_DEFAULT_REPO}", 'artifactory-credentials') {
                          sh "skaffold build --skip-tests" 
                      }
                    }
                }
            }
        }
    }
}