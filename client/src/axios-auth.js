import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://localhost:3000/'
})

instance.defaults.headers.common['SOMETHING'] = 'something'
instance.defaults.baseURL = 'http://localhost:3000'
instance.defaults.headers.common['Authorization'] = 'fasfdsa'
instance.defaults.headers.get['Accepts'] = 'application/json'

instance.interceptors.request.use(config => {
  console.log('Request Interceptor', config)
  return config
})

instance.interceptors.response.use(res => {
  console.log('Response Interceptor', res)
  return res
})

export default instance