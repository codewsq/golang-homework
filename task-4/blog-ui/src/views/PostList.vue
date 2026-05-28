<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Message, Modal } from '@arco-design/web-vue'
import { getPosts, deletePost } from '@/api/post'
import { useUserStore } from '@/stores/user'
import { formatDateTime } from '@/utils/format'
import type { Post } from '@/types'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const posts = ref<Post[]>([])

async function fetchPosts() {
  loading.value = true
  try {
    const { data } = await getPosts()
    posts.value = data.data?.posts ?? []
  } finally {
    loading.value = false
  }
}

function isOwner(post: Post) {
  return post.user_id === userStore.user?.id
}

function handleDelete(post: Post) {
  Modal.warning({
    title: '确认删除',
    content: `确定要删除文章「${post.title}」吗？此操作不可恢复。`,
    hideCancel: false,
    onOk: async () => {
      await deletePost(post.ID)
      Message.success('删除成功')
      await fetchPosts()
    },
  })
}

onMounted(fetchPosts)
</script>

<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-empty v-if="!loading && posts.length === 0" description="暂无文章，快来写一篇吧" />

    <a-space v-else direction="vertical" fill :size="16" style="width: 100%">
      <a-card
        v-for="post in posts"
        :key="post.ID"
        hoverable
        class="post-card"
        @click="router.push({ name: 'PostDetail', params: { id: post.ID } })"
      >
        <template #title>
          <span class="post-title">{{ post.title }}</span>
        </template>
        <template #extra>
          <a-space @click.stop>
            <a-tag v-if="post.published" color="green">已发布</a-tag>
            <a-tag v-else color="gray">草稿</a-tag>
            <template v-if="isOwner(post)">
              <a-button
                type="text"
                size="small"
                @click="router.push({ name: 'PostEdit', params: { id: post.ID } })"
              >
                编辑
              </a-button>
              <a-button type="text" status="danger" size="small" @click="handleDelete(post)">
                删除
              </a-button>
            </template>
          </a-space>
        </template>
        <p class="post-excerpt">{{ post.content }}</p>
        <div class="post-meta">
          <a-space>
            <span><icon-user /> {{ post.user?.username ?? '未知' }}</span>
            <span><icon-clock-circle /> {{ formatDateTime(post.created_at) }}</span>
          </a-space>
        </div>
      </a-card>
    </a-space>
  </a-spin>
</template>

<style scoped>
.post-card {
  cursor: pointer;
}

.post-title {
  font-size: 18px;
}

.post-excerpt {
  margin: 0 0 12px;
  color: var(--color-text-2);
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  line-height: 1.6;
  white-space: pre-wrap;
}

.post-meta {
  font-size: 13px;
  color: var(--color-text-3);
}
</style>
