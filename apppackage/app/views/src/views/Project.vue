<template>
  <div>
    <el-card>
      <div class="toolbar">
        <el-input
          v-model="query.name"
          placeholder="工程名称"
          clearable
          style="width: 220px"
          @keyup.enter.native="handleSearch"
          @clear="handleSearch"
        />
        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
        <el-button type="primary" icon="el-icon-plus" style="margin-left: auto" @click="handleAdd">新增工程</el-button>
      </div>
      <el-table v-loading="loading" :data="list" border style="width: 100%">
        <el-table-column prop="name" label="名称" min-width="140" />
        <el-table-column prop="gitUrl" label="git地址" min-width="220" show-overflow-tooltip />
        <el-table-column label="类型" width="110">
          <template slot-scope="{ row }">
            <el-tag :type="typeTagType(row.type)" size="small">{{ typeLabel(row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="测试分支/版本" min-width="150">
          <template slot-scope="{ row }">{{ row.testBranch }} / {{ row.testVersion || '-' }}</template>
        </el-table-column>
        <el-table-column label="正式分支/版本" min-width="150">
          <template slot-scope="{ row }">{{ row.prodBranch }} / {{ row.prodVersion || '-' }}</template>
        </el-table-column>
        <el-table-column prop="updatedAt" label="更新时间" width="170" />
        <el-table-column label="操作" width="220" fixed="right">
          <template slot-scope="{ row }">
            <el-button type="text" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="text" size="small" @click="handleBuild(row)">构建</el-button>
            <el-button type="text" size="small" style="color: #f56c6c" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog :title="form.id ? '编辑工程' : '新增工程'" :visible.sync="dialogVisible" width="560px">
      <el-form ref="form" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="工程名称" />
        </el-form-item>
        <el-form-item label="git地址" prop="gitUrl">
          <el-input v-model="form.gitUrl" placeholder="https:// 或 git@ 地址" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择类型" style="width: 100%">
            <el-option label="移动应用" value="mobile" />
            <el-option label="桌面应用" value="desktop" />
            <el-option label="linux应用" value="linux" />
            <el-option label="web应用" value="web" />
          </el-select>
        </el-form-item>
        <el-form-item label="测试分支" prop="testBranch">
          <el-input v-model="form.testBranch" placeholder="如 develop" />
        </el-form-item>
        <el-form-item v-if="form.id" label="测试版本">
          <el-input v-model="form.testVersion" disabled />
        </el-form-item>
        <el-form-item label="正式分支" prop="prodBranch">
          <el-input v-model="form.prodBranch" placeholder="如 master" />
        </el-form-item>
        <el-form-item v-if="form.id" label="正式版本">
          <el-input v-model="form.prodVersion" disabled />
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" :loading="saving" @click="handleSave">保 存</el-button>
      </div>
    </el-dialog>

    <el-dialog title="选择构建环境" :visible.sync="buildVisible" width="420px">
      <el-radio-group v-model="buildEnv">
        <el-radio label="test">测试环境</el-radio>
        <el-radio label="prod">正式环境</el-radio>
      </el-radio-group>
      <div slot="footer">
        <el-button @click="buildVisible = false">取 消</el-button>
        <el-button type="primary" :loading="building" @click="confirmBuild">开始构建</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import request from '../utils/request'

const typeMap = {
  mobile: { label: '移动应用', tag: 'success' },
  desktop: { label: '桌面应用', tag: 'warning' },
  linux: { label: 'linux应用', tag: 'info' },
  web: { label: 'web应用', tag: 'primary' }
}

function emptyForm() {
  return {
    id: 0,
    name: '',
    gitUrl: '',
    type: '',
    testBranch: '',
    testVersion: '',
    prodBranch: '',
    prodVersion: ''
  }
}

export default {
  name: 'ProjectPage',
  data() {
    return {
      loading: false,
      saving: false,
      list: [],
      total: 0,
      query: { page: 1, size: 10, name: '' },
      dialogVisible: false,
      form: emptyForm(),
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        gitUrl: [{ required: true, message: '请输入git地址', trigger: 'blur' }],
        type: [{ required: true, message: '请选择类型', trigger: 'change' }],
        testBranch: [{ required: true, message: '请输入测试分支', trigger: 'blur' }],
        prodBranch: [{ required: true, message: '请输入正式分支', trigger: 'blur' }]
      },
      buildVisible: false,
      building: false,
      buildEnv: 'test',
      buildProject: null
    }
  },
  created() {
    this.loadList()
  },
  methods: {
    typeLabel(type) {
      return (typeMap[type] || {}).label || type
    },
    typeTagType(type) {
      return (typeMap[type] || {}).tag || 'info'
    },
    async loadList() {
      this.loading = true
      try {
        const res = await request.get('/project/list', { params: this.query })
        this.list = res.data.list || []
        this.total = res.data.total || 0
      } finally {
        this.loading = false
      }
    },
    handleSearch() {
      this.query.page = 1
      this.loadList()
    },
    handleAdd() {
      this.form = emptyForm()
      this.dialogVisible = true
    },
    handleEdit(row) {
      this.form = { ...row }
      this.dialogVisible = true
    },
    handleSave() {
      this.$refs.form.validate(async valid => {
        if (!valid) return
        this.saving = true
        try {
          const { id, name, gitUrl, type, testBranch, prodBranch } = this.form
          await request.post('/project/save', { id, name, gitUrl, type, testBranch, prodBranch })
          this.$message.success('保存成功')
          this.dialogVisible = false
          this.loadList()
        } finally {
          this.saving = false
        }
      })
    },
    handleDelete(row) {
      this.$confirm(`确认删除工程「${row.name}」？`, '提示', { type: 'warning' })
        .then(async () => {
          await request.post('/project/delete', { id: row.id })
          this.$message.success('删除成功')
          this.loadList()
        })
        .catch(() => {})
    },
    handleBuild(row) {
      this.buildProject = row
      this.buildEnv = 'test'
      this.buildVisible = true
    },
    async confirmBuild() {
      this.building = true
      try {
        await request.post('/build/start', { projectId: this.buildProject.id, env: this.buildEnv })
        this.$message.success('构建任务已提交')
        this.buildVisible = false
        this.$router.push('/build-record')
      } finally {
        this.building = false
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
</style>
