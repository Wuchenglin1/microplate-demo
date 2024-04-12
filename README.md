##### 1.部署并运行Shifu

- 安装`Shifudemo`

![image-20240412230731017](G:/typoraImg/README/image-20240412230731017.png)

- 安装完成

![image-20240412231029698](G:/typoraImg/README/image-20240412231029698.png)

- 确认`Shifudemo`运行

![image-20240412231115426](G:/typoraImg/README/image-20240412231115426.png)



##### 2. 运行一个酶标仪的数字孪生

- 先运行一个`nginx`

![image-20240413021659606](G:/typoraImg/README/image-20240413021659606.png)

- 启动酶标仪的数字孪生：

![image-20240413021504756](G:/typoraImg/README/image-20240413021504756.png)

- 查看：

![image-20240413021741940](G:/typoraImg/README/image-20240413021741940.png)

- 交互：

![image-20240413021815357](G:/typoraImg/README/image-20240413021815357.png)



##### 3. 编写一个Go应用

1. 定期轮询获取酶标仪的/get_measurement接口，并将返回值平均后打印出来，轮询时间可自定义
2. Go的应用需要容器化
3. Go的应用需要运行在Shifu的k8s集群当中
4. 最终通过kubectl logs命令可以查看打印的值



- 创建查询代码，生成`go.mod`，写好`Dockerfile`(详情在git提交的仓库中)

- 创建应用

```bash
root@vultr:~/GoProjects/src# docker build --tag plate-reader-detector:v0.0.3 .
[+] Building 45.4s (13/13) FINISHED                                                                                            docker:default
 => [internal] load build definition from Dockerfile                                                                                     0.0s
 => => transferring dockerfile: 213B                                                                                                     0.0s
 => resolve image config for docker-image://docker.io/docker/dockerfile:1                                                                0.6s
 => CACHED docker-image://docker.io/docker/dockerfile:1@sha256:dbbd5e059e8a07ff7ea6233b213b36aa516b4c53c645f1817a4dd18b83cbea56          0.0s
 => [internal] load metadata for docker.io/library/golang:1.21-alpine                                                                    0.5s
 => [internal] load .dockerignore                                                                                                        0.0s
 => => transferring context: 2B                                                                                                          0.0s
 => [1/6] FROM docker.io/library/golang:1.21-alpine@sha256:ed8ce6c22dd111631c062218989d17ab4b46b503cbe9a9cfce1517836e65298a              0.0s
 => [internal] load build context                                                                                                        0.0s
 => => transferring context: 705B                                                                                                        0.0s
 => CACHED [2/6] WORKDIR /app                                                                                                            0.0s
 => CACHED [3/6] COPY go.mod ./                                                                                                          0.0s
 => CACHED [4/6] RUN go mod download                                                                                                     0.0s
 => [5/6] COPY *.go ./                                                                                                                   0.0s
 => [6/6] RUN go build -o /main                                                                                                         43.1s
 => exporting to image                                                                                                                   0.7s
 => => exporting layers                                                                                                                  0.7s
 => => writing image sha256:1f6b8e9bbbdb1b430bb4c6b630cb08f45a0b5ee6c6a9d49365f852930633985f                                             0.0s
 => => naming to docker.io/library/plate-reader-detector:v0.0.3                                                                          0.0s
```



- 加载应用镜像到`kind`中

```bash
root@vultr:~/GoProjects/src# kind load docker-image plate-reader-detector:v0.0.3
Image: "plate-reader-detector:v0.0.3" with ID "sha256:1f6b8e9bbbdb1b430bb4c6b630cb08f45a0b5ee6c6a9d49365f852930633985f" not yet present on node "kind-control-plane", loading...
```



- 运行容器 Pod

```bash
root@vultr:~/GoProjects/src# kubectl run plate-reader-detector --image=plate-reader-detector:v0.0.3 -n deviceshifu
pod/plate-reader-detector created
root@vultr:~/GoProjects/src# 
```

- 检查应用输出

```bash
root@vultr:~/shifudemo/shifudemos# kubectl logs -n deviceshifu plate-reader-detector -f
 总共有97组数据，平均数为135.970
  总共有194组数据，平均数为125.720
  总共有291组数据，平均数为128.620
  总共有388组数据，平均数为139.230
  总共有485组数据，平均数为136.760
  总共有582组数据，平均数为115.550
  总共有679组数据，平均数为142.630
  总共有776组数据，平均数为127.680
  总共有873组数据，平均数为127.640
```



