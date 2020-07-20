pipeline {
    agent any

    stages {
        stage('Build auth service') {
            steps {
                sh 'go build -o auth-service/bin/main auth-service/main.go'
                sh 'rm -rf auth-service/*.go'
                sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                sh 'docker build -t alexandruubytex/golang_marketplace_auth_service:latest /auth-service/'
                sh 'docker push alexandruubytex/golang_marketplace_auth_service:latest'
            }
        }

        stage('Deploy auth service') {
            steps {
                echo 'Deploy auth'
            }
        }
    }
}