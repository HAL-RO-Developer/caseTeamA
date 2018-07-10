<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <div class="buttons" v-for="item in menu_items" :key="item.id">
                <span class="button full-width" @click="click(item.id)">{{item.title}}</span>
            </div>
        </div>
        <app-footer></app-footer>
        <under-tab :index='2'></under-tab>
    </div>
</template>
<script>
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'

export default {
    name :"settings",
    components:{
        UnderTab,
        AppHeader,
        AppFooter
    },
    data() {
        return {
            title:"設定",
            menu_items: [
                { id:1, title:'子ども一覧' },
                { id:2, title:'メッセージ設定' },
                { id:3, title:'取扱説明書'},
                { id:4, title:'ログアウト' }
            ]
        }
    },
    methods:{
        click(id){
            switch (id) {
                case 1:
                    this.$router.push({path:"/children"})
                    break;
                case 2:
                    this.$router.push({path:"/messages"})
                    break;
                case 3:
                    this.$router.push({path:"/manual"})
                    break;
                default:
                    http.RemoveToken()
                    this.$router.push({path:"/"})
                    break;
            }
        }
    }
}
</script>

<style>
</style>