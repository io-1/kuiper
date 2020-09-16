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
                /* sh 'make -C ${GOPATH}/${SRC_PATH} test-unit 2>&1 | go-junit-report > ${WORKSPACE}/unit-report.xml' */

                // junit plugin
                /* junit 'unit-report.xml' */

                // run tests
                sh 'echo "mode: set" > ${WORKSPACE}/coverage.out'
                sh 'go test -v -coverprofile ${WORKSPACE}/coverage.out --tags unit ${GOPATH}/${SRC_PATH}/... > ${WORKSPACE}/unit-tests.txt 2>&1'

                // archive unit tests 
                archiveArtifacts artifacts: 'unit-tests.txt' 

                sh 'cat ${WORKSPACE}/unit-tests.txt 2>&1 | go-junit-report > ${WORKSPACE}/unit-report.xml'

                // junit plugin
                junit 'unit-report.xml'


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
    }
    post {
        always {
            sh 'docker rmi $(docker images -aq) || exit 0'
        }
    }
}
