pipeline {
    agent none
    tools {
        go 'go-1.14.6'
    }
    environment {
        GO114MODULE = 'on'
    }
    stages {
        stage('Build auth service') {
            agent {
                docker { image 'docker:18.09.7' }
            }
            steps {
                dir("auth-service") {
                    sh 'apk add go'
                    sh 'apk add git'
                    sh 'apk add libc-dev'
                    sh 'go build -o bin/main main.go'
                    sh 'rm -rf *.go'
                    sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                    sh 'docker build -t alexandruubytex/golang_marketplace_auth_service:latest .'
                    sh 'docker push alexandruubytex/golang_marketplace_auth_service:latest'
                }
            }
        }

        stage('Build UI service') {
            steps {
                echo 'Build UI service'
            }
        }
    }
}