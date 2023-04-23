# 1. 容器方式安装

## 1.1 安装规划

本实验前提已经在OCI中创建好VCN和VM，且每台vm增加200G块存储。

|序号      | 服务器名  | IP地址  | 数据和日志存储路径 | 对应盘符|
|:--------:|:---------|:--------|:---------|:---------|
|1   |hand-rabbitmq-node1| 10.0.0.226|/var/log/rabbitmq and /var/lib/rabbitmq |/dev/sdb  |
|2   |hand-rabbitmq-node2|10.0.0.208 |/var/log/rabbitmq and /var/lib/rabbitmq |/dev/sdb  |
|3   |hand-rabbitmq-node3|10.0.0.238 |/var/log/rabbitmq and /var/lib/rabbitmq |/dev/sdb  |

在对应VM的/etc/hosts中增加对应的域名解释

```text
10.0.0.226 hand-rabbitmq-node1
10.0.0.208 hand-rabbitmq-node2
10.0.0.238 hand-rabbitmq-node3
```

## 1.2 环境准备

* 服务器环境准备
  
依次在服务器：hand-rabbitmq-node1，hand-rabbitmq-node2，hand-rabbitmq-node3上执行

```bash
# 关闭selinux模式为 /etc/selinux/config  SELINUX=permissive
[root@hand-rabbitmq-node3 ~]#setenforce 0
[root@hand-rabbitmq-node3 ~]# getenforce
Permissive
```

* 创建文件系统

```bash
# 新建RabbitMQ 数据存储文件系统
[root@hand-rabbitmq-node1 ~]# pvcreate /dev/sdb
  Physical volume "/dev/sdb" successfully created.
```

```bash
[root@hand-rabbitmq-node1 ~]# vgcreate datavg /dev/sdb
  Volume group "datavg" successfully created
[root@hand-rabbitmq-node1 ~]# lvcreate -n lvrabbitmq -L  100G datavg
  Logical volume "lvrabitmq" created.
[root@hand-rabbitmq-node1 ~]# lvcreate -n lvrabbitmqlog -l 100%FREE datavg
  Logical volume "lvrabbitmqlog" created.
```

```bash
[root@hand-rabbitmq-node1 ~]# mkfs.xfs /dev/datavg/lvrabbitmq
[root@hand-rabbitmq-node1 ~]# mkfs.xfs /dev/datavg/lvrabbitmqlog
[root@hand-rabbitmq-node1 ~]# mkdir /var/log/rabbitmq
[root@hand-rabbitmq-node1 ~]# mkdir /var/lib/rabbitmq
```

```bash
[root@hand-rabbitmq-node1 ~]# lsblk -f
NAME               FSTYPE      LABEL UUID                                   MOUNTPOINT
sda                                                                         
├─sda1             vfat              6AD3-CAF6                              /boot/efi
├─sda2             xfs               670a020a-d604-4da4-a096-65c005812ef2   /boot
└─sda3             LVM2_member       fpoIDN-OdPb-RvyS-0pPP-wSpq-ECy1-2qWXNb 
  ├─ocivolume-root xfs               ac429d9c-399a-4194-8e43-c92679901401   /
  └─ocivolume-oled xfs               b4ce40f5-2ade-42e5-8371-b67fcc5e1e06   /var/oled
sdb                LVM2_member       FoLlbP-RfDc-wym6-mooe-8Mg9-SNtG-fnAyVH 
└─datavg-lvrabitmq xfs               bd6f0a53-ed8c-404b-9325-da981fee3bf1   
```

* 增加文件系统挂载点

```bash
[root@hand-rabbitmq-node1 ~]# vi /etc/fstab 
UUID=3e2833e3-01b6-4806-8914-78c0b0a424a6  /var/log/rabbitmq  xfs     defaults        0 2
UUID=4e2833e3-01b6-4806-8944-78c0b0a4446  /var/lib/rabbitmq   xfs     defaults        0 2
```

* 挂载文件系统/var/log/rabbitmq 和 /var/lib/rabbitmq

```bash
[root@hand-rabbitmq-node1 ~]# mount -a
[root@hand-rabbitmq-node1 ~]# df -h
Filesystem                    Size  Used Avail Use% Mounted on
devtmpfs                      7.6G     0  7.6G   0% /dev
tmpfs                         7.7G     0  7.7G   0% /dev/shm
tmpfs                         7.7G  8.7M  7.7G   1% /run
tmpfs                         7.7G     0  7.7G   0% /sys/fs/cgroup
/dev/mapper/ocivolume-root     36G  8.0G   28G  23% /
/dev/mapper/ocivolume-oled     10G  111M  9.9G   2% /var/oled
/dev/sda2                    1014M  324M  691M  32% /boot
/dev/sda1                     100M  5.0M   95M   6% /boot/efi
tmpfs                         1.6G     0  1.6G   0% /run/user/0
tmpfs                         1.6G     0  1.6G   0% /run/user/987
tmpfs                         1.6G     0  1.6G   0% /run/user/1000
/dev/mapper/datavg-lvrabitmqlog  100G  746M  99G   1% /var/log/rabbitmq
/dev/mapper/datavg-lvrabitmq  100G  746M  100G   1% /var/lib/rabbitmq
```

* 关闭系统防火墙

```bash
systemctl stop firewalld
```

## 1.3 安装 Docker容器引擎

```bash
sudo yum-config-manager --enable ol7_addons
sudo yum install docker-engine
sudo systemctl start docker  
sudo systemctl enable docker
sudo groupadd docker
sudo service docker restart
sudo usermod -a -G docker opc

```
## 1.4 部署RabbitMQ Container

```bash
cd /var/lib/rabbitmq
mkdir config
cd config
nano enabled_plugins
```

enabled_plugins 配置格式:

[rabbitmq_federation,rabbitmq_management,rabbitmq_mqtt,rabbitmq_shovel,rabbitmq_shovel_management,rabbitmq_stomp,rabbitmq_web_mqtt,rabbitmq_web_stomp,rabbitmq_web_stomp_examples].

运行RabbitMQ Container容器：
```bash
docker run -d \
--restart always \
--hostname rabbitmq-demo.h2lcloud.com \
-p 80:15672 -p 1883:1883 -p 443:15671 \
-e RABBITMQ_DEFAULT_USER=admin \
-e RABBITMQ_DEFAULT_PASS=Str0ngPassword \
-v /var/lib/rabbitmq/config/enabled_plugins:/etc/rabbitmq/enabled_plugins \
-v /var/lib/rabbitmq/:/var/lib/rabbitmq \
--name rabbitmq \
rabbitmq:3-management
```

 MQTT support with something like Mosquitto：

 ```bash
 # subscribe to a topic:
mosquitto_sub \
-t demo/topic \
-h rabbitmq-demo.h2lcloud.com \
-u admin \
-P Str0ngPassword 
-p 1883

# publish a message to a topic:
mosquitto_pub \
-t demo/topic \
-h rabbitmq-demo.h2lcloud.com \
-u admin \
-P Str0ngPassword \
-p 1883 \
-m "hello, world"
 ```