
## winserver 给普通用户启动，重启IIS操作的项目

#### 服务端提供API接口，供客户端去调用

i. iis操作
>  - [GET] /iis/restart    重启iis
>  - [GET] /iis/start      启动iis
>  - [GET] /iis/stop        关闭iis

ii. 7netvfs操作

> - [GET] /7netvfs/restart   重启7netvfs服务
> - [GET] /7netvfs/start  启动7netvfs服务
> - [GET] /7netvfs/stop   停止7netvfs服务

iii. 获取权限
> - [POST] /admin/add   添加管理员
> - [POST] /admin/del   移出管理员   
 