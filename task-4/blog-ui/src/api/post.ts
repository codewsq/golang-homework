import request from './request'
import type { ApiResponse, Post } from '@/types'

export function getPosts() {
  return request.get<ApiResponse<{ posts: Post[] }>>('/post/getPosts')
}

export function getPost(id: number) {
  return request.get<ApiResponse<{ post: Post }>>(`/post/get/${id}`)
}

export function createPost(title: string, content: string) {
  return request.put<ApiResponse<{ post: Post }>>('/post/create', { title, content })
}

export function updatePost(id: number, title: string, content: string) {
  return request.post<ApiResponse<{ post: Post }>>(`/post/update/${id}`, { title, content })
}

export function deletePost(id: number) {
  return request.delete<ApiResponse>(`/post/delete/${id}`)
}
