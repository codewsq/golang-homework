<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Message, Modal } from '@arco-design/web-vue'
import { getPost, deletePost } from '@/api/post'
import { getComments, createComment } from '@/api/comment'
import { useUserStore } from '@/stores/user'
import { formatDateTime, commentText, commentAuthor } from '@/utils/format'
import type { Post, Comment } from '@/types'

const props = defineProps<{ id: string }>()
const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const commentLoading = ref(false)
const submitCommentLoading = ref(false)
const post = ref<Post | null>(null)
const comments = ref<Comment[]>([])
const commentContent = ref('')

const postId = computed(() => Number(props.id))
const isOwner = computed(() => post.value?.user_id === userStore.user?.id)

async function fetchPost() {
  loading.value = true
  try {
    const { data } = await getPost(postId.value)
    post.value = data.data?.post ?? null
  } finally {
    loading.value = false
  }
}

async function fetchComments() {
  commentLoading.value = true
  try {
    const { data } = await getComments(postId.value)
    comments.value = Array.isArray(data.data) ? data.data : []
  } finally {
    commentLoading.value = false
  }
}

async function handleSubmitComment() {
  const text = commentContent.value.trim()
  if (!text) {
    Message.warning('请输入评论内容')
    return
  }
  submitCommentLoading.value = true
  try {
    await createComment(text, postId.value)
    commentContent.value = ''
    Message.success('评论发表成功')
    await fetchComments()
  } finally {
    submitCommentLoading.value = false
  }
}

function handleDelete() {
  if (!post.value) return
  Modal.warning({
    title: '确认删除',
    content: `确定要删除文章「${post.value.title}」吗？`,
    hideCancel: false,
    onOk: async () => {
      await deletePost(post.value!.ID)
      Message.success('删除成功')
      router.push({ name: 'PostList' })
    },
  })
}

onMounted(async () => {
  await fetchPost()
  await fetchComments()
})
</script>

<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-page-header
      v-if="post"
      :title="post.title"
      @back="router.push({ name: 'PostList' })"
    >
      <template #subtitle>
        <a-space>
          <span>{{ post.user?.username }}</span>
          <span>{{ formatDateTime(post.updated_at) }}</span>
        </a-space>
      </template>
      <template #extra>
        <a-space v-if="isOwner">
          <a-button @click="router.push({ name: 'PostEdit', params: { id: post.ID } })">
            编辑
          </a-button>
          <a-button status="danger" @click="handleDelete">删除</a-button>
        </a-space>
      </template>
    </a-page-header>

    <a-card v-if="post" class="post-body">
      <pre class="post-content">{{ post.content }}</pre>
    </a-card>

    <a-card title="评论" class="comment-section">
      <a-form :model="{ comment: commentContent }" layout="vertical">
        <a-form-item field="comment" label="发表评论">
          <a-textarea
            v-model="commentContent"
            placeholder="写下你的想法..."
            :auto-size="{ minRows: 3, maxRows: 6 }"
            :max-length="500"
            show-word-limit
          />
        </a-form-item>
        <a-form-item>
          <a-button
            type="primary"
            :loading="submitCommentLoading"
            @click="handleSubmitComment"
          >
            发表评论
          </a-button>
        </a-form-item>
      </a-form>

      <a-divider />

      <a-spin :loading="commentLoading">
        <a-empty v-if="comments.length === 0" description="暂无评论" />
        <a-comment
          v-for="item in comments"
          :key="item.ID"
          :author="commentAuthor(item)"
          :datetime="formatDateTime(item.created_at)"
          align="right"
          class="comment-item"
        >
          <template #avatar>
            <a-avatar :style="{ backgroundColor: '#165DFF' }">
              {{ commentAuthor(item).charAt(0).toUpperCase() }}
            </a-avatar>
          </template>
          <template #content>
            <div class="comment-body">{{ commentText(item) }}</div>
          </template>
        </a-comment>
      </a-spin>
    </a-card>
  </a-spin>
</template>

<style scoped>
.post-body {
  margin-bottom: 16px;
}

.post-content {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-word;
  font-family: inherit;
  font-size: 15px;
  line-height: 1.8;
  color: var(--color-text-1);
  background: transparent;
  border: none;
}

.comment-section {
  margin-top: 16px;
}

.comment-item {
  margin-bottom: 8px;
}

.comment-body {
  color: var(--color-text-1);
  line-height: 1.6;
  white-space: pre-wrap;
}
</style>
