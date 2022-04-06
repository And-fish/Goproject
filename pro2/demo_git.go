package main

func main() {
	/*
			git本地有四个工作区域:
				工作目录(Workpace):平时存放项目代码的地方
				暂存区(Index/Stage):用于临时存放你的改动，事实上它只是以一个文件，保存即将提交到文件列表信息
				资源库(Repository或Git Directory):安全存放数据的位置，这里面有你提交到所有版本的数据。其中HEAD指向最新放入仓库的版本
				git仓库(Remote Directory):远程仓库，托管代码的服务器，可以简单的认为是你项目组中的一台电脑用于远程数据交换

			workspace ---add---> index ---commit---> repository ---push---> remote	从本地工作目录上传数据到远程仓库
			repository ---checkout---> workspace
			remote ---clone--->repository
			remote ---pull---> workspace	从远处仓库拉取数据到本地

			git的工作流程:
				1.在工作目录中添加、修改文件;
				2.将需要进行版本管理的文件放入暂存区域;
				3.将暂存区域的文件提交到git仓库
			因此git管理的文件有三种状态:已修改，已暂存，已提交

			本地文件的四个状态:
				刚新建是untracked(未跟踪)状态，没有加入到git库，不参与版本控制，通过git add变成staged
				文件已经入库，未修改，即版本库中的文件快照内容与文件夹中完全一致。
				修改后变成了modified状态，在提交就变成了staged。或者把库中的remove掉，就变成了untracked状态
				staged状态commit后变成unmodified状态


			git branch 	查看本地所有分支
			git	status	查看当前状态
			git	commit	提交
			git	branch -a	查看所有分支
		!!! git commit -am "init"	提交并注释
			git remote add origin git@192.168.1.119:ndshow	远程添加
			git	push origin master	将文件推送到服务器上
			git	remote show origin	显示远程库origin里面的资源
			git	push origin master:hb-dev	将本地库与服务器上的库进行关联
			git checkout --track origin/dev	切换到远程dev分支
			git brach -D master develop	删除本地库的develop
			git	checkout -b dev	建立一个新的本地分支dev
			git	merge origin/dev 将分支dev和当前分支合并
			git checkout dev	切换到本地分支dev
			git remote show	查看远程库
			git	add .
			git rm 文件名	从git中删除某个文件
			git clone git://github.com/schacon/grit.git 从服务器上将代码拉下来
			git	config --list	看所有用户
			git	ls-files	看已经提交的
			git commit -a 提交当前repos的所有改变
			git	add [file name]	添加以恶搞文件到git index
			git	commit -v	可以看commit的差异
			git	commit -m	添加commit信息，注释
			git	commit -a	添加到index再commit
			git	commit -a -v	一般提交命令
			git	rm -f	强行移除修改后的文件
			git diff --cached	查看尚未提交的更新
			git stash push	将文件给push到一个临时空间
			git	stash pop	将文件从临时空间pop下来
			git remote add origin git@github.com:username/hello-woeld.git
			git push origin master	将本地项目给提交到服务器
			git pull	同步本地和服务器
			git push (远程仓库名) (分支名)	将本地分支推送到服务器上去
			git push origin serverfix:awesomebranch
			git fetch	相当于是从远程获取最新版本到本地，不会自动merge


			mkdir directory
			cd directory
			git init	本地初始化
			touch README
			git add README	添加文件
			git commit -m 'first commit'
			git remote add origin git@guthub.com:username/directory.git


			git用户配置:
				1.单独配置:
					git config username
					git config useremail
				2.全局配置:
					git config --global username
					git config --global useremail

			git初始化:
				git init
			.git目录会为你的项目存储所有历史和元信息的目录，这些对象指向不同的分支
			.git目录下面的文件:
				1.HEAD文件：是我们常说的haed指针，指向了当前的分支
				2.config：保存当前仓库的配置信息
				3.description：项目的描述信息
				4.hooks/：系统默认狗子脚本目录
				5.index/：索引文件，暂存区
				6.logs/：各个refs的历史信息
				7.objects/：本地仓库的所有对象(commits，trees，blobs，tags)
				8.refs/：标识项目里的每个分支指向了那个提交(commit)

			git状态:
				git status
			三种提示信息:
				1.Changes to be committed：表示已经从工作区add到暂存区file(文件或文件夹)
				2.Changes not staged for commit：表示工作区、暂时区都存在的file(文件或文件夹)，在工作区进行修改或删除，但是没有add到暂存区
				3.Untracked files：表示只在工作区有的file，也就是在暂时区没有该file。

			git将代码添加到暂存区:
				git add [file1] [file2]
				git add [dir]
				git add .
			git add . 和git add -u 和 git add -A的区别:
				git add . :会监控工作区的状态树，使用这个命令会把工作室所有的变化提交到暂存区，包括文件修改和新文件，但不包括被删除的文件。
				git add -u :只监控已经被add的文件，它会将被修改的文件提交到暂存区。add -u不会提交新文件
				git add -A :是上面两个功能的集合

			git提交代码:
				git commit [file1] [file2] -m [msessage]
			直接提交，不需要git add添加到暂存区:
				git commit -am [message]
			修改并提交:
				1.修改代码文件，查看状态(git status)：
					现在是modified状态
				2.再次添加到暂存区(git add filename)，再次查看状态(git status)
				3.查看版本日志(git reflog)

			git项目版本切换：
				先通过git reflog或者git log查看版本信息
				再用git reset --hard 版本号xxx
			查看.git下面的refs目录:
				.git/refs/heads/master

			git分支简介:
				1.当为项目增加新特性，但很可能会影响到当前可正常使用的代码。
				  与其将特性加入到其他人正在使用的master分支，更好的方法是在仓库的其他分支中变更代码
				2.更重要的是git设计是用于协同开发。如果多人使用master分支，可能会引发混乱。
				  使用分支可以让你核验他人的贡献并选择适合的加入到项目中
			分支的好处:
				1.同时推进多个功能的开发，提高开发效率
				2.各个分支在开发的过程中，如果一个分支开发失败，不会对其他分支有任何影响

			git分支常用命令：
				git branch -v
					查看分支，"*"号代表当前在那个分支
				git branch BranchName
					新建分支
				git branch -d BranchName
					删除分支
				git branch -m OldBranchName NewBranchName
					修改分支
				git checkout BranchName
					切换分支
				git merge BranchName
					合并分支
			git合并分支引起的冲突：
				合并分支时，两个分支在同一个代码的同一位置都有修改，git无法判断如何去留
			实例：
				vim FileName	在主分支上修改文件
				git commit -am 'modifiy by master' FileName	主分支提交文件
				git checkout branch1	切换分支
				vim FileName	在分支上修改同一个文件
				git commit -am 'mobifiy by branch1' FileName	副分支上提交文件
				git checkout master	切换到主分支
			如何解决：
				vim FileName	手动打开文件，手动选择需要保留的内容
				git commit -am 'merge commit'	手动提交选择后的内容


	*/
}
