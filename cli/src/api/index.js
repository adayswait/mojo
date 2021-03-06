
import axios from 'axios'

const httpClient = axios.create({
    baseURL: 'http://10.1.1.248/',
    timeout: 60000,
    headers: { 'X-Custom-Header': 'mojo' },
    withCredentials: true
});

// request拦截器
httpClient.interceptors.request.use(
    config => {
        // 根据条件加入token-安全携带
        // if (true) { // 需自定义
        //     // 让每个请求携带token
        //     config.headers['User-Token'] = '';
        // }
        return config;
    },
    error => {
        // 请求错误处理
        Promise.reject(error);
    }
)

// respone拦截器
httpClient.interceptors.response.use(
    response => {
        // 统一处理状态
        const res = response.data;
        if (res.code != 0) { // 需自定义
            // 返回异常
            window.console.error(res)
            return Promise.reject(res);
        } else {
            return res;
        }
    },
    // 处理处理
    error => {
        const ret = {
            code: -1,
            data: null
        }
        if (error && error.response) {
            ret.code = error.response.status;
            switch (error.response.status) {
                case 400:
                    error.message = '错误请求';
                    break;
                case 401:
                    error.message = '未授权，请重新登录';
                    break;
                case 403:
                    error.message = '拒绝访问';
                    break;
                case 404:
                    error.message = '请求错误,未找到该资源';
                    break;
                case 405:
                    error.message = '请求方法未允许';
                    break;
                case 408:
                    error.message = '请求超时';
                    break;
                case 500:
                    error.message = '服务器端出错';
                    break;
                case 501:
                    error.message = '网络未实现';
                    break;
                case 502:
                    error.message = '网络错误';
                    break;
                case 503:
                    error.message = '服务不可用';
                    break;
                case 504:
                    error.message = '网络超时';
                    break;
                case 505:
                    error.message = 'http版本不支持该请求';
                    break;
                default:
                    error.message = `未知错误${error.response.status}`;
            }
            ret.data = error.message;
        } else {
            ret.data = "连接到服务器失败";
        }
        window.console.error(ret);
        return Promise.reject(ret);
    }
)

export function get(url, params = {}) {
    return new Promise((resolve, reject) => {
        httpClient({
            url: url,
            method: 'get',
            params: params
        }).then(response => {
            resolve(response);
        }).catch(error => {
            reject(error);
        });
    });
}

export function put(url, params = {}) {
    return new Promise((resolve, reject) => {
        httpClient({
            url: url,
            method: 'put',
            data: params
        }).then(response => {
            resolve(response);
        }).catch(error => {
            reject(error);
        });
    });
}

export function del(url, params = {}) {
    return new Promise((resolve, reject) => {
        httpClient({
            url: url,
            method: 'delete',
            data: params
        }).then(response => {
            resolve(response);
        }).catch(error => {
            reject(error);
        });
    });
}

export function post(url, params = {}) {
    return new Promise((resolve, reject) => {
        httpClient({
            url: url,
            method: 'post',
            data: params
        }).then(response => {
            resolve(response);
        }).catch(error => {
            reject(error);
        });
    });
}

/*
 *  文件上传
 *  url:请求地址
 *  params:参数
 * */
export function fileUpload(url, params = {}) {
    return new Promise((resolve, reject) => {
        httpClient({
            url: url,
            method: 'post',
            data: params,
            headers: { 'Content-Type': 'multipart/form-data' }
        }).then(response => {
            resolve(response);
        }).catch(error => {
            reject(error);
        });
    });
}

export function isMobile() {
    if ((navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i))) {
        return true // 手机端
    } else {
        return false // alert('PC端')
    }
}

export default {
    get,
    post,
    del,
    put,
    fileUpload,
    isMobile
}