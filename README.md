# 自动化部署github项目
### 初衷
  写这个项目的项目就是做一个在个人服务下的部署自动部署后端项目的一个东西，因为公司中也有一些自动话的ci/cd流程，可以push完代码就完成项目的部署，采用了jenkins加k8s加一些东西，比较复杂
  
  个人的项目的就是越简单越好，在网上看了一下github的actions，但是呢我觉得从权限的角度来看一个项目的部署不应该是这样的，不应该ssh去链接服务器，而应该是github告诉服务器去拉去代码来部署
### 项目逻辑
有点简单
1. github在发布一个版本后，发送一个created的webhook到服务器（webhook最好不要太多，webhook上可利用东西不多）
   ![image-20230408223244877](https://zlx-1259122459.cos.ap-guangzhou.myqcloud.com/image-20230408223244877.png)
2. 系统接收到webhook会的得到项目地址和项目名字还有版本号，拉去代码
3. 执行docker-compose up -d指令

### 项目条件
1. docker
2. docker-compose
3. 项目dockerfile可以build项目
4. docker-compose.yaml可以管理项目
5. build.sh需要复制到 quck项目下，清理原来docker和部署新docker都是build.sh执行，需要有执行权限
  


### 缺点
1. 无法知道报错，无法知道具体有没有成功