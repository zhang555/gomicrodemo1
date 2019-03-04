/**
 * Created by zzmhot on 2017/3/23.
 *
 * 主程序入口
 *
 * @author: zzmhot
 * @github: https://github.com/zzmhot
 * @email: zzmhot@163.com
 * @Date: 2017/3/23 18:19
 * @Copyright(©) 2017 by zzmhot.
 *
 */

//导入样式
import 'normalize.css'

import axios from 'axios'

import 'font-awesome/scss/font-awesome.scss'
import 'element-ui/lib/theme-chalk/index.css';
//导入Vue框架
import Vue from 'vue'
//导入element组件
import ElementUI from 'element-ui'
//导入组件
import router from './router'
//导入状态管理器
import store from 'store'
//导入请求框架
// import api from './api'
//导入自定义插件
// import Plugins from 'plugins'
//导入主视图文件
import App from './App'

import '@/assets/scss/table.css';
import {cookieStorage} from 'common/storage'
// import 'element-ui/lib/theme-default/index.css'


//使用element-ui
Vue.use(ElementUI)

//使用自定义插件
// Vue.use(Plugins)

//使用api
// Vue.use(api)

//发布后是否显示提示
Vue.config.productionTip = false

//是否开启工具调试
Vue.config.devtools = process.env.NODE_ENV === 'development'


axios.defaults.withCredentials = true
axios.defaults.timeout = 1000 * 10 // 请求的超时时间


axios.interceptors.request.use(
    config => {

        config.headers["X-CSRF-Token"] = cookieStorage.get('token');

        // console.log("headers")
        // console.log(config)
        // console.log(config.headers)


        // if (store.state.token) {  // 判断是否存在token，如果存在的话，则每个http header都加上token
        //     config.headers.Authorization = `token ${store.state.token}`;
        // }
        return config;
    },
    err => {
        return Promise.reject(err);
    });

axios.interceptors.response.use(
    response => {
        if (response) {
            switch (response.data.code) {
                case 4011:
                    store.commit("set_user_session", {});
                    router.push({name: 'login'})
                    break

                case 401:
                    break


                // router.push({path: '/'})
            }
        }

        return response;
    },
    error => {
        // if (error.response) {
        //     switch (error.response.status) {
        //         case 401:
        //
        //             break
        //     }
        // }
        return Promise.reject(error)   // 返回接口返回的错误信息
    });


// axios.defaults.headers.post['token'] = cookieStorage.get('token');
// axios.defaults.headers.post['token'] = "123123123123";


Vue.prototype.$axios = axios


// console.log(Vue.prototype.$axios)


new Vue({
    router,
    store,
    axios,
    ...App
}).$mount('mainbody')

//
// new Vue({
//     el: '#app',
//     router,
//     store,
//     axios,
//     components: {App},
//     template: '<App/>'
// })