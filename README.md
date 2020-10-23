# alipay_f2fpay
## 阿里云 当面付2.0 DEMO GO

基于 github.com/smartwalle/alipay/v3 ，采样ECHO框架，把官方的当面付demo移植到GO平台<br>
使用步骤:<br>
1.先在支付宝开放平台申请开通当面付2.0，按官方教程配置密钥<br>
2.修改/f2fpay/conf/server.conf文件，填入alipay_public_key，private_key，appid<br>
3.运行f2fpay.exe<br>
4.访问http://127.0.0.1 进行测试<br>

![image](https://github.com/bestyun-xyz/alipay_f2fpay/blob/main//screenshot/index.png)
![image](https://github.com/bestyun-xyz/alipay_f2fpay/blob/main//screenshot/barpay.png)
![image](https://github.com/bestyun-xyz/alipay_f2fpay/blob/main//screenshot/qrpay.png)
![image](https://github.com/bestyun-xyz/alipay_f2fpay/blob/main//screenshot/query.png)
![image](https://github.com/bestyun-xyz/alipay_f2fpay/blob/main//screenshot/refund.png)

<br>
### 2020.10.23 测试所有功能正常
