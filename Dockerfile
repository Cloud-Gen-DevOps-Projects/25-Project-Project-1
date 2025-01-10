FROM tomcat:9.0-jdk11
LABEL maintainer="ravindra.devops@gmail.com"
COPY ./register-form.war /usr/local/tomcat/webapps/register-form.war
EXPOSE 8080
CMD ["catalina.sh", "run"]

