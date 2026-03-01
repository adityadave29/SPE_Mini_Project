pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
    }

    stages {

        stage('Run Unit Tests') {
        steps {
            sh '/opt/homebrew/bin/go version'
            sh '/opt/homebrew/bin/go mod tidy'
            sh '/opt/homebrew/bin/go test -v ./...'
        }
    }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE:latest .'
            }
        }

        stage('Push To Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                    sh 'docker push $DOCKER_IMAGE:latest'
                }
            }
        }
    }

    post {
        success {
            echo "Build and Push Successful"
        }
        failure {
            echo "Build Failed"
        }
    }
}