# MySQL 최신 이미지를 기반으로 설정
FROM mysql:latest

# 환경 변수 설정
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=metting
ENV MYSQL_USER=metting
ENV MYSQL_PASSWORD=metting

# 초기화 스크립트 복사
COPY ./initdb.sql /docker-entrypoint-initdb.d/

# MySQL 포트 공개
EXPOSE 3306
