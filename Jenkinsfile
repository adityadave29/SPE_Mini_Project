pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
    }

    triggers {
        pollSCM('* * * * *')  
    }

    stages {

        stage('Checkout Code') {
            steps {
                git branch: 'master',
                    url: 'https://github.com/adityadave29/SPE_Mini_Project.git'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go mod tidy'
                sh 'go test -v ./...'
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
                    usernameVariable: 'adityadave29',
                    passwordVariable: 'Dave@aditya29'
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