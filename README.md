# CI/CD

## サンプルサーバーの動作方法

```shell
$ cd lambda
$ env GOOS=linux GOARCH=amd64 go build -o "main" -tags netgo -ldflags='-s -w -extldflags "-static"' .
```

## コンテナ
### コンテナイメージのビルド

```shell
$ docker build -t wscicd-test . 
```

### コンテナ起動  

```shell
$ docker run -p 8080:8080 -it wscicd-test:latest
```

### コンテナイメージのプッシュ

```shell
aws ecr get-login-password --region ap-northeast-1 --profile=wscicd | docker login --username AWS --password-stdin 932657704375.dkr.ecr.ap-northeast-1.amazonaws.com
docker tag wscicd-test:latest 932657704375.dkr.ecr.ap-northeast-1.amazonaws.com/student_o_repository:latest
docker push 932657704375.dkr.ecr.ap-northeast-1.amazonaws.com/student_o_repository:latest
```
