pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
    }

    triggers {
        pollSCM('* * * * *')
    }

    stages {

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