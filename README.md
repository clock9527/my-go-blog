# my-go-blog

## 一、用户登录注册

### 1.1、登录 POST

url:     <http://localhost:8888/goblog/user/login>
param:   用户登录信息：{"username":"user2","password":"123456"}

### 1.2、注册 POST

url:     <http://localhost:8888/goblog/user/reg>
param:   用户注册信息：{"usernaem":"user2","password":"123456","email":"qwe@q.com"}

## 二、文章

### 2.1、文章列表 GET

url:     <http://localhost:8888/goblog/post/query/list>
param:   无

### 2.2、文章详情 GET

uri:     <http://localhost:8888/goblog/post/query/detail/1>
param:   文章ID:":id"

### 2.3、创建文章 POST

url:     <http://localhost:8888/goblog/post/modify/create>
headers: Authorization:登录成功后获得的秘钥
param:   文章内容：{"title":"test2","content":"Hi,moning"}

### 2.4、修改文章 POST

url:     <http://localhost:8888/goblog/post/modify/update/3>
headers: Authorization:登录成功后获得的秘钥
param:   文章id:":id"，
         文章内容：{"content":"hahahah"}

### 2.5、删除文章 DELETE

url:     <http://localhost:8888/goblog/post/modify/delete/3>
headers: Authorization:登录成功后获得的秘钥
param:   文章ID：":id"

## 三、评论

### 3.1、查看评论 GET

url:   <http://localhost:8888/goblog/comment/list/4>
param:  文章ID：":id"

### 3.2、评论创建 PUT

url:     <http://localhost:8888/goblog/comment/create/4>
headers: Authorization:登录成功后获得的秘钥
param:   文章ID：":id",
         评论内容：{"content":"asdfasd}
