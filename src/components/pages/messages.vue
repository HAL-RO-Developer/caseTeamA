<template>
   <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>        
        <div class="contents">
            <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
            <b-field>
                <b-select placeholder="Select a child" v-model="child_id" @input="getMessage" expanded> 
                    <option v-for="option in options.children" :key="option.child_id" :value="option.child_id">{{option.nickname}}</option>
                </b-select>
            </b-field>
            <card ref="card" v-for="(message, index) in messages.child_messages" 
                :key="index"
                :condition="message.condition"
                :messageCall="message.message_call"
                :messages="message.messages"
                @remove="removeMessage"
                ></card>
        </div>
        <fab :icon="fabIcon" @click="isComponentModalActive = true"></fab>
        <app-footer></app-footer>
        <under-tab :index='2'></under-tab>
        <b-modal :active.sync="isComponentModalActive" has-modal-card>
            <modal-form @add="addMessage"></modal-form>
        </b-modal>
    </div>
</template>
<script>
    import http from '../../service/service';
    import UnderTab from '../modules/underTab.vue'
    import AppHeader from '../modules/header.vue'
    import AppFooter from '../modules/footer.vue'
    import Card from '../modules/messages/card.vue'
    import ModalForm from '../modules/messages/modal.vue'
    import Fab from '../modules/fab.vue'

    export default {
        name :"messages",
        components:{
            UnderTab,
            AppHeader,
            AppFooter,
            Card,
            ModalForm,
            Fab,
        },
        data() {
            return {
                title: "メッセージ設定",
                child_id: "",
                fabIcon: "plus",
                isComponentModalActive: false,
                isLoading: false,
                messages:[],
                options:{
                    children:[]
                },
            }
        },
        methods:{
            addMessage(data){
                this.isComponentModalActive = false
                http.addMessage(Number(this.child_id), Number(data.condition), Number(data.message_call), data.message )
                    .then((response)=>{
                        this.getMessage()
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
            getMessage(id){
                this.isLoading = true
                if(id!=null){
                    this.child_id = id;
                }
                http.getMessage(this.child_id)
                    .then((response)=>{
                        this.isLoading = false
                        this.messages = response.data.messages
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
            removeMessage(id,message){
                this.$dialog.confirm({
                title: 'メッセージ削除',
                message: '『'+ message +'』を削除しますか？',
                confirmText: '削除',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => 
                    http.removeMessage(id)
                        .then((response)=>{
                            this.getMessage()
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
            getChild(){
            http.getChild()
                .then((response)=>{
                    this.options.children = response.data.children
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
        },
        created(){
            this.child_id = localStorage.getItem('child_id')
            this.getChild()
            this.getMessage()
        }
    }
</script>

<style>
</style>