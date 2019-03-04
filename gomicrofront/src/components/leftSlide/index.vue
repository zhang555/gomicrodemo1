<template>
    <div class="left-side">
        <div class="left-side-inner">
            <router-link to="/" class="logo block">
                <img src="./images/logo.png" alt="AdminX">
            </router-link>
            
            
            <el-menu class="menu-box"
                     router
                     style="border-right-width: 0px;"
                     :default-active="$route.path">
                
                
                <div v-if="user_session.roleName === 'admin' ">
                    <div v-for="(item, index) in nav_menu_data" :key="index">
                        <el-submenu :index="item.title">
                            <template slot="title">
                                
                                <i class="icon" :class="item.icon"></i>
                                <span v-text="item.title" class="text"></span>
                            </template>
                            
                            <el-menu-item class="menu-list"
                                          v-for="(sub_item, sub_index) in item.child"
                                          :index="sub_item.name + sub_item.path"
                                          @click="menuItemClick( sub_item.name  , sub_item.path)"
                                          :key="sub_index">
                                <span v-text="sub_item.title" class="text"></span>
                            </el-menu-item>
                        </el-submenu>
                    </div>
                </div>
                
                <!--
                    v-if="user_session.roleName === 'user' "
                -->
                <div>
                    
                    <!--
                    v-if="item.user === true"
                    -->
                    <div v-for="(item, index) in nav_menu_data" :key="index">
                        <el-submenu :index="item.title">
                            <template slot="title">
                                
                                <i class="icon" :class="item.icon"></i>
                                <span v-text="item.title" class="text"></span>
                            </template>
                            
                            <!--
                                  v-if="sub_item.user === true"
                            -->
                            <el-menu-item class="menu-list"
                                          v-for="(sub_item, sub_index) in item.child"
                                          :index="sub_item.name + sub_item.path"
                                          @click="menuItemClick( sub_item.name  , sub_item.path)"
                                          :key="sub_index">
                                <span v-text="sub_item.title" class="text"></span>
                            </el-menu-item>
                        </el-submenu>
                    </div>
                </div>
            
            
            </el-menu>
        
        
        </div>
    </div>
</template>


<script type="text/javascript">
    import {mapState} from 'vuex';

    export default {
        name: 'slide',
        data() {
            return {

                nav_menu_data: [
                    // {title: "主页", path: "/home", icon: "fa-home"},
                    {
                        title: "table", path: "/table", icon: "el-icon-check",

                        user: true,

                        child: [
                            {title: "user", name: "user", user: false,},

                            {title: "article", name: "article", user: true,},
                            {title: "message", name: "message", user: true,},


                        ]
                    },

                ]
            }
        },

        computed: {
            ...mapState(["user_session"]),
        },

        methods: {

            menuItemClick(name, path) {

                if (name !== undefined) {
                    this.$router.push({name: name,})
                }

                if (path !== undefined) {
                    this.$router.push({path: path,})
                }

                // this.$router.push({path: path, params: {title: title}})
                // this.$router.push({name: 'student', params: {title: title}})

                // this.$router.push({name: 'courseFile', params: {id: row.id}})

            },
        }
    }
</script>
