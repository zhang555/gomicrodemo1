<template>
    <div class="login-body">
        <div class="loginWarp"
             v-loading="load_data"
             element-loading-text="正在登陆中...">
            
            <!--
               @keyup.enter="submit_form"
            
            -->
            <div class="login-title">
                <img src="./images/login_logo.png"/>
            </div>
            <div class="login-form">
                <el-form ref="form" :model="form" :rules="rules" label-width="0">
                    <el-form-item prop="username" class="login-item">
                        <el-input v-model="form.username" placeholder="请输入账户名：" class="form-input"
                                  :autofocus="true"></el-input>
                    </el-form-item>
                    <el-form-item prop="password" class="login-item">
                        <el-input type="password" v-model="form.password" placeholder="请输入账户密码："
                                  class="form-input"></el-input>
                    </el-form-item>
                    <el-form-item class="login-item">
                        
                        <!--
                        icon="check"
                        -->
                        <el-button size="large" class="form-submit" @click="login">登陆</el-button>
                        
                        
                        <!--<p>user: user</p>-->
                        <!--<p>user1: user1</p>-->
                        <!--<p>admin: admin</p>-->
                    
                    </el-form-item>
                </el-form>
            </div>
            
            <!--<el-button size="large" class="form-submit" @click="loginMethod(1)">user登陆</el-button>-->
            <!--<el-button size="large" class="form-submit" @click="loginMethod(1)">user1登陆</el-button>-->
            <!--<el-button size="large" class="form-submit" @click="loginMethod(1)">user1登陆</el-button>-->
        
        </div>
        
        <!--<p>formjson:{{formjson}}</p>-->
    </div>
</template>


<script type="text/javascript">

    import {mapMutations, mapState} from 'vuex';
    import global_ from '@/tool/Global'

    // import {mapActions} from 'vuex'
    // import {port_code, port_user} from 'common/port_uri'
    // import {SET_USER_INFO} from 'store/actions/type'

    const url = global_.urlPrefix + "/rpc"


    export default {
        data() {
            return {
                form: {
                    username: "",
                    password: "",
                },
                rules: {
                    username: [{required: true, message: '请输入账户名！', trigger: 'blur'}],
                    password: [{required: true, message: '请输入账户密码！', trigger: 'blur'}]
                },
                //请求时的loading效果
                load_data: false
            }
        },

        computed: {
            // ...mapState(["count"]),
            ...mapState([
                "user_session",
                'token',
            ]),


            formjson() {
                return JSON.stringify(this.form)
            }
        },

        methods: {
            ...mapMutations([
                'set_user_session',
                'set_token',
            ]),

            // ...mapActions({
            //     set_user_info: SET_USER_INFO
            // }),

            getUserSession() {
                this.$axios.post(url, {
                    service: "go.micro.srv.user",
                    method: "UserService.GetUserSessionByToken",
                    request: {
                        token: this.token,
                    }
                })
                    .then((response) => {
                        this.set_user_session(response.data.user)
                        global_.handleResponseGetThen(response)
                    })
                    .catch((error) => {
                        // this.load_data = false
                        this.$message.error(error.response.data.message);
                        // this.errordata = error.response.data
                    })
            },


            login() {

                let bean = {
                    service: "go.micro.srv.user",
                    method: "UserService.Auth",
                    request: this.form,
                }

                this.$axios.post(url, bean)
                    .then((response) => {
                        this.load_data = false
                        // if (response.data.success === true) {
                        if (response.data.code === 200) {
                            let token = response.data.token
                            // let result = response.data.token

                            // this.set_user_session(result)
                            this.set_token(token)
                            // this.$message.success(response.data.message);
                            this.$message.success("登陆成功");

                            this.getUserSession()

                            setTimeout(this.$router.push({path: '/user'}), 500)


                        }
                        if (response.data.code !== 200) {
                            // this.$message(response.data.message);
                            this.$message(response.data.msg);
                        }
                    })
                    .catch((error) => {
                        this.load_data = false
                        this.$message("登陆超时");
                    })

            },

            loginMethod(num) {

                let form = {}

                switch (num) {
                    case 1:
                        form = {
                            username: "user",
                            password: "user",
                        }
                        break

                    case 1:
                        form = {
                            username: "user1",
                            password: "user1",
                        }
                        break
                    case 1:
                        form = {
                            username: "admin",
                            password: "admin",
                        }
                        break
                }

                // let form = JSON.stringify(this.form)

                this.$axios.post(global_.urlPrefix + "/login", form)
                    .then((response) => {
                        this.load_data = false
                        if (response.data.success === true) {
                            let result = response.data.result

                            this.set_user_session(result)
                            this.$message.success(response.data.message);
                            setTimeout(this.$router.push({path: '/'}), 500)
                        }
                        if (response.data.success === false) {
                            this.$message(response.data.message);
                        }
                    })
                    .catch((error) => {
                        this.load_data = false
                        this.$message("登陆超时");
                    })

            },


            //提交
            submit_form() {
                this.$refs.form.validate((valid) => {
                    if (!valid) return false
                    this.load_data = true
                    //登录提交
                    this.$fetch.api_user.login(this.form)
                        .then(({data, msg}) => {
                            this.set_user_info({
                                user: data,
                                login: true
                            })
                            this.$message.success(msg)
                            setTimeout(this.$router.push({path: '/'}), 500)
                        })
                        .catch(({code}) => {
                            this.load_data = false
                            if (code === port_code.error) {
                                this.$notify.info({
                                    title: '温馨提示',
                                    message: '账号和密码都为：admin'
                                })
                            }
                        })
                })
            }
        }
    }
</script>
<style lang="scss" type="text/scss" rel="stylesheet/scss">
    .login-body {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        background-image: url(./images/login_bg.jpg);
        background-repeat: no-repeat;
        background-position: center;
        background-size: cover;
        .loginWarp {
            width: 300px;
            padding: 25px 15px;
            margin: 100px auto;
            background-color: #fff;
            border-radius: 5px;
            .login-title {
                margin-bottom: 25px;
                text-align: center;
            }
            .login-item {
                .el-input__inner {
                    margin: 0 !important;
                }
            }
            .form-input {
                input {
                    margin-bottom: 15px;
                    font-size: 12px;
                    height: 40px;
                    border: 1px solid #eaeaec;
                    background: #eaeaec;
                    border-radius: 5px;
                    color: #555;
                }
            }
            .form-submit {
                width: 100%;
                color: #fff;
                border-color: #6bc5a4;
                background: #6bc5a4;
                &:active, &:hover {
                    border-color: #6bc5a4;
                    background: #6bc5a4;
                }
            }
        }
    }
</style>
