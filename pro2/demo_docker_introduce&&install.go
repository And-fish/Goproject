// package main

// func main() {
// 	/*
// 		企业中存在一些问题:
// 			1.软件更新发布及部署抵消，过程烦琐需要人工介入
// 			2.环境一致性难以保证
// 			3.不同环境之间迁移成本过高

// 		Docker（Dock Worker 码头工人）是一种运行与Linux和Windows上的软件，用于创建、管理和编排容器

// 		Docker能解决的问题：
// 			1.使用简单，
// 				开发的角度来看可以分成三个部分:构建、运输、运行。其中关键环节是构建，即打包镜像文件。
// 				从测试和运维的角度来看就是两个部分：复制和运行。有了这个镜像文件，想复制到哪里运行都可以，和平台无关
// 			2.Docker隔离除了独立的运行空间，不会和其他应用争用系统资源，不需要考虑应用之间的相互影响
// 			3.因为在构建镜像文件的时候就处理好了对于系统的依赖，所以对于测试和运维而言，工作量减少了
// 			4.Docker为开发之提供了一种开发环境的管理方法，帮助测试人员保证环境的同步，为运维提供了可移植的标准化部署流程
// 	*/
// 	/*
// 		uname -a	保证内核在3.8以上
// 		yum update	更新yum包
// 		yum -y install gcc -y
// 		yum -y install gcc-c++ -y	这两个是提供device mapper驱动依赖
// 		yum install -y yum-utils device-mapper-persistent-data	yum-util提供了yum-config-manage功能
// 		yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo	设置docker-ce的yum源
// 		yum install docker-ce	安装
// 		systemctl start docker	启动Docker
// 		systemctl enable docker	加入开机启动
// 		docker version	验证安装
// 		测试:
// 			docker pull hello-world
// 			docker images
// 			docker run hello-world

// 		删除:
// 			yum list installed grep
// 			yum remove **** -y

// 	*/
// }
