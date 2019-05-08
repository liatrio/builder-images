pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  environment {
    SKAFFOLD_DEFAULT_REPO = 'docker.artifactory.liatr.io/liatrio'
  }
  stages {
    stage('Promote version') {
      when { branch "master" }
      steps {
        container('skaffold') {
          script {
            def tagHead = sh returnStdout: true, script: 'git tag -l --points-at HEAD v*.*.* | tail -1'
            echo "Tag HEAD ${tagHead}"
            if (!headTag) {
              def tagLast = sh returnStdout: true, script: 'git tag -l v*.*.*'
              if (tagLast == '')
              echo "Tag last ${tagLast}"
              def tagParts = tagLast.substring(1).split('.')
              echo "Tag parts ${tagParts[0]} ${tagParts[1]} ${tagParts[2]}"
              tag = "v${tagParts[0]}.${tagParts[1]}.${tagParts[2] + 1}"
              echo "Tag ${tag}"
              sh "git tag -a -m 'releasing ${tag}' ${tag}"
              sh "git push origin ${tag}"
            }
          }
        }
      }
    }
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
