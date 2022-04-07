// package main

// func main() {
// 	/*
// 				官网:
// 					https://hub.docker.com
// 				国内：
// 					1.网易蜂巢：https://c.163.com/hub
// 					2.daoCloud: https://hub.daocloud.io/
// 					3.搭建私服

// 					搭建私服:
// 						1.修改/etc/docker/daemon.json文件
// 							添加:
// 								{
// 		        					"regitstry-mirrors":["https://docker-cn.com"],
// 		        					"insecure-registries":["ip:port"]
// 								}
// 						2.重启服务:
// 								systemctl daemon-reload
// 								systemctl restart docker

// 				拉取镜像到本地：(以tomcat为例)
// 					docker pull daocloud.io/library/tomcat:8.0.45
// 				查看全部本地镜像：
// 					docker images
// 				删除本地镜像：
// 					docker rmi 镜像标识(IMAGE ID 通过docker images查看)
// 				将本地的镜像导出：
// 					docker save -o ./tomcat.img fa6a
// 					(docker save -o 导出路径	镜像id)
// 				导入本地镜像文件：
// 					docker load -i tomcat.img
// 					(docker load -i 文件名)
// 				修改镜像名称:
// 					docker tag fa daocloud.io/library/tomcat:8.0.45
// 					(docker tag 镜像id 镜像名称:版本号)
// 				运行容器:
// 					1.简单操作:
// 						docker run 镜像id和镜像名称
// 					2.常用的参数:
// 						docker run -d -p 8080:8080 --name tomcat fa
// 						(docker run -d -p 宿主机端口:容器端口 --name 容器名称 镜像id或者[镜像名称:版本号])
// 					-d :代表后台运行容器
// 					-p 宿主机端口:容器端口 :为了映射当前Linux的端口和容器的端口
// 					--name 容器名称 ：指定容器的名称
// 				查看正在运行的容器：
// 					docker ps [-qa]
// 					-q ：查看容器的id
// 					-a ：查看全部容器的包括没有运行的
// 				查看容器日志
// 					docker logs -f 11
// 					-f ：可以滚动查看日志最后几行
// 				加入容器内部
// 					docker exec -it 11 bash
// 					(docker exec -it 容器id bash)
// 					退出:exit
// 				停止容器运行:
// 					docker stop 11
// 					(docker stop 容器id)
// 					docker stop $(docker ps -qa)	停止全部容器
// 				启动容器:
// 					docker start 11
// 					（docker start 容器id)
// 				删除容器：
// 					docker rm 11
// 					(docker rm 容器ID)
// 					docker rm $(docker ps -qa)	删除全部容器
// 	*/
// }
