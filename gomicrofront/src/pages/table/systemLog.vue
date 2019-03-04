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
                <!--<el-table-column prop="id" label="id" width=""></el-table-column>-->
                <el-table-column prop="method" label="method" width=""></el-table-column>
                <el-table-column prop="url" label="url" width=""></el-table-column>
                <!--<el-table-column prop="url" label="url" width=""></el-table-column>-->
                <el-table-column prop="remoteaddr" label="remoteaddr" width=""></el-table-column>
                <el-table-column prop="status" label="status" width=""></el-table-column>
                <el-table-column prop="latency" label="latency" width=""></el-table-column>
                <el-table-column prop="duration" label="duration" width=""></el-table-column>
                <el-table-column prop="createTime" label="createTime" width=""></el-table-column>
                <!--<el-table-column prop="latency" label="latency" width=""></el-table-column>-->
                <!--<el-table-column prop="sex" label="性别" width="">-->
                <!--<template slot-scope="scope">-->
                <!--<span v-text="scope.row.sex == 1 ? '男' : '女'"></span>-->
                <!--</template>-->
                <!--</el-table-column>-->
                
                
                <!--<el-table-column label="操作" width="180">-->
                <!--<template slot-scope="scope">-->
                <!---->
                <!--<el-button type="info" size="small"-->
                <!---->
                <!--@click="showPut(scope.row)"-->
                <!--icon="edit">修改-->
                <!--</el-button>-->
                <!---->
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
        
        
        <!--<p>total :{{total}}</p>-->
    
    </div>
</template>


<script type="text/javascript">
    import {bottomToolBar, panelTitle} from 'components'
    import global_ from '@/tool/Global'


    const url = global_.urlPrefix + "/systemLog"


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
                // let this = this
                this.$axios.get(url + '?pageNum=' + pagenum + '&pageSize=' + global_.pageSizeString)
                    .then((response) => {
                        this.load_data = false
                        if (response.data.success === true) {
                            this.table_data = response.data.result
                            // this.pageMap = response.data.page
                            this.total = response.data.page.count
                        }

                        global_.handleResponseGetThen(response)
                    })
                    .catch((error) => {
                        this.load_data = false
                        this.$message.error(error.response.data.message);
                    })
            },


            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            },


        }
    }
</script>
