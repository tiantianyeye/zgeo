# 国家与地区查询服务
## 导入与使用
1.启动时需要调用 LoadGetLiteCityDB 方法加载对应的 GeoLite2-City.mmdb 数据库
2.使用时将符合geo规则的IP地址传入 GetCityInfoByGeoIp 方法
3.方法返回，国家，地区，以及可能的错误信息

## 配置依赖
对应的 GeoLite2-City.mmdb 数据库已经被加载到对应的docker镜像中，镜像以 alpine:3.7 版本的linux系统作为基础，并将 GeoLite2-City.mmdb 数据库添加到linux系统的/app目录下