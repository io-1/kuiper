pipeline {
    agent any
    environment {
        SRC_PATH = 'src/github.com/io-1/kuiper'
    }
    stages {
        stage('Run Unit Tests') {
            agent { 
            docker { 
                    image 'golang' 
                    args  '-e GOCACHE=/tmp'
                } 
            }
            steps {

                // copy over the workspace files into the correct dir
                sh 'mkdir -p ${GOPATH}/${SRC_PATH}'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/${SRC_PATH}'

                // get dependencies
                sh 'make -C ${GOPATH}/${SRC_PATH} get'

                // run tests - convert to junit
                sh 'make -C ${GOPATH}/${SRC_PATH} test-unit 2>&1 | go-junit-report > ${WORKSPACE}/unit-report.xml'

                // junit plugin
                junit 'unit-report.xml'

                // creating code coverage
                sh 'echo "mode: set" > ${WORKSPACE}/coverage.out'
                sh 'go test -v -coverprofile ${WORKSPACE}/coverage.out --tags unit ${GOPATH}/${SRC_PATH}/...'

                // create coberuta report
                sh 'gocov convert ${WORKSPACE}/coverage.out | gocov-xml > ${WORKSPACE}/coverage-report.xml'
                cobertura coberturaReportFile: 'coverage-report.xml', enableNewApi: true

                // create html and archive it
                sh 'go tool cover -html ${WORKSPACE}/coverage.out -o ${WORKSPACE}/coverage.html'
                publishHTML (target: [
                  allowMissing: false,
                  alwaysLinkToLastBuild: false,
                  keepAll: true,
                  reportDir: '.',
                  reportFiles: 'coverage.html',
                  reportName: "Code Coverage Report"
                ])
            }
        }
        stage('Deploy') {
            steps {
                echo 'deploying..'
            }
        }
        stage('clean up') {
            steps {
                sh 'docker rmi $(docker images -aq) || exit 0'
            }
        }
    }
    post {
        always {
            sh 'docker rmi $(docker images -aq) || exit 0'
        }
    }
}
