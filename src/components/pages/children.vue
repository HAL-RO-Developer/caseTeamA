"<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
            <card v-for="(child, index) in children" 
                :key="index"
                :selected="selected"
                :child="child"
                @remove="removeChild"
                @select="select"></card>
        </div>
        <fab icon="plus" @click="isComponentModalActive = true"></fab>
        <app-footer></app-footer>
        <under-tab :index='2'></under-tab>
        <b-modal :active.sync="isComponentModalActive" has-modal-card>
            <modal-form @add="addChild"></modal-form>
        </b-modal>
        
    </div>
</template>
<script>
import moment from "moment";
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
import Card from '../modules/children/card.vue'
import ModalForm from '../modules/children/modal.vue'
import Fab from '../modules/fab.vue'

export default {
    name :"children",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
        Card,
        ModalForm,
        Fab
    },
    data() {
        return {
            title: "子ども一覧",
            children: [],
            selected: "",
            isComponentModalActive: false,
            isLoading: false
        }
    },
    methods:{
        getChild(){
            this.isLoading = true
            http.getChild()
                .then((response)=>{
                    this.isLoading = false
                    this.children = []
                    this.children = response.data.children
                })
                .catch((err)=>{
                    this.isLoading = false
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                             icon: 'alert-circle',
                        })
                        switch(err.response.status){
                            case 401:
                                http.RemoveToken()
                                this.$router.push({path:'/'})
                                break;
                            default:
                                break;
                        }
                    }
                });
        },
        addChild(data){
            this.isComponentModalActive = false
            var birthday = moment(data.birthday)
            http.addChild(data.nickname, birthday.format("YYYY-MM-DD"), Number(data.sex))
                .then((response)=>{
                    this.getChild()
                    this.$toast.open(data.nickname+'を追加しました')
                    //this.$router.push({path: '/pages/device' })
                })
                .catch((err)=>{
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                             icon: 'alert-circle',
                        })
                        switch(err.response.status){
                            case 401:
                                http.RemoveToken()
                                this.$router.push({path:'/'})
                                break;
                            default:
                                break;
                        }
                    }
                });
        },
        removeChild(id, nickname){
            this.$dialog.confirm({
                title: '子ども情報削除',
                message: '『'+ nickname +'』を削除しますか？',
                confirmText: '削除',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => 
                    http.removeChild(id)
                        .then((response)=>{
                            var child_id = localStorage.getItem('child_id');
                            if( child_id = id ){
                                localStorage.removeItem('child_id')
                            }
                            this.getChild()
                        })
                        .catch((err)=>{
                            if(err){
                                this.$dialog.alert({
                                    title: 'Error',
                                    message: err.response.data.error,
                                    type: 'is-danger',
                                    hasIcon: true,
                                    icon: 'alert-circle',
                                })
                                switch(err.response.status){
                                    case 401:
                                        http.RemoveToken()
                                        this.$router.push({path:'/'})
                                        break;
                                    default:
                                        break;
                                }
                            }
                        })
            })
            
        },
        select(id){
            localStorage.setItem('child_id', id)
            this.selected = id
            //this.$router.push({path: '/pages/records'}) 
        }
    },
    created(){
        this.selected = localStorage.getItem('child_id');
        this.getChild()
    }
}
</script>

<style>
</style>