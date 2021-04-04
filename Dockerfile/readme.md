#### 编译
docker build -f Dockerfile/Dockerfile -t redis-like-image .
#### 运行
docker run -itd -p 6379:6379 --name r-l  redis-like-image:latest