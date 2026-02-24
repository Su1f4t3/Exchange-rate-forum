# 汇流社区

一个基于 Go 后端和 Vue.js 前端的汇率论坛应用，用户可以查看汇率、发布文章和互动交流。

## 功能特性

- 用户注册和登录（JWT 认证）
- 查看和创建汇率信息
- 发布和浏览文章/论坛帖子
- 文章点赞功能（Redis 缓存）
- 响应式设计（支持桌面和移动端）

## 技术栈

### 后端
- **框架**: Gin
- **数据库**: MySQL + GORM
- **缓存**: Redis
- **认证**: JWT + bcrypt

### 前端
- **框架**: Vue 3 + TypeScript
- **构建工具**: Vite
- **状态管理**: Pinia
- **UI 组件**: Element Plus（桌面端）、Vant（移动端）
- **HTTP 客户端**: Axios

## 快速开始

### 环境要求

- Go 1.25+
- Node.js 16+
- MySQL 8.0+
- Redis 6.0+

### 数据库配置

1. 创建 MySQL 数据库：
```sql
CREATE DATABASE test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 默认数据库配置：
- 主机：127.0.0.1:3306
- 数据库名：test
- 用户名：root
- 密码：123456

3. 确保 Redis 运行在 localhost:6379

### 后端启动

```bash
cd Exchangeapp_backend

# 安装依赖
go mod tidy

# 启动服务（端口 8080）
go run main.go
```

### 前端启动

```bash
cd Exchangeapp_frontend

# 安装依赖
npm install

# 启动开发服务器（端口 5173）
npm run dev

# 构建生产版本
npm run build
```

## API 文档

### 认证接口（公开）
- `POST /auth/login` - 用户登录
- `POST /auth/register` - 用户注册

### 汇率接口
- `GET /exchangerates` - 获取所有汇率（公开）
- `POST /exchangerates` - 创建汇率（需要登录）

### 文章接口
- `GET /articles` - 获取所有文章列表（Redis 缓存 5 分钟）
- `GET /articles/:id` - 获取指定文章
- `POST /articles` - 创建文章（需要登录）

### 点赞接口
- `POST /articles/:id/like` - 点赞文章（需要登录）
- `GET /articles/:id/like` - 获取文章点赞数

## 项目结构

```
Exchange_app/
├── Exchangeapp_backend/     # Go 后端
│   ├── config/             # 配置文件
│   ├── controllers/        # 控制器
│   ├── models/            # 数据模型
│   ├── router/            # 路由配置
│   ├── middlewares/       # 中间件
│   ├── utils/             # 工具函数
│   └── main.go            # 入口文件
└── Exchangeapp_frontend/   # Vue 前端
    ├── src/
    │   ├── views/         # 页面组件
    │   ├── components/    # 公共组件
    │   ├── router/        # 路由配置
    │   ├── store/         # 状态管理
    │   └── types/         # TypeScript 类型
    ├── package.json
    └── vite.config.ts     # Vite 配置
```

## 配置说明

### 后端配置
配置文件位于 `Exchangeapp_backend/config/config.yml`：

```yaml
app:
  name: CurrencyExchangeApp
  port: :3000  # 注意：实际运行端口为 8080

database:
  dsn: "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
  MaxIdleConns: 10
  MaxOpenConns: 100

redis:
  Addr: "localhost:6379"
```

### 前端配置
Vite 配置位于 `Exchangeapp_frontend/vite.config.ts`，API 请求会代理到后端服务器。

## 开发注意事项

1. 后端服务实际运行在 8080 端口，而不是配置文件中的 3000 端口
2. JWT 密钥硬编码为 "secret"，生产环境请修改
3. 文章列表使用 Redis 缓存，缓存时间为 5 分钟
4. 点赞数据存储在 Redis 中，使用 `article:{id}:likes` 格式
5. CORS 已配置允许来自 `http://localhost:5173` 的请求

## 贡献

欢迎提交 Issue 和 Pull Request 来改进项目。

## 许可证

ISC License
