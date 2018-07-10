<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <button class="button is-medium is-primary full-width" type="button" @click="change">サーバー変更</button>
            <section class="modal-card-body" lebel="現サーバー">
                <b-input
                    type="text"
                    v-model="now_server"
                    placeholder="現在のサーバー"
                    readonly>
                </b-input>
            </section>
            <section class="modal-card-body" label="新サーバー">
                <b-input
                    type="text"
                    v-model="new_server"
                    placeholder="新サーバー">
                </b-input>
            </section>
            <button class="button is-medium is-primary full-width" type="button" @click="$router.push({path: '/admin/question'})">問題作成</button>
        </div>
        <app-footer></app-footer>
    </div>
</template>

<script>
import http from '../../../service/service';
import UnderTab from '../../modules/underTab.vue'
import AppHeader from '../../modules/header.vue'
import AppFooter from '../../modules/footer.vue'

export default {
    name :"admin",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
    },
    data() {
        return {
            title: "admin",
            child_id: "",
            now_server:localStorage.getItem("server"),
            new_server:""
        }
    },
    methods:{
       change(){
            localStorage.setItem("server", this.new_server)
            this.now_server = localStorage.getItem("server")
            this.new_server = ""
            http.Load()
       }
    },
    created(){
    }
}
</script>

<style>
    button{
        display: block;
    }
    .list{
        position: relative;
        overflow: auto;
        height: 43vh;
    }
</style>