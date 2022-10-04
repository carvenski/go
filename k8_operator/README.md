# K8 CRD/Operator Demo

# 创建一个报表自定义资源Report CRD
# 编写一个Report Operator来监听并触发业务逻辑

## Report CRD
name
type  finebi或fine report类型报表
author
datasource  是否配置数据源权限
tables  是否配置表权限
row_filters  是否开启行列分公司权限
colume_filters
test_env_ready 是否发布到测试环境
prod_env_ready

## Report Operator
监听到Report资源的更新事件,根据其中字段的变化,触发相应的业务逻辑.
比如,报表开发时,自动配置其数据源和表等权限,
报表开发完成后,自动同步文件并发布到测试和生产环境,
自动生成其对应的内外网service和ingress访问链接等动作.
实现BI报表分析平台的云原生自动化CICD等功能,提高分析效率.

# Operator开发
- 参考示例代码sample-controller/ 参考https://github.com/kubernetes/sample-controller
- 参考k8s/pkg/controller/目录里面的各种controller源码
- 熟练使用client-go,它就是操作k8 api的go库

# 编译sample-controller
go build sample-controller


