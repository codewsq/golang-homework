<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { register } from '@/api/auth'
import { APP_TITLE } from '@/config'

const router = useRouter()
const loading = ref(false)

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
})

async function handleSubmit({ errors }: { errors?: Record<string, unknown> }) {
  if (errors) return
  if (form.password !== form.confirmPassword) {
    Message.warning('两次输入的密码不一致')
    return
  }
  loading.value = true
  try {
    await register(form.username, form.email, form.password)
    Message.success('注册成功，请登录')
    router.push({ name: 'Login' })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-page">
    <a-card :title="`注册 · ${APP_TITLE}`" class="auth-card">
      <a-form :model="form" layout="vertical" @submit="handleSubmit">
        <a-form-item
          field="username"
          label="用户名"
          :rules="[
            { required: true, message: '请输入用户名' },
            { minLength: 3, message: '至少 3 个字符' },
          ]"
        >
          <a-input v-model="form.username" placeholder="3-50 个字符" allow-clear />
        </a-form-item>
        <a-form-item
          field="email"
          label="邮箱"
          :rules="[
            { required: true, message: '请输入邮箱' },
            { type: 'email', message: '邮箱格式不正确' },
          ]"
        >
          <a-input v-model="form.email" placeholder="example@email.com" allow-clear />
        </a-form-item>
        <a-form-item
          field="password"
          label="密码"
          :rules="[
            { required: true, message: '请输入密码' },
            { minLength: 6, message: '至少 6 个字符' },
          ]"
        >
          <a-input-password v-model="form.password" placeholder="至少 6 个字符" />
        </a-form-item>
        <a-form-item
          field="confirmPassword"
          label="确认密码"
          :rules="[{ required: true, message: '请再次输入密码' }]"
        >
          <a-input-password v-model="form.confirmPassword" placeholder="再次输入密码" />
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" long :loading="loading">
            注册
          </a-button>
        </a-form-item>
        <div class="auth-footer">
          已有账号？
          <router-link to="/login">去登录</router-link>
        </div>
      </a-form>
    </a-card>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--color-fill-2) 0%, var(--color-fill-3) 100%);
  padding: 24px;
}

.auth-card {
  width: 100%;
  max-width: 400px;
}

.auth-footer {
  text-align: center;
  color: var(--color-text-3);
  font-size: 14px;
}

.auth-footer a {
  color: rgb(var(--primary-6));
}
</style>
