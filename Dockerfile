FROM tomcat:9.0-jdk17
LABEL maintainer="ravindra.devops@gmail.com"
copy ./*.war /usr/local/tomcat/webapps/
EXPOSE 8080
CMD ["catalina.sh", "run"]

