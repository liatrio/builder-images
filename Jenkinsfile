pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  environment {
    SKAFFOLD_DEFAULT_REPO = 'docker.artifactory.liatr.io/liatrio'
  }
  stages {
    stage('Promote version') {
      //when { branch "master" }
      steps {
        container('skaffold') {
          script {
            def headTags = sh returnStdout: true, script: 'git tag -l --points-at HEAD --sort=creatordate v*.*.*'
            echo "Tag HEAD ${headTags}"
            if (!headTags) {
              def tag
              def tags = sh returnStdout: true, script: 'git tag -l --sort=creatordate v*.*.*'
              if (tags) {
                echo "Tags ${tags}"
                def tagParts = tags.split("\n")[-1].substring(1).split('\\.')
                echo "Tag parts ${tagParts[0]} ${tagParts[1]} ${tagParts[2]}"
                tag = "v${tagParts[0]}.${tagParts[1]}.${tagParts[2].toInteger() + 1}"
              } else {
                tag = 'v0.0.1'
              }
              echo "Tag ${tag}"
              sh "git config --global user.email 'jenkins@liatr.io'"
              sh "git config --global user.name 'Liatrio Jenkins Automation'"
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
