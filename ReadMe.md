## 使用 kratos 和 gorm 搭建代码托管平台

### 1、项目搭建

这里需要先把 kratos 整下来

```bash
kratos new kg-git -r https://gitee.com/go-kratos/kratos-layout.git

# 这里后面的一套主要是直接New出不来，网不好

# 项目创建好后

cd kg-git
go mod tidy # 用来拉取依赖

# 试试看能不能跑

kratos run
```

### 2、快速体验 git 远程仓库

1. 初始化空的仓库

```shell
git init --bare /root/git-test/hello.git
```

2. 生成 ssh 密钥

```shell
ssh-keygen -t rsa -C "984274788@qq.com"
```

3. 将客户端生成的公钥赋值到服务器

```shell
cat ~/.ssh/id_rsa.pub

# 服务端的配置文件路径
vi ~/.ssh/authorized_keys
```

4. 操作远程仓库

```shell
# clone

git clone root@192.27.164.148:/root/git-test/hello.git

# 添加远程仓库

git remote add origin root@192.27.164.148:/root/git-test/hello.git

# 推送本地代码到远程仓库
git push -u origin master
```

### 3、系统模块

- 仓库管理
  - 仓库列表
  - 新增仓库
  - 修改仓库
  - 删除仓库
  - 授权用户
- 用户管理
  - 登录
- git 服务 ⏎
