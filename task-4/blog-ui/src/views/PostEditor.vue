<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { getPost, createPost, updatePost } from '@/api/post'

const props = defineProps<{ id?: string }>()
const route = useRoute()
const router = useRouter()

const loading = ref(false)
const submitting = ref(false)

const isEdit = computed(() => route.name === 'PostEdit')
const postId = computed(() => Number(props.id))

const form = reactive({
  title: '',
  content: '',
})

async function loadPost() {
  if (!isEdit.value || !postId.value) return
  loading.value = true
  try {
    const { data } = await getPost(postId.value)
    const post = data.data?.post
    if (post) {
      form.title = post.title
      form.content = post.content
    }
  } finally {
    loading.value = false
  }
}

async function handleSubmit({ errors }: { errors?: Record<string, unknown> }) {
  if (errors) return
  submitting.value = true
  try {
    if (isEdit.value) {
      await updatePost(postId.value, form.title, form.content)
      Message.success('更新成功')
      router.push({ name: 'PostDetail', params: { id: postId.value } })
    } else {
      const { data } = await createPost(form.title, form.content)
      const newId = data.data?.post?.ID
      Message.success('创建成功')
      if (newId) {
        router.push({ name: 'PostDetail', params: { id: newId } })
      } else {
        router.push({ name: 'PostList' })
      }
    }
  } finally {
    submitting.value = false
  }
}

onMounted(loadPost)
</script>

<template>
  <a-spin :loading="loading" style="width: 100%">
    <a-page-header
      :title="isEdit ? '编辑文章' : '写文章'"
      @back="router.back()"
    />

    <a-card>
      <a-form :model="form" layout="vertical" @submit="handleSubmit">
        <a-form-item
          field="title"
          label="标题"
          :rules="[{ required: true, message: '请输入标题' }]"
        >
          <a-input
            v-model="form.title"
            placeholder="文章标题"
            :max-length="255"
            show-word-limit
            allow-clear
          />
        </a-form-item>
        <a-form-item
          field="content"
          label="正文"
          :rules="[{ required: true, message: '请输入正文' }]"
        >
          <a-textarea
            v-model="form.content"
            placeholder="支持纯文本，后续可接入 Markdown 编辑器"
            :auto-size="{ minRows: 12, maxRows: 24 }"
          />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" html-type="submit" :loading="submitting">
              {{ isEdit ? '保存' : '发布' }}
            </a-button>
            <a-button @click="router.back()">取消</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>
  </a-spin>
</template>
