pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
        DOCKER_CONFIG = "${WORKSPACE}/.docker"
    }

    stages {

        stage('Prepare Docker Config') {
            steps {
                sh 'mkdir -p $DOCKER_CONFIG'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh '/opt/homebrew/bin/go mod tidy'
                sh '/opt/homebrew/bin/go test -v ./...'
            }
        }

        stage('Docker Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'docker-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh 'echo $DOCKER_PASS | /usr/local/bin/docker login -u $DOCKER_USER --password-stdin'
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