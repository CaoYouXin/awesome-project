import axios, {type AxiosRequestConfig} from 'axios'

interface Req {
  url: string,
  data: object,
  method: 'GET' | 'POST' | 'PUT' | 'DELETE'
}

function httpReq<T>(obj: Req, config: AxiosRequestConfig =
  {}): Promise<T> {
  return new Promise<T>((resolve, reject) => {
    let options: AxiosRequestConfig = {
      url: '',
      method: 'GET',
      data: {},
      params: {},
      timeout: 80 * 1000,
      baseURL: '',
      responseType: config.responseType,
      headers: {
        'Content-Type': 'application/json',
        'X-requested-With': 'XMLHttpRequest',
        'Authorization': 'Basic Zm9vOmJhcg==',
      },
    }

    options = {
      ...options, ...obj
    }

    if (options.method === 'GET') {
      options.params = options.data
    }

    axios(options).then(res => {
      if (res.status !== 200) {
        reject('HTTP:状态码异常！')
        return
      }

      let data = res.data
      // if (config.responseType === 'arraybuffer') {
      //   let blob = {}
      //   if (config.fileType.indexOf('ppt') > -1) {
      //     blob = new Blob([res.data], {
      //       type: 'application/vnd.ms-powerpoint'
      //     })
      //   } else if (config.fileType.indexOf('xls') > -1) {
      //     blob = new Blob([res.data], {
      //       type: 'application/vnd.ms-excel'
      //     })
      //   }
      //
      //   let downloadElement = document.createElement('a')
      //   let href = window.URL.createObjectURL(blob) //创建下载的链接
      //   downloadElement.href = href
      //   downloadElement.download = config.fileName + config.fileType //下载后文件名
      //   // downloadElement.download = "test.xlsx"; //下载后文件名
      //   document.body.appendChild(downloadElement)
      //   downloadElement.click() //点击下载
      //   document.body.removeChild(downloadElement) //下载完成移除元素
      //   window.URL.revokeObjectURL(href) //释放掉blob对象
      //   return
      // }

      if (parseInt(data.code) === 200 || data.success) {
        resolve(data.data || '');
      } else {
        reject(data.message);
      }
    }).catch(e => {
      console.log(e.response)
      reject('网络异常，请稍后再试!');
    })
  })
}

function HTTP<T>(obj: Req, config: AxiosRequestConfig = {}): Promise<T> {

  let defaultConfig: AxiosRequestConfig = {
    responseType: 'json',
  };

  config = {
    ...defaultConfig, ...config
  };

  return httpReq<T>(obj, config);
}

export default {
  GET: async <T>(url: string, data = {}, config: AxiosRequestConfig = {}) => {
    return HTTP<T>({
      url: await window.getApiRoot() + url, data, method: 'GET'
    }, config);
  },
  POST: async <T>(url: string, data = {}, config: AxiosRequestConfig = {}) => {
    return HTTP<T>({
      url: await window.getApiRoot() + url, data, method: 'POST'
    }, config);
  },
  PUT: async <T>(url: string, data = {}, config: AxiosRequestConfig = {}) => {
    return HTTP<T>({
      url: await window.getApiRoot() + url, data, method: 'PUT'
    }, config);
  },
  DELETE: async <T>(url: string, data = {}, config: AxiosRequestConfig = {}) => {
    return HTTP<T>({
      url: await window.getApiRoot() + url, data, method: 'DELETE'
    }, config);
  }
}
