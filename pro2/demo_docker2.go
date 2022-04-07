// package main

// func main() {
// 	/*
// 		一、准备一个web项目：
// 			1.创建一个文件夹，例如demo
// 			2.创建一个文件夹WEB-INF
// 			3.在WEB-INF中创建一个文件web.xml
// 			4.在web.xml中添加:
// 				<web-app xmlns="https://xmlns.jcp.org/xml/ns/javaee"
// 						xmlns:xsi="https://www.w3.org/2001/XMLSchema-instance"
// 						xsi:schemaLocation="https://xmlns.jcp.org/xml/ns/jvaee
// 						https://xmlns.jcp.org/xml/ns/javaee/web-app_3_1.xsd"
// 						version="3.1"
// 						metadata-complete="true">
// 				</web-app>
// 			5.创建一个index.html，添加"hello world"

// 		二、创建MySQL容器
// 			1.创建MySQL镜像:
// 				docker pull daocloud.io/library/mysql:5.7.5-m15
// 			2.查看MySQL镜像:
// 				docker images
// 			3.创建MySQL容器:
// 				docker run -d -p 3303:3306 -e MYSQL_ROOT_PASSWORD=123456 --name mysql 644
// 			4.查看容器:
// 				docker ps

// 		三、创建Tomcat容器
// 			1.创建Tomcat镜像:
// 				docker pull daocloud.io/library/tomcat:8.0.45
// 			2.查看Tomcat镜像:
// 				docker images
// 			3.创建Tomcat容器:
// 				docker run -d -p 8080:8080 --name tomcat fa
// 			4.查看容器:
// 				docker ps
// 		四、将项目部署到Tomcat
// 			1.将web项目拷贝到Tomcat的webapps目录下面:
// 				docker cp demo 容器id:/user/local/tomcat/webapps
// 			2.进入容器目录:
// 				docker exec -it 容器id bash
// 			3.查看容器日志:
// 				docker logs -f 容器id

// 		五、Docker数据卷
// 			为了部署项目，需要使用cp命令将宿主机内的demo文件复制到容器内部
// 			1.数据卷:
// 				将宿主机的一个目录，映射到容器的一个目录，可以在宿主机中操作目录中的内容，那么容器内部映射的文件也会跟着一起改变
// 			2.创建数据集:
// 				docker volume create Docker_Data
// 				(docker volume create 数据卷name)
// 				默认存放到 /var/lib/docker/volumes/数据卷name/_data 目录
// 			3.查看数据卷的详细信息:
// 				docker volume inspect Docker_Data
// 				(docker volume inspect 数据卷name)
// 			4.查看所有的数据卷:
// 				docker volume ls
// 			5.删除数据卷:
// 				docker volume rm Docker_Data/
// 				(docker volume rm 数据卷name)
// 			6.应用数据卷:
// 				docker run -v 数据卷名称:容器内部路径 镜像id
// 			或者docker run -v 数据卷路径:容器内部路径 镜像id
// 			7.实例，将/root映射到tomcat容器中的webapps:
// 				docker run -d -p 8080:8080 --name tomcat -v /root:/usr/local/tomcat/webapps fa

// 	*/
// }
