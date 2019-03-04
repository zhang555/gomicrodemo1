<template>
    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!--<el-button @click.stop="on_refresh" size="small">-->
            <!--<i class="fa fa-refresh"></i>-->
            <!--</el-button>-->
            <!--<router-link :to="{name: 'tableAdd'}" tag="span">-->
            <!--<el-button type="primary" icon="plus" size="small">添加数据</el-button>-->
            <!--</router-link>-->
        </panel-title>
        <div class="panel-body">
            
            
            <el-row>
                <el-button @click.stop="get_table_data" size="small" type="primary">
                    <i class="el-icon-refresh"></i>
                    <span type="text">刷新</span>
                </el-button>
                
                <!--<el-button @click.stop="showPost" size="small" type="primary">-->
                    <!--<i class="el-icon-plus "></i><span type="text">创建</span>-->
                <!--</el-button>-->
                <br>
                <br>
            </el-row>
            
            
            <el-table
                :data="table_data"
                v-loading="load_data"
                element-loading-text="拼命加载中"
                border
                style="width: 100%;">
                
                <!--<el-table-column type="selection" width="55"></el-table-column>-->
                <el-table-column prop="id" label="id" width=""></el-table-column>
                <el-table-column prop="userId" label="userId" width=""></el-table-column>
                <el-table-column prop="message" label="message" width=""></el-table-column>
                <!--<el-table-column prop="content" label="content" width=""></el-table-column>-->
                <!--<el-table-column prop="userId" label="userId" width=""></el-table-column>-->
                <!--<el-table-column prop="" label="nickname" width="">-->
                    
                    <!--<template slot-scope="scope">-->
                        <!--<span>{{getNickname(scope.row.userId)}}</span>-->
                    <!--</template>-->
                <!--</el-table-column>-->
                
                <!--<el-table-column prop="sex" label="性别" width="">-->
                <!--<template slot-scope="scope">-->
                <!--<span v-text="scope.row.sex == 1 ? '男' : '女'"></span>-->
                <!--</template>-->
                <!--</el-table-column>-->
                
                <el-table-column prop="createTime" label="创建时间" width="200"></el-table-column>
                <el-table-column prop="updateTime" label="修改时间" width="200"></el-table-column>
                
                <!--<el-table-column label="操作" width="180">-->
                    <!--<template slot-scope="scope">-->
                        <!--<router-link :to="{name: 'tableUpdate', params: {id: scope.row.id}}" tag="span">-->
                        <!--<el-button type="info" size="small" icon="edit">修改</el-button>-->
                        <!--</router-link>-->
                
                        <!--
                          @click="showViewDialog(scope.row)"
                        -->
                        <!--<el-button type="info" size="small"-->
                                   <!--@click="showPut(scope.row)"-->
                                   <!--icon="edit">修改-->
                        <!--</el-button>-->
                        
                        <!--<el-button type="danger" size="small" icon="delete" @click="delete_data(scope.row)">删除-->
                        <!--</el-button>-->
                    <!--</template>-->
                <!--</el-table-column>-->
            </el-table>
            
            <el-pagination
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-size="10"
                layout="total, prev, pager, next"
                :total="total">
            </el-pagination>
        
        </div>
        
        
        <el-dialog :title="dialogTypeName" :visible.sync="dialogFormVisible">
            <el-form :model="form" :rules="rules" ref="form" label-width="100px">
                
                <el-form-item label="主题" prop="title">
                    <el-input v-model.trim="form.title"></el-input>
                </el-form-item>
                
                <el-form-item label="主要内容" prop="content">
                    <el-input type="textarea"
                              :autosize="{ minRows: 5,}"
                              v-model="form.content"></el-input>
                </el-form-item>
            </el-form>
            
            <div slot="footer" class="dialog-footer" v-if="dialogType === 1 || dialogType === 2">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <!--<el-button type="primary" @click="postOrPut()">确 定</el-button>-->
                <el-button type="primary" @click="submitForm('form')">确 定</el-button>
            
            </div>
        
        </el-dialog>
        
        <!--<p>total :{{total}}</p>-->
    
    </div>
</template>


<script type="text/javascript">
    import {bottomToolBar, panelTitle} from 'components'
    import global_ from '@/tool/Global'
    import {mapMutations, mapState} from 'vuex';
    
    const url = global_.urlPrefix + "/rpc"

    export default {
        data() {
            return {
                table_data: null,
                //当前页码
                currentPage: 1,
                //数据总条目
                total: 0,
                //每页显示多少条数据
                length: 15,
                //请求时的loading效果
                load_data: true,
                //批量选择数组
                batch_select: [],
                
                dialogFormVisible: false,
                dialogTypeName: '修改',  //创建
                dialogType: 1,
                
                form: {},

                rules: {
                    // title: [
                    //     {required: true, message: '不能为空', trigger: 'blur'},
                    //     {type: 'string', message: '长度需要在1~50之间', min: 1, max: 50}
                    // ],
                    // content: [
                    //     {required: true, message: '不能为空', trigger: 'blur'},
                    //     {type: 'string', message: '长度需要在1~400之间', min: 1, max: 400}
                    // ],
                },

                users: [],

            }
        },
        components: {
            panelTitle,
            bottomToolBar
        },
        created() {
            this.get_table_data()
        },
        methods: {
            ...mapMutations([
                'set_user_session',
                'set_token',
            ]),

            //刷新
            on_refresh() {
                this.get_table_data()
            },
            //获取数据
            get_table_data() {
                this.get_method(this.currentPage)
            },

            get_method(pagenum) {
                this.load_data = true

                this.$axios.post(url, {
                    service: "go.micro.srv.message",
                    method: "MessageService.GetMany",
                    request: {
                        PageNum: pagenum,
                        PageSize: 10,
                    },
                })
                    .then((response) => {
                        this.load_data = false

                        this.table_data = response.data.messages
                        this.total = response.data.page.count
                        
                        global_.handleResponseGetThen(response)


                        let set = new Set()
                        
                        for (let i = 0, len = this.table_data.length; i < len; i++) {
                            set.add(this.table_data[i].userId)
                        }
                        
                        this.$axios.post(url, {
                            service: "go.micro.srv.user",
                            method: "UserService.GetNicknamesByIds",
                            request: {
                                ids: [...set],
                            },
                        })
                            .then((response) => {
                                // this.load_data = false
                                this.users = response.data.users
                                // this.total = response.data.page.count
                                // global_.handleResponseGetThen(response)
                            })
                            .catch((error) => {
                                this.$message.error(error.response.data.msg);
                            })


                    })
                    .catch((error) => {
                        this.load_data = false
                        // this.$message.error(error.response.data.message);
                        // this.errordata = error.response.data
                    })
            },


            showPost(data) {
                this.dialogType = 1
                this.dialogTypeName = '创建'
                this.form = {}

                this.dialogFormVisible = true

            },

            showPut(data) {
                this.dialogType = 2
                this.dialogTypeName = '修改'

                let bean = {
                    service: "go.micro.srv.article",
                    method: "ArticleService.GetOne",
                    request: {
                        id: data.id,
                    },
                }

                this.$axios.post(url, bean)
                    .then((response) => {
                        console.log(response)

                        if (response.data.code === 200) {
                            this.form = JSON.parse(JSON.stringify(response.data.article))
                            this.dialogFormVisible = true
                        }
                    })
                    .catch((error) => {
                        console.log(error)
                    })
            },

            postOrPut() {
                this.dialogFormVisible = false;

//                post
                if (this.dialogType === 1) {
                    // var str = JSON.stringify(this.form)

                    let bean = {
                        service: "go.micro.srv.article",
                        method: "ArticleService.Post",
                        request: this.form,
                    }

                    this.$axios.post(url, bean)
                        .then((response) => {
                            global_.handleResponseThenCustomizePost(response)
                            // global_.handleResponseThenCustomizePost(response)
                            // self.$message.success(response.data.message);
                            // console.log(response)

                            this.get_table_data()
                            console.log(response)

                        })
                        .catch((error) => {
                            // this.$message.error(error);
                            console.log(error)
                            // console.log(response)
                            global_.handleResponseCatch()
                        })
                }

                // put
                if (this.dialogType === 2) {

                    let form = JSON.parse(JSON.stringify(this.form))
                    delete form.createTime
                    delete form.updateTime

                    let bean = {
                        service: "go.micro.srv.article",
                        method: "ArticleService.Put",
                        request: form,
                    }

                    this.$axios.post(url, bean)
                        .then((response) => {
                            global_.handleResponseThenCustomizePut(response)
                            console.log(response)

                            this.get_table_data()
                            // self.$message.success(response.data.message);
                        })
                        .catch((error) => {
                            console.log(error)

                            global_.handleResponseCatch()
                        })
                }
            },


            submitForm(formName) {
                let self = this
                this.$refs[formName].validate((valid) => {
                    if (valid) {

                        self.postOrPut()

                    } else {
                        return false;
                    }
                });
            },


            //单个删除
            delete_data(item) {
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                })
                    .then(() => {

                        let bean = {
                            service: "go.micro.srv.article",
                            method: "ArticleService.Delete",
                            request: {
                                id: item.id,
                            },
                        }

                        this.$axios.post(url, bean)
                            .then((response) => {
                                global_.handleResponseThenCustomizeDelete(response)
                                this.get_table_data()
                                // self.$message.success(response.data.message);
                            })
                            .catch((error) => {
                                global_.handleResponseCatch()
                                // self.get_table_data()
                                //  self.$message.error(error.response.data.message);
                            })

                    })
                    .catch(() => {
                    })
            },


            getNickname(id) {
                for (let i = 0, len = this.users.length; i < len; i++) {
                    if (this.users[i].id === id) {
                        return this.users[i].nickname
                    }
                }
            },

            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            },

        }
    }
</script>
