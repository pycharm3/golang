# golang
git提交代码到gitbug操作流程：
-------
1.在github创建一个仓库\<br>
2.本地下载git并安装\<br>
3.设置用户信息\<br>
$ git config --global user.name "填写用户名"\<br>
$ git config --global user.email 填写邮箱\<br>
4.查看配置信息\<br>
$ git config --list\<br>
5.git初始化在本地创建实例文件夹.git\<br>\<br>
$ git init\<br>
6.将README加入缓存\<br>
$ git add README.md\<br>
$ git status //查看是否加入缓存\<br>

7.提交打tag\<br>\<br>
$ git commit -m "first commit"
8.提交到仓库\<br>
$ git remote add origin 仓库地址\<br>
8.push到master\<br>
$ git push -u origin master\<br>

第二次提交:\<br>
$ git add 要提交的文件\<br>
$ git commit -m "打tag"\<br>
$ git push -u origin master\<br>
