import axios from 'axios'
import { Message } from '@arco-design/web-vue'
import {
  API_BASE_URL,
  TOKEN_STORAGE_KEY,
  USER_STORAGE_KEY,
  USE_BEARER_PREFIX,
} from '@/config'
import router from '@/router'

const request = axios.create({
  baseURL: API_BASE_URL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

request.interceptors.request.use((config) => {
  const token = localStorage.getItem(TOKEN_STORAGE_KEY)
  if (token) {
    config.headers.Authorization = USE_BEARER_PREFIX ? `Bearer ${token}` : token
  }
  return config
})

request.interceptors.response.use(
  (response) => {
    const body = response.data
    if (body && typeof body.code === 'number' && body.code >= 400) {
      const msg = body.error || body.message || '请求失败'
      Message.error(msg)
      return Promise.reject(new Error(msg))
    }
    return response
  },
  (error) => {
    const status = error.response?.status
    const data = error.response?.data
    const requestUrl = error.config?.url ?? ''
    const msg = data?.error || data?.message || error.message || '网络错误'
    const isLoginRequest = requestUrl.includes('/login')
    const isRegisterRequest = requestUrl.includes('/register')

    if (status === 401) {
      if (isLoginRequest) {
        // 登录页凭证错误，与 token 过期区分
        Message.error('用户名或密码错误')
      } else if (isRegisterRequest) {
        Message.error(msg)
      } else {
        localStorage.removeItem(TOKEN_STORAGE_KEY)
        localStorage.removeItem(USER_STORAGE_KEY)
        Message.warning('登录已过期，请重新登录')
        if (router.currentRoute.value.name !== 'Login') {
          router.push({ name: 'Login' })
        }
      }
    } else {
      Message.error(msg)
    }
    return Promise.reject(error)
  },
)

export default request
