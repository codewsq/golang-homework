/** 格式化后端返回的时间字符串 */
export function formatDateTime(value?: string): string {
  if (!value || value.startsWith('0001-01-01')) return '-'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

/** 评论正文字段（后端可能为 Content 或 content） */
export function commentText(c: { Content?: string; content?: string }): string {
  return c.Content ?? c.content ?? ''
}

/** 评论作者名称（来自嵌套的 user.username） */
export function commentAuthor(c: { user?: { username?: string } }): string {
  const name = c.user?.username?.trim()
  return name || '未知用户'
}
