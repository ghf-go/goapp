<template>
  <el-container class="layout">
    <el-aside width="220px" class="aside">
      <div class="side-header">
        <div class="system-name">打包构建系统</div>
        <div class="username">{{ username }}</div>
      </div>
      <el-menu
        class="side-menu"
        :default-active="$route.path"
        router
        background-color="#001529"
        text-color="#a6adb4"
        active-text-color="#ffffff"
      >
        <el-menu-item index="/project">
          <i class="el-icon-folder-opened"></i>
          <span slot="title">工程管理</span>
        </el-menu-item>
        <el-menu-item index="/build-record">
          <i class="el-icon-document"></i>
          <span slot="title">构建记录</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header" height="56px">
        <el-button type="text" icon="el-icon-switch-button" @click="handleLogout">退出登录</el-button>
      </el-header>
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: 'MainLayout',
  computed: {
    username() {
      return localStorage.getItem('username') || ''
    }
  },
  methods: {
    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.layout {
  height: 100%;
}
.aside {
  background-color: #001529;
}
.side-header {
  padding: 20px 16px;
  color: #ffffff;
}
.system-name {
  font-size: 18px;
  font-weight: bold;
}
.username {
  margin-top: 6px;
  font-size: 13px;
  color: #a6adb4;
}
.side-menu {
  border-right: none;
}
.header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  border-bottom: 1px solid #e6e6e6;
  background: #ffffff;
}
.main {
  background: #f0f2f5;
}
</style>
