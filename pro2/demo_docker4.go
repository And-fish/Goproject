// package main

// func main() {
// 	/*
// 		自建docker镜像:
// 			1.新建test目录，把demo复制到test文件夹
// 				cp -r demo test
// 			2.新建一个dockerfile文件,在里面添加:
// 				FROM daocloud.io/library/tomcat:8.0.45
// 				#RUN jar -cvfM0 demo.war demo/*		#这行在外面手动输出了，因为docker找不到jar
// 				ADD demo.war /usr/local/tomcat/webapps
// 			3.自建镜像:
// 				docker build -t tomcat:1.0 .
// 			4.检查容器:
// 				docker images
// 				docker run -d -p 8080:8080 --name TC 45
// 				docker ps
// 				访问http://192.168.25.128:8080/demo/demo/	输出了test decker

// 		Docker-Compose：
// 			之前运行一个镜像，需要添加大量的参数。可以通过Docker-Compose编写这些参数
// 			Docker-compose可以帮助批量的管理容器，只需要通过以恶搞docker-compose.yml文件去维护即可

// 		安装docker-compose:
// 			1.下载https://github.com/docker/compose/releases/download/1.24.1/docker-compose-Linux-x86_64
// 			2.把下载好的文件放到linux里面
// 			3.修改文件名称为docker-compose将文件移动到/usr/local/bin下面并给文件一个777权限:
// 				mv docker-compose-Linux-x86_64 docker-compose
// 				mv docker-compose /usr/local/bin/
// 				chmod 777 docker-compose
// 			3.5.修改环境变量:
// 				vim /etc/profile
// 				加上export PATH=路径:$PATH

// 		新建docker-compose.yml:
// 			D:\Users\Gamisuu\Desktop\GOproject\pro2\docker-compose.yml

// 		基于docker-compose.yml启动管理的容器:
// 			docker-compose up -d
// 		开启/关闭/重启的容器:
// 			docker-compose start/stop/restart
// 		关闭并删除容器:
// 			docker-compose down
// 		查看由docker-compose管理的容器:
// 			docker-compose ps
// 		查看日志:
// 			docker-compose logs -f

// 		基于dockerfile和docker-compose管理容器
// 		修改docker-compose.yml为
// 			D:\Users\Gamisuu\Desktop\GOproject\pro2\docker-compose.yml

// 	*/

// }
