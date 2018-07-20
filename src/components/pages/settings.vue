<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <div class="buttons" v-for="item in menu_items" :key="item.id">
                <span class="button is-medium is-success full-width" @click="click(item.id)">
                    <b-icon :icon="item.icon"></b-icon>
                    <span style="padding-left:.5vh;">{{item.title}}</span>
                </span>
            </div>
        </div>
        <under-tab :index='2'></under-tab>
        <app-footer></app-footer>
        
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
                { id:1, title:'子ども一覧', icon:'human-child' },
                { id:2, title:'メッセージ設定', icon:'message-outline' },
                { id:3, title:'BOCCOとつなげる', icon:'robot'},
                //{ id:4, title:'取扱説明書' icon:'note'},
                { id:5, title:'ログアウト', icon:'logout-variant'}
            ]
        }
    },
    methods:{
        click(id){
            switch (id) {
                case 1:
                    this.$router.push({path:"/pages/children"})
                    break;
                case 2:
                    this.$router.push({path:"/pages/messages"})
                    break;
                case 3:
                    this.$router.push({path:"/pages/bocco"})
                    break;
                case 4:
                    this.$router.push({path:"/pages/manual"})
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
    .illust{
        position: fixed;
        right: -3vw;
        bottom: 20vh;
        width: 40vw;
        -webkit-transform: rotate(15deg);
        -ms-transform: rotate(15deg);
        transform: rotate(15deg);
    }
</style>