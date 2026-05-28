# 个人博客前端 (blog-ui)

基于 **Vue 3 + Vite + TypeScript + Arco Design Vue** 的个人博客管理界面，对接后端 REST API（默认 `http://localhost:8080`）。

## 功能概览

| 模块 | 说明 |
|------|------|
| 用户注册 / 登录 | 对接 `/api/register`、`/api/login`，JWT 本地持久化 |
| 文章列表 | 展示全部文章，支持查看详情 |
| 文章 CRUD | 创建、编辑、删除（仅作者本人可操作） |
| 评论 | 在文章详情页发表评论、查看评论及评论人 |

---

## 环境要求

| 工具 | 建议版本 |
|------|----------|
| Node.js | ≥ 18（推荐 20 LTS） |
| npm | ≥ 9 |
| 后端服务 | 见 `../blog-server`，默认端口 **8080** |

---

## 项目结构

```text
blog-ui/
├── index.html
├── package.json
├── vite.config.ts          # Vite 构建与开发代理
├── src/
│   ├── config/index.ts     # ★ 应用级可配置项（优先修改）
│   ├── api/                # 接口封装（axios）
│   ├── layouts/            # 布局
│   ├── router/             # 路由与鉴权
│   ├── stores/             # Pinia 状态
│   ├── types/              # TypeScript 类型
│   ├── utils/              # 工具函数
│   └── views/              # 页面
└── dist/                   # 生产构建输出（npm run build 后生成）
```

---

## 配置修改方式

本项目配置分两层：**应用配置** 与 **构建/代理配置**。按场景选择修改位置。

### 1. 应用配置（推荐优先改这里）

文件：[`src/config/index.ts`](src/config/index.ts)

| 配置项 | 说明 | 示例 |
|--------|------|------|
| `APP_TITLE` | 浏览器标题、顶栏站点名 | `'个人博客'` |
| `API_BASE_URL` | 接口基础路径 | 开发：`'/api'`；生产：`'https://api.example.com/api'` |
| `TOKEN_STORAGE_KEY` | localStorage 中 token 键名 | `'blog_token'` |
| `USER_STORAGE_KEY` | localStorage 中用户信息键名 | `'blog_user'` |
| `USE_BEARER_PREFIX` | Authorization 是否加 `Bearer ` 前缀 | 当前后端为裸 token，保持 `false` |

`API_BASE_URL` 当前逻辑：

```ts
export const API_BASE_URL = import.meta.env.DEV ? '/api' : 'http://localhost:8080/api'
```

- **开发环境**（`npm run dev`）：固定走 `/api`，由 Vite 代理到后端，无需改此项即可本地联调。
- **生产环境**（`npm run build`）：请将 `http://localhost:8080/api` 改为实际后端公网/内网地址。

### 2. 开发代理配置（仅本地联调）

文件：[`vite.config.ts`](vite.config.ts)

```ts
server: {
  port: 5173,  // 前端开发端口，可按需修改
  proxy: {
    '/api': {
      target: 'http://localhost:8080',  // 后端地址
      changeOrigin: true,
    },
  },
},
```

当后端不在本机或端口不是 8080 时，修改 `proxy['/api'].target` 即可，**无需**改 `API_BASE_URL`（开发环境仍用 `/api`）。

### 3. 其他可调项

| 位置 | 作用 |
|------|------|
| `src/main.ts` | Arco 组件全局 `size`（mini / small / medium / large） |
| `src/App.vue` | Arco 语言包（当前为 `zh-cn`） |
| `index.html` | 页面 `<title>`（也可由 `APP_TITLE` 在运行时体现） |

---

## 测试环境启动方式

测试环境指：**本地开发 + 本地后端**，用于功能联调与接口验证。

### 步骤 1：启动后端

进入后端目录（与 `blog-ui` 同级）：

```bash
cd ../blog-server
```

任选一种方式：

```bash
# 方式一：直接运行可执行文件（Windows）
./server.exe

# 方式二：Go 命令启动
go run main.go
```

确认后端可用：浏览器或 curl 访问 `http://localhost:8080/health`，应返回 `status: ok`。

### 步骤 2：安装前端依赖（首次）

```bash
cd ../blog-ui
npm install
```

### 步骤 3：启动前端开发服务

```bash
npm run dev
```

### 步骤 4：访问

- 前端地址：**http://localhost:5173**
- 接口请求：`/api/*` → Vite 代理 → `http://localhost:8080/api/*`

### 常用测试命令

```bash
# 类型检查 + 生产构建（验证能否正常打包）
npm run build

# 本地预览构建结果（需先 build）
npm run preview
```

`preview` 默认端口 **4173**，此时 **不会** 走 Vite 代理，需保证 `src/config/index.ts` 中生产用 `API_BASE_URL` 指向可访问的后端，或仅用于静态资源检查。

---

## 生产环境启动方式

生产环境指：将前端构建为静态资源，由 Web 服务器托管，并访问**已部署的后端 API**。

### 步骤 1：修改生产 API 地址

编辑 [`src/config/index.ts`](src/config/index.ts)，将非开发环境下的 `API_BASE_URL` 改为真实后端，例如：

```ts
export const API_BASE_URL = import.meta.env.DEV
  ? '/api'
  : 'https://your-domain.com/api'
```

### 步骤 2：构建

```bash
npm install
npm run build
```

产物目录：`dist/`（`index.html` + `assets/`）。

### 步骤 3：部署静态资源

将 `dist/` 内全部文件上传到 Nginx、Caddy、对象存储静态站点等。

**Nginx 示例**（仅前端；API 由独立域名/路径提供时需自行处理跨域或由网关反代）：

```nginx
server {
    listen 80;
    server_name blog.example.com;
    root /var/www/blog-ui/dist;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

若希望生产环境也通过同源 `/api` 访问后端，可在 Nginx 增加反代（此时 `API_BASE_URL` 可设为 `'/api'`）：

```nginx
location /api/ {
    proxy_pass http://127.0.0.1:8080/api/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
}
```

### 步骤 4：启动/访问

- 用户通过配置的域名访问前端（如 `http://blog.example.com`）。
- 确保后端服务已在线上运行，且前端 `API_BASE_URL` 与网络策略（防火墙、HTTPS）一致。

### 生产环境本地快速验证（可选）

```bash
npm run build
npm run preview
```

访问 **http://localhost:4173**。注意：若 `API_BASE_URL` 仍为 `http://localhost:8080/api`，需本机 8080 后端同时运行。

---

## 部署手册（简要 checklist）

1. [ ] 后端已部署并可访问（如 `/health` 正常）。
2. [ ] 修改 `src/config/index.ts` 中生产 `API_BASE_URL`。
3. [ ] 执行 `npm run build`，确认无报错。
4. [ ] 将 `dist/` 部署到 Web 服务器或 CDN。
5. [ ] 配置 SPA 路由回退（`try_files ... /index.html`）。
6. [ ] （推荐）配置 HTTPS 与 `/api` 反向代理，减少跨域问题。
7. [ ] 浏览器访问：注册 → 登录 → 发文 → 评论，做一次冒烟测试。

---

## 与后端的对接说明

| 说明 | 详情 |
|------|------|
| 默认后端地址 | `http://localhost:8080` |
| 认证方式 | 请求头 `Authorization: <token>`（无 Bearer 前缀） |
| 接口文档 | 见 `../个人博客系统后端接口文档.md` |
| 登录失败 | HTTP 401，前端提示「用户名或密码错误」 |
| Token 过期 | 其他接口 401，清除本地登录态并跳转登录页 |

---

## 常见问题

**Q：开发时接口 404 或跨域？**  
A：确认后端已启动；确认 `vite.config.ts` 中 `proxy.target` 为 `http://localhost:8080`；开发环境 `API_BASE_URL` 应为 `'/api'`。

**Q：生产环境登录/列表空白？**  
A：检查 `API_BASE_URL` 是否指向正确后端；浏览器开发者工具 Network 查看请求是否失败或 CORS 被拦。

**Q：登录成功但刷新后又要登录？**  
A：检查浏览器是否禁用 localStorage；`TOKEN_STORAGE_KEY` 是否被其他脚本清除。

---

## 脚本说明

| 命令 | 说明 |
|------|------|
| `npm run dev` | 测试环境：启动 Vite 开发服务器（端口 5173） |
| `npm run build` | 生产构建，输出到 `dist/` |
| `npm run preview` | 本地预览 `dist/`（端口 4173） |
