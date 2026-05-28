/**
 * 应用全局配置
 * 部署或切换环境时，可优先修改本文件中的项
 */

/** 站点标题，会显示在浏览器标签与顶栏 */
export const APP_TITLE = '个人博客'

/**
 * API 基础路径
 * - 开发环境：使用 '/api'，由 vite.config.ts 代理到后端
 * - 生产环境：可改为完整地址，如 'http://your-server.com/api'
 */
export const API_BASE_URL = import.meta.env.DEV ? '/api' : 'http://localhost:8080/api'

/** localStorage 中 JWT 的键名 */
export const TOKEN_STORAGE_KEY = 'blog_token'

/** localStorage 中用户信息的键名 */
export const USER_STORAGE_KEY = 'blog_user'

/**
 * 是否在 Authorization 头前加 "Bearer " 前缀
 * 当前后端直接读取 token 字符串（无 Bearer），请保持 false
 * 若后端改为标准 Bearer 格式，可设为 true
 */
export const USE_BEARER_PREFIX = false

/**
 * Arco Design 主题色（primary）
 * 可选：arcoblue | blue | green | orange | red | purple 等
 */
export const ARCO_PRIMARY_COLOR = 'arcoblue'
