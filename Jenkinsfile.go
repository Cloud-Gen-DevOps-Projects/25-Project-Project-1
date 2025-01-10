pipeline{
	agent any 
	tools{
		jdk "jdk17"
		maven "maven3"
	}
	environment{
		SCANNER_HOME= tool "sonar-task"
	}
	stages{
		stage("Code Checkout"){
			steps{
				git branch: 'main', url: 'https://github.com/Cloud-Gen-DevOps-Projects/25-Projects-Project-1.git'
			}
		}
		
		stage("Code Analysis"){
			steps{
				withSonarQubeEnv('sonar-token') {
					sh "mvn clean install sonar:sonar"
 					}
			}
		}
		
		
		
		
		
		
		
	}
}

				
