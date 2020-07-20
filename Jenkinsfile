pipeline {
    agent any
    tools {
        go 'go-1.14.6'
    }
    environment {
        GO114MODULE = 'on'
    }
    stages {
        stage('Build auth service') {
            steps {
                sh 'cd auth-service'
                sh 'go mod init local.com/golangMarketplace/authService'
                sh 'go build -o bin/main main.go'
                sh 'rm -rf *.go'
                sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                sh 'docker build -t alexandruubytex/golang_marketplace_auth_service:latest .'
                sh 'docker push alexandruubytex/golang_marketplace_auth_service:latest'
            }
        }

        stage('Deploy auth service') {
            steps {
                echo 'Deploy auth service'
            }
        }
    }
}