import request from './request'
import type { ApiResponse, Comment } from '@/types'

export function createComment(content: string, postId: number) {
  return request.put<ApiResponse<Comment>>('/comment/create', {
    content,
    post_id: postId,
  })
}

/** 获取某篇文章的全部评论，data 为评论数组 */
export function getComments(postId: number) {
  return request.get<ApiResponse<Comment[]>>(`/comment/get/${postId}`)
}
