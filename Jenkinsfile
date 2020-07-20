pipeline {
    agent any

    stages {
        stage('Build auth service') {
            steps {
                sh 'go build -o auth-service/bin/main auth-service/main.go'
                sh 'rm -rf auth-service/*.go'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploy'
            }
        }
    }
}