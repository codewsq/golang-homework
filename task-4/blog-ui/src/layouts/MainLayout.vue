<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { APP_TITLE } from '@/config'

const router = useRouter()
const userStore = useUserStore()

function handleLogout() {
  userStore.logout()
  router.push({ name: 'Login' })
}
</script>

<template>
  <a-layout class="main-layout">
    <a-layout-header class="header">
      <div class="header-inner">
        <router-link to="/" class="logo">{{ APP_TITLE }}</router-link>
        <a-space>
          <a-button type="primary" @click="router.push({ name: 'PostCreate' })">
            <template #icon><icon-plus /></template>
            写文章
          </a-button>
          <a-dropdown trigger="click">
            <a-button type="text">
              <icon-user />
              {{ userStore.user?.username }}
            </a-button>
            <template #content>
              <a-doption @click="handleLogout">退出登录</a-doption>
            </template>
          </a-dropdown>
        </a-space>
      </div>
    </a-layout-header>
    <a-layout-content class="content">
      <div class="content-inner">
        <router-view />
      </div>
    </a-layout-content>
    <a-layout-footer class="footer">
      <!-- 页脚文案可按需修改 -->
      © {{ new Date().getFullYear() }} {{ APP_TITLE }}
    </a-layout-footer>
  </a-layout>
</template>

<style scoped>
.main-layout {
  min-height: 100vh;
}

.header {
  background: var(--color-bg-2);
  border-bottom: 1px solid var(--color-border);
}

.header-inner {
  max-width: 960px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
  padding: 0 16px;
}

.logo {
  font-size: 18px;
  font-weight: 600;
  color: rgb(var(--primary-6));
  text-decoration: none;
}

.content {
  background: var(--color-fill-2);
  min-height: calc(100vh - 64px - 48px);
}

.content-inner {
  max-width: 960px;
  margin: 0 auto;
  padding: 24px 16px;
}

.footer {
  text-align: center;
  color: var(--color-text-3);
  font-size: 13px;
}
</style>
