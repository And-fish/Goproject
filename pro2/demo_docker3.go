// package main

// func main() {
// 	/*
// 		什么时候Dokerfile?
// 			Dockerfile是一个包含用于组合映像的命令的文本文档。
// 			可以使用在命令行中调用的任何命令。
// 			Docker通过读取Dockerfile中的指令自动生成映像

// 			docker build命令用于从Dockerfile构建映像。
// 			(docker bulid -f Dockerfile路径)

// 		Dockerfile的剧本结构:
// 			1.基础镜像信息
// 			2.维护者信息
// 			3.镜像操作指令
// 			4.容器启动时执行指令

// 		Dockerfile常用指令:
// 			1.FROM:指定基础镜像，必须为第一个命令
// 				FROM mysql:5.6
// 				(	FROM <image>
// 					FROM <image>:<tag>
// 					FROM <image>@<digest>)	tag和digest时可选的，如果不使用，会使用latest版本的基础镜像
// 			2.MAINTAINER:维护者信息
// 				MAINTAINER xxh <1626245483@qq.com>
// 			or	MAINTAINER 1626245483@qq.com
// 			or	MAINTAINER xxh
// 				(MAINTAINER <name>)
// 			3.RUN:构建镜像时执行的命令
// 				RUN用于在镜像容器中执行命令，其中有以下两种命令执行方法
// 				RUN指令创建的中间镜像会被缓存，并会在下次构件中使用。如果不想使用这些缓存，可以在docker build时使用--no-cache参数，如:docker build --no-cache
// 					shell格式:
// 						RUN <command>
// 					exec格式:
// 						RUN ["executable","param1","param2"]
// 						示例:
// 							RUN apk update
// 							RUN ["/etc/execfile","arg1","arg2"]
// 			4.ADD:将本地文件添加到容器中，tar类型文件会自动解压，可以访问网络资源，类似于wget
// 					ADD hom* /mydir/	#添加所有以"hom"开头的文件
// 					ADD hom?.txt /mydir/	#?替代以恶搞单字符，例如hemo.txt
// 					ADD test relativeDir/	#添加"test"到`WORKDIR`/relativeDir/
// 					ADD test /absoluteDir/	#添加"test"到/absoluteDir/
// 			5.COPY:类似于ADD但是不会自动解压文件，也不能访问网络
// 			6.CMD:构建容器后调用，也就是在容器启动时才进行调用
// 				CMD echo "test"|wc -m
// 				CMD ["/usr/bin/wc","--help"]
// 				(	CMD ["executable","param1","param2"]
// 					CMD	command param1 param2)
// 			7.ENTRYPOINT:配置容器，使其可执行化，配额和CMD可省去"application"，只使用参数
// 				ENTRYPOINT ["EXECUTABLE","param1","param2"]
// 				ENTRYPOINT	command param1 param2
// 				示例:
// 					FROM ubuntu
// 					ENTRYPOINT ["top","-b"]
// 					CMD ["-c"]
// 			8.LABEL:用于为镜像添加元数据
// 				LABEL version="1.0" description="这是一个web服务器" by="xxh"
// 				(LABEL <key>=<value> <key>=<value>)
// 			9.ENV:设置环境变量
// 				ENV myname xxh
// 				ENV myname=xxh myemail="1626245483@qq.com"
// 				(	ENV <key> <value>
// 					ENV <key>=<value>)
// 			10.EXPOSE:指定于外界交互的端口
// 				EXPOSE并不会让容器的端口访问到主机。要是其可访问，需要在docker run运行容器时通过-p来发布这些端口，或来发布expose导出的所有端口
// 				EXPOSE 80 443
// 				(EXPOSE <port> [<port>])
// 			11.VOLUME:用于指定数据卷
// 				VOLUME ["/data"]
// 			12.WORKDIR:工作目录，类似于CD命令
// 					WORKDIR /A	(此时工作目录为/A)
// 					WORKDIR B	(此时工作目录为/A/B)
// 			13.USER:指定运行容器的用户名
// 			14.AGR:用于指定床底给构建运行时的变量
// 				AGR name="xxh"
// 			15.ONBUILD:用于设置镜像触发器
// 				ONBUILD ADD . /app/src
// 				(ONBUILD [INSTRUCTION])
// 	*/
// }
