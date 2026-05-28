/** 后端统一响应结构 */
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data?: T
  error?: string
}

export interface User {
  id?: number
  ID?: number
  username: string
  email?: string
}

export interface Post {
  ID: number
  title: string
  content: string
  user_id: number
  user?: User
  published: boolean
  created_at: string
  updated_at: string
}

export interface Comment {
  ID: number
  Content: string
  content?: string
  user_id: number
  user?: User
  post_id: number
  created_at: string
}

export interface LoginResult {
  token: string
  user: User
}

export interface RegisterResult {
  user: User
}
