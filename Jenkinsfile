pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "adityadave29/calculator"
        DOCKER_CONFIG = "${WORKSPACE}/.docker"
        DOCKER_HOST = "unix:///Users/adityadave/.colima/default/docker.sock"
        PATH = "/usr/bin:/bin:/usr/sbin:/sbin:/opt/homebrew/bin"
    }

    stages {

        stage('Checkout Source Code') {
            steps {
                git url: 'https://github.com/adityadave29/SPE_Mini_Project.git', branch: 'master'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh '''
                go mod tidy
                go test -v ./...
                '''
            }
        }

        stage('Start Docker Environment') {
            steps {
                sh 'colima start || true'
            }
        }

        stage('Docker Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'docker-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    '''
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                docker build -t $DOCKER_IMAGE:$BUILD_NUMBER .
                docker tag $DOCKER_IMAGE:$BUILD_NUMBER $DOCKER_IMAGE:latest
                '''
            }
        }

        stage('Push To Docker Hub') {
            steps {
                sh '''
                docker push $DOCKER_IMAGE:$BUILD_NUMBER
                docker push $DOCKER_IMAGE:latest
                '''
            }
        }

        stage('Deploy with Ansible') {
            steps {
                sh '''
                ansible-playbook -i inventory.ini deploy.yml
                '''
            }
        }
    }

    post {

        success {
            emailext(
                subject: "Jenkins Build SUCCESS: ${JOB_NAME} #${BUILD_NUMBER}",
                body: """
                Build Successful!

                Job Name: ${JOB_NAME}
                Build Number: ${BUILD_NUMBER}
                Build URL: ${BUILD_URL}

                Docker Image: ${DOCKER_IMAGE}:${BUILD_NUMBER}
                """,
                to: "daveaditya2004@gmail.com",
                from: "daveaditya2004@gmail.com"
            )
        }

        failure {
            emailext(
                subject: " Jenkins Build FAILED: ${JOB_NAME} #${BUILD_NUMBER}",
                body: """
                Build Failed!

                Job Name: ${JOB_NAME}
                Build Number: ${BUILD_NUMBER}
                Build URL: ${BUILD_URL}

                Please check Jenkins logs.
                """,
                to: "daveaditya2004@gmail.com"
            )
        }

        always {
            echo "Build Finished at: ${new Date()}"
        }
    }
}