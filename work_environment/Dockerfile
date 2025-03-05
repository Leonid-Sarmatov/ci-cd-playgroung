FROM ubuntu:latest

WORKDIR /my_work_directory

# Установка зависимостей
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y ca-certificates 
RUN apt-get install -y git 
RUN apt-get install -y docker.io 
RUN apt-get install -y gnupg2

# Добавление репозитория GitLab
RUN curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | bash

# Установка GitLab Runner
RUN apt install -y gitlab-runner

# Копирование скрипта запуска GitLab Runner
COPY . .

# Запуск GitLab Runner
ENTRYPOINT ["sh", "./entrypoint.sh"]

#CMD ["gitlab-runner", "run"]

# Регистрация раннера (не рекомендуется в Dockerfile!) GR1348941Epajxbqb9LAxAwumY4cM glrt-t3_maf17fN5zoq3hkDhag6G
#CMD ["gitlab-runner", "register", "--url", "http://localhost", "--token", "GR1348941Epajxbqb9LAxAwumY4cM"]
#CMD ["gitlab-runner", "run"]
#ENTRYPOINT ["cat", "/dev/zero"]