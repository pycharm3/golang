# golang
git提交代码到gitbug操作流程：
1.在github创建一个仓库
2.本地下载git并安装
3.设置用户信息
$ git config --global user.name "填写用户名"
$ git config --global user.email 填写邮箱
4.查看配置信息
$ git config --list
5.git初始化在本地创建实例文件夹.git
$ git init
6.将README加入缓存
$ git add README.md
$ git status //查看是否加入缓存

7.提交打tag
$ git commit -m "first commit"
8.提交到仓库
$ git remote add origin 仓库地址
8.push到master
$ git push -u origin master

第二次提交:
$ git add 要提交的文件
$ git commit -m "打tag"
$ git push -u origin master