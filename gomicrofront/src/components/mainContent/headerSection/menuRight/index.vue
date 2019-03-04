<template>
    <!--<div class="menu-right" v-if="get_user_info.login">-->
    <div class="menu-right" v-if="user_session.username !== '' ">
        <div class="notification-menu">
            
            <!--<el-tag class="tag" type="primary">-->
            <!--roleName：{{user_session.roleName}}-->
            <!--</el-tag>-->
            
            
            <el-tag class="tag" type="primary">
                token：{{token}}
            </el-tag>
            
            
            <el-dropdown trigger="click" class="notification-list">
                <div class="notification-btn">
                    <!--<img :src="get_user_info.user.avatar" :alt="get_user_info.user.name"/>-->
                    <!--<span v-text="get_user_info.user.name"></span>-->
                    
                    
                    <!--<img :src="get_user_info.user.avatar" :alt="get_user_info.user.name"/>-->
                    <span v-text="user_session.username"></span>
                    
                    
                    <span class="icon"></span>
                </div>
                
                
                <el-dropdown-menu slot="dropdown" class="dropdown-menu">
                    
                    
                    <!--<el-dropdown-item class="dropdown-list">-->
                    <!--<a href="javascript:" class="dropdown-btn" @click="user_click(0)">-->
                    <!--<i class="icon fa fa-user"></i>-->
                    <!--<span>个人信息</span>-->
                    <!--</a>-->
                    <!--</el-dropdown-item>-->
                    <!--<el-dropdown-item class="dropdown-list">-->
                    <!--<a href="javascript:" class="dropdown-btn" @click="user_click(0)">-->
                    <!--<i class="icon fa fa-cog"></i>-->
                    <!--<span>设置</span>-->
                    <!--</a>-->
                    <!--</el-dropdown-item>-->
                    <el-dropdown-item class="dropdown-list">
                        <a href="javascript:" class="dropdown-btn" @click="user_click(0)">
                            <i class="icon fa fa-sign-out"></i>
                            <span>安全退出</span>
                        </a>
                    </el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
        
        <p>user_session:{{user_session}}</p>
    </div>
</template>


<script type="text/javascript">
    // import {mapActions, mapGetters, mapMutations, mapState} from 'vuex'
    import {mapMutations, mapState} from 'vuex';
    import global_ from '@/tool/Global'

    // import {GET_USER_INFO} from 'store/getters/type'
    // import {SET_USER_INFO} from 'store/actions/type'

    const url = global_.urlPrefix + "/rpc"


    const USER_OUT = 0
    const USER_INFO = 1
    const USER_SETTING = 2

    export default {
        computed: {
            // ...mapGetters({
            //     get_user_info: GET_USER_INFO,
            // }),

            ...mapState([
                "user_session",
                "token",
            ])


        },
        methods: {
            // ...mapActions({
            //     set_user_info: SET_USER_INFO
            // }),

            ...mapMutations([
                'set_user_session',
                'set_token',

            ]),


            //退出
            user_out() {
                // this.$confirm('此操作将退出登录, 是否继续?', '提示', {
                //     confirmButtonText: '确定',
                //     cancelButtonText: '取消',
                //     type: 'warning'
                // }).then(() => {
                //     this.logout()
                //
                //
                // }).catch(() => {
                //
                // })
                this.logout()

            },


            logout() {

                // if (this.token === {}) {
                //     this.set_token("")
                // }


                let bean = {
                    service: "go.micro.srv.user",
                    method: "UserService.Logout",
                    request: {
                        token: this.token,
                    },
                }

                this.$axios.post(url, bean)
                    .then((response) => {
                        // this.set_user_session(null)
                        this.set_token(null)
                        this.$router.push({path: '/login'})
                        // setTimeout(), 500)
                    })
                    .catch((error) => {

                    })
            },


            user_info() {
                //个人信息
            },
            user_setting() {
                //设置
            },
            user_click(type) {
                switch (type) {
                    case USER_OUT :
                        //退出
                        this.user_out()
                        break
                    case USER_INFO:
                        //个人信息
                        break
                    case USER_SETTING:
                        //设置
                        break
                }
            }
        }
    }
</script>


<style lang="scss" type="text/scss" rel="stylesheet/scss" scoped>
    
    .tag {
        /*padding-right: 30px;*/
        margin-right: 30px;
    }
</style>
