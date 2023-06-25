```
 # 版號更新時務必使用 1.0.1 三個數字組成
 # 安裝
 go install github.com/gps-gaming-gs/goleaf@latest 
 
 # 使用
 ### 建立Proto範本
 goleaf rpc -o slot001.proto 
 
 ### 透過slot001.proto產生 leafServer
 goleaf rpc protoc slot001.proto --zrpc_out=. --home template
 
 ### 或透過slot001.proto產生 leafServer && cocos project(範本為3.7.2)
 goleaf rpc protoc slot001.proto --zrpc_out=. --home template --cocos=CocosProjectName
 ### cocos專案中的Makefile可用來將proto編譯成typescript檔，第一步：make env、第二步：make grpc，查看scripts/models
 
 ### 產生Controller/View的模板
 goleaf cocos --controller LobbyUI
 
  ### 產生Proxy的模板
 goleaf cocos --proxy Chat
 
 # 安裝依賴 
 go mod tidy 
 
 # 範本
 goleaf template init --home $(pwd)/template

```