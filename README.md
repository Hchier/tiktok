# tiktok

## 项目介绍

实现极简版抖音。每一个应用都是从基础版本逐渐发展迭代过来的，希望同学们能通过实现一个极简版的抖音，来切实实践课程中学到的知识点，如Go语言编程，常用框架、数据库、对象存储等内容，同时对开发工作有更多的深入了解与认识，长远讲能对大家的个人技术成长或视野有启发。



## 接口

[接口文档在线分享 - Apifox](https://www.apifox.cn/apidoc/shared-099e9908-7f5b-4f15-b4a1-c2816811dbb9)



## 部署

### 环境

1. redis

2. mysql

    数据库见tiktok.sql

3. nginx

    ```nginx
    server {
        listen       8500;
        location / {
        # 由于StaticResourcePathPrefix=/projects_related/static/tiktok/，
        # 且 StaticResourceUrlPrefix=http://120.77.13.126:8500/tiktok/
        # 故 http://120.77.13.126:8500/tiktok/ --> /projects_related/static/tiktok/
        root   "/projects_related/static";
        }
    }
    ```

4. fmmgeg

    1. 下载 https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
    2. 解压tar -zxvf ffmpeg-release-amd64-static.tar.xz
    3. 环境变量 export PATH=$PATH:/www/server/ffmpeg/ffmpeg-5.1.1-amd64-static



### 问题

部署时发现了个问题，把`HostPorts`改为`120.77.13.126:8080`时，项目启动报错，改为`127.0.0.1:8080`后可以正常启动，但是外界又无法访问，于是只能用nginx做了一层反向代理。

```nginx
server{
        listen 80;
        server_name tiktok.hchier.xyz;
        location / {
		proxy_pass http://127.0.0.1:8080/;
	}
}
```



## 使用

1. 安装apk并打开。
2. 双击“我”，配置服务端前缀地址：http://tiktok.hchier.xyz/
