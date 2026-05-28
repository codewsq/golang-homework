import request from './request'
import type { ApiResponse, LoginResult, RegisterResult } from '@/types'

export function login(username: string, password: string) {
  return request.post<ApiResponse<LoginResult>>('/login', { username, password })
}

export function register(username: string, email: string, password: string) {
  return request.post<ApiResponse<RegisterResult>>('/register', {
    username,
    email,
    password,
  })
}
