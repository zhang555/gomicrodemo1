<script type="">
    import Vue from 'vue'

    const isDevelopment = process.env.NODE_ENV === 'development'


    // let url = window.location.href
    // let index = url.lastIndexOf("#")
    // let suffix = url.slice( 0 ,index - 1)
    let suffix = document.domain


    export default {

        // urlPrefix: isDevelopment ? "http://localhost:8228" : "http://192.168.0.112:8228",
        // urlPrefix: isDevelopment ? "http://localhost:8080" : "http://" + suffix + ":8228",
        urlPrefix: isDevelopment ? "http://localhost:8228" : "http://" + suffix + ":8228",

        pageSizeNum: 10,
        pageSizeString: "10",


        //
        handleResponseGetThen(response) {
            if (response.data.success === false) {
                if (response.data.code === 300000) {
                    Vue.prototype.$message("登陆过期，请重新登录")
                } else {
                    Vue.prototype.$message.error(response.data.message)
                }
                // Vue.prototype.$message("请重新登录")
            }
        },


        //
        handleResponseThen(response) {
            if (response.data.code === 200) {
                Vue.prototype.$message.success(response.data.msg)
            }
            if (response.data.code !== 200) {
                Vue.prototype.$message.error(response.data.msg)
            }
        },


        //
        handleResponseThenCustomizePost(response) {
            if (response.data.code === 200) {
                Vue.prototype.$message.success("创建成功")
            }
            if (response.data.code !== 200) {
                Vue.prototype.$message.error(response.data.msg)
            }
        },


        //
        handleResponseThenCustomizePut(response) {
            if (response.data.code === 200) {
                Vue.prototype.$message.success("修改成功")
            }
            if (response.data.code !== 200) {
                Vue.prototype.$message.error(response.data.msg)
            }
        },

        handleResponseThenCustomizeDelete(response) {
            if (response.data.code === 200) {
                Vue.prototype.$message.success("删除成功")
            }
            if (response.data.code !== 200) {
                Vue.prototype.$message.error(response.data.msg)
            }
        },


        //
        handleResponseCatch(error) {
            // Vue.prototype.$message("失败了哦");
            Vue.prototype.$message("服务器内部出错");
        },


        //
        handleUploadError(err) {
            Vue.prototype.$message("上传失败，请重试");
        },


        // getUserSession() {
        //     Vue.prototype.$axios.post(this.urlPrefix + '/rpc', {
        //         service: "go.micro.srv.user",
        //         method: "UserService.GetUserSessionByToken",
        //         request: {
        //             token: this.token,
        //         }
        //     })
        //         .then((response) => {
        //             this.userSession = response.data.user
        //
        //             global_.handleResponseGetThen(response)
        //
        //         })
        //         .catch((error) => {
        //             // this.load_data = false
        //             this.$message.error(error.response.data.message);
        //             // this.errordata = error.response.data
        //         })
        //
        // },

        methods: {},


    }
</script>


<style scoped>

</style>

