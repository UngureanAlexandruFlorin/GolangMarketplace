pipeline {
    agent any
    tools {
        go 'go-1.14.6'
        nodejs 'node-14.6.0'
    }
    environment {
        GO114MODULE = 'on'
    }
    stages {
        stage('Build UI service') {
            steps {
                dir("ui-service") {
                    sh 'npm i'
                    sh 'npm run build'
                    sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                    sh 'docker build -t alexandruubytex/golang_marketplace_ui_service:latest .'
                    sh 'docker push alexandruubytex/golang_marketplace_ui_service:latest'
                }
            }
        }

        stage('Build auth service') {
            steps {
                dir("auth-service") {
                    
                    sh 'go build -o ./bin/main main.go'
                    sh 'rm -rf *.go'
                    sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                    sh 'docker build -t alexandruubytex/golang_marketplace_auth_service:latest .'
                    sh 'docker push alexandruubytex/golang_marketplace_auth_service:latest'
                }
            }
        }

        stage('Build product service') {
            steps {
                dir("product-service") {
                    
                    sh 'go build -o ./bin/main main.go'
                    sh 'rm -rf *.go'
                    sh 'docker login -u \'alexandruubytex\' -p \'333Albastru333\''
                    sh 'docker build -t alexandruubytex/golang_marketplace_product_service:latest .'
                    sh 'docker push alexandruubytex/golang_marketplace_product_service:latest'
                }
            }
        }
    }
}