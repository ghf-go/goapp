import axios from 'axios'
import { Message } from 'element-ui'
import router from '../router'

const request = axios.create({
  baseURL: '/api',
  timeout: 30000
})

request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = token
  }
  return config
})

request.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code === 0) {
      return res
    }
    if (res.code === 401) {
      handleUnauthorized(res.msg)
      return Promise.reject(new Error(res.msg || '未登录'))
    }
    Message.error(res.msg || '请求失败')
    return Promise.reject(new Error(res.msg || '请求失败'))
  },
  error => {
    if (error.response && error.response.status === 401) {
      handleUnauthorized('登录已过期，请重新登录')
    } else {
      const msg = (error.response && error.response.data && error.response.data.msg) || error.message || '网络错误'
      Message.error(msg)
    }
    return Promise.reject(error)
  }
)

function handleUnauthorized(msg) {
  localStorage.removeItem('token')
  localStorage.removeItem('username')
  if (msg) {
    Message.error(msg)
  }
  if (router.currentRoute.path !== '/login') {
    router.push('/login')
  }
}

export default request
