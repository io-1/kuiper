pipeline {
    agent any
    stages {
        stage('unit tests') {
            agent { docker { image 'golang' } }
            steps {                 
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github.com/io-1/kuiper'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/io-1/kuiper'
                sh 'cd ${GOPATH}/src/github.com/io-1/kuiper'

                sh 'go clean -testcache'
                sh 'go get -v -d ./...'
                sh 'go test -v -short --tags unit ./...'
            }
        }
        stage('deploy') {
            steps {
                echo 'Deploying..'
            }
        }
    }
}
