/**
 * Created by zzmhot on 2017/3/21.
 *
 * @author: zzmhot
 * @github: https://github.com/zzmhot
 * @email: zzmhot@163.com
 * @Date: 2017/3/21 16:04
 * @Copyright(©) 2017 by zzmhot.
 *
 */
import * as type from 'store/mutations/type'
import {cookieStorage} from 'common/storage'

export default {
    //设置用户信息和是否登录
    [type.SET_USER_INFO](state, userinfo) {
        state.user_info = userinfo || {}
        if (userinfo === null) {
            cookieStorage.remove('user_info')
        } else {
            cookieStorage.set('user_info', userinfo)
        }
    },


    set_user_session(state, user_session) {

        state.user_session = user_session || {};

        if (user_session === null) {
            cookieStorage.remove('user_session')
        } else {
            cookieStorage.set('user_session', user_session)
        }

    },

    set_token(state, token) {

        state.token = token || {};

        if (token === null) {
            cookieStorage.remove('token')
        } else {
            cookieStorage.set('token', token)
        }

    },

}
