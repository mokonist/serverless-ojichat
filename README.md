# serverless-ojichat
[ojichat](https://github.com/greymd/ojichat) をAWS API Gateway + AWS Lambdaなサーバーレス構成にしました。

## 使い方
インフラのプロビジョニングに [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) を利用しています。
```shell script
sam build 
sam deploy --guided
```

### ライセンス

利用している greymd/ojichat はMIT Licenseで提供されています。
[https://github.com/greymd/ojichat/blob/master/LICENSE](https://github.com/greymd/ojichat/blob/master/LICENSE)

本リポジトリも [MIT License](https://github.com/mokocm/serverless-ojichat/LICENSE) となります。
