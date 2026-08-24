<template>
  <div>
    <el-card>
      <div class="toolbar">
        <el-select v-model="query.projectId" placeholder="全部工程" clearable filterable style="width: 220px" @change="handleSearch">
          <el-option v-for="p in projectOptions" :key="p.id" :label="p.name" :value="p.id" />
        </el-select>
        <el-button type="primary" icon="el-icon-search" @click="handleSearch">查询</el-button>
      </div>
      <el-table v-loading="loading" :data="list" border style="width: 100%">
        <el-table-column prop="projectName" label="工程名" min-width="150" />
        <el-table-column label="环境" width="90">
          <template slot-scope="{ row }">
            <el-tag :type="row.env === 'prod' ? 'danger' : 'success'" size="small">
              {{ row.env === 'prod' ? '正式' : '测试' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="branch" label="分支" min-width="120" />
        <el-table-column prop="version" label="版本" min-width="120">
          <template slot-scope="{ row }">{{ row.version || '-' }}</template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template slot-scope="{ row }">
            <el-tag :type="statusTagType(row.status)" size="small">{{ statusLabel(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="170" />
        <el-table-column label="操作" width="90">
          <template slot-scope="{ row }">
            <el-button type="text" size="small" @click="handleLog(row)">日志</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        class="pager"
        background
        layout="total, prev, pager, next"
        :total="total"
        :page-size="query.size"
        :current-page.sync="query.page"
        @current-change="loadList"
      />
    </el-card>

    <el-drawer :title="`构建日志 #${logRecord ? logRecord.id : ''}`" :visible.sync="drawerVisible" size="60%" @close="stopLogPolling">
      <div class="log-wrap">
        <pre class="log-content">{{ logText || '暂无日志' }}</pre>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import request from '../utils/request'

const statusMap = {
  pending: { label: '排队中', tag: 'info' },
  running: { label: '构建中', tag: 'warning' },
  success: { label: '成功', tag: 'success' },
  failed: { label: '失败', tag: 'danger' }
}

export default {
  name: 'BuildRecordPage',
  data() {
    return {
      loading: false,
      list: [],
      total: 0,
      query: { page: 1, size: 10, projectId: '' },
      projectOptions: [],
      listTimer: null,
      drawerVisible: false,
      logRecord: null,
      logText: '',
      logTimer: null
    }
  },
  created() {
    this.loadProjectOptions()
    this.loadList()
  },
  beforeDestroy() {
    this.stopListPolling()
    this.stopLogPolling()
  },
  methods: {
    statusLabel(status) {
      return (statusMap[status] || {}).label || status
    },
    statusTagType(status) {
      return (statusMap[status] || {}).tag || 'info'
    },
    async loadProjectOptions() {
      try {
        const res = await request.get('/project/list', { params: { page: 1, size: 1000 } })
        this.projectOptions = res.data.list || []
      } catch (e) {
        // 工程下拉加载失败不影响列表
      }
    },
    async loadList() {
      this.loading = true
      try {
        const res = await request.get('/build/record/list', { params: this.query })
        this.list = res.data.list || []
        this.total = res.data.total || 0
        this.checkListPolling()
      } finally {
        this.loading = false
      }
    },
    handleSearch() {
      this.query.page = 1
      this.loadList()
    },
    checkListPolling() {
      const active = this.list.some(r => r.status === 'pending' || r.status === 'running')
      if (active && !this.listTimer) {
        this.listTimer = setInterval(() => this.loadList(), 3000)
      } else if (!active) {
        this.stopListPolling()
      }
    },
    stopListPolling() {
      if (this.listTimer) {
        clearInterval(this.listTimer)
        this.listTimer = null
      }
    },
    async handleLog(row) {
      this.logRecord = row
      this.logText = ''
      this.drawerVisible = true
      await this.loadLog()
    },
    async loadLog() {
      if (!this.logRecord) return
      const res = await request.get('/build/record/detail', { params: { id: this.logRecord.id } })
      this.logRecord = res.data
      this.logText = res.data.log || ''
      this.checkLogPolling()
    },
    checkLogPolling() {
      const active = this.logRecord && (this.logRecord.status === 'pending' || this.logRecord.status === 'running')
      if (active && !this.logTimer) {
        this.logTimer = setInterval(() => this.loadLog(), 3000)
      } else if (!active) {
        this.stopLogPolling()
      }
    },
    stopLogPolling() {
      if (this.logTimer) {
        clearInterval(this.logTimer)
        this.logTimer = null
      }
    }
  }
}
</script>

<style scoped>
.toolbar {
  display: flex;
  gap: 10px;
  margin-bottom: 16px;
}
.pager {
  margin-top: 16px;
  text-align: right;
}
.log-wrap {
  padding: 0 20px 20px;
  height: 100%;
  box-sizing: border-box;
}
.log-content {
  margin: 0;
  padding: 12px;
  background: #1e1e1e;
  color: #d4d4d4;
  font-family: Menlo, Monaco, Consolas, 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.5;
  white-space: pre-wrap;
  word-break: break-all;
  overflow: auto;
  height: 100%;
  box-sizing: border-box;
}
</style>
