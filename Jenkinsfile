pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
    }

    stages {

        stage('Run Unit Tests') {
            steps {
                sh '/opt/homebrew/bin/go mod tidy'
                sh '/opt/homebrew/bin/go test -v ./...'
            }
        }

        stage('Docker Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '/usr/local/bin/docker login -u $DOCKER_USER -p $DOCKER_PASS'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '/usr/local/bin/docker build -t $DOCKER_IMAGE:latest .'
            }
        }

        stage('Push To Docker Hub') {
            steps {
                sh '/usr/local/bin/docker push $DOCKER_IMAGE:latest'
            }
        }
    }
}