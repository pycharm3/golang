# golang
reflect.go-golang反射最佳实践
-------
* 声明一个struct并在main函数中实例化传入函数中
* 获取传入参数的TypeOf和ValueOf
* 获取ValueOf的Kind进行判断如果是非reflect.Value类型则退出程序
* 再通过rValue.NumMethod()获取到该value的字段个数
* 遍历该value
* `rTyp.Field(i).Tag.Get("json")`i代表第i个字段，如果字段tag的key匹配到json则返回该字段tag否则返回为nil
* `rValue.NumMethod()` 可用来获取绑定value.type的方法数,注意方法排序是根据方法名的ASCII码排序
* `rValue.Method(1).Call(nil)` 1代表调用第一个方法，nil代表传入参数为空，如果方法需要参数则在此传入参数，注意参数为value.[]slice类型
* `var reslice []reflect.Value` 声明一个reflect.Value的slice可以添加参数传给需要的方法
* 额代码有注释不想写了

golang反射最佳实践总结
-------
* 反射顾名思义能够反向操作经典案例golang中的testing单元测试框架，函数名要求Test开头不需要用户调度则底层能够自动运行，就是用到了反射<br>
* 代码编译检测到规定格式开头的函数就自动加载执行，然后对执行函数进行反射操作获取其对应的Type和Value然后再使用TypeOf和ValueOf下的方法对其进一步操作，获取其字段及方法等各个元素，底层调用执行，所以不需要我们一个个去进行调度执行反射底层就可以做好<br>
* 反射的这些特性特别适合进行框架的封装，用户只需要写好需要的上下文资源由反射进行调度大大增强了可操作性<br>

git提交代码到gitbug操作流程：
-------
1.在github创建一个仓库<br>
2.本地下载git并安装<br>
3.设置用户信息<br>
$ git config --global user.name "填写用户名"<br>
$ git config --global user.email 填写邮箱<br>
4.查看配置信息<br>
$ git config --list<br>
5.git初始化在本地创建实例文件夹.git<br>
$ git init<br>
6.将README加入缓存<br>
$ git add README.md<br>
$ git status //查看是否加入缓存<br>

7.提交打tag<br>
$ git commit -m "first commit"<br>
8.提交到仓库<br>
$ git remote add origin 仓库地址<br>
8.push到master<br>
$ git push -u origin master<br>

第二次提交:<br>
$ git add 要提交的文件<br>
$ git commit -m "打tag"<br>
$ git push -u origin master<br>
