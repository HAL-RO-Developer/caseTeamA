<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <div class="lavel">
                <div class="level-center">
                    <div class="level-item">
                        <p class="title is-5">
                            {{date}}の回答記録
                        </p>
                    </div>
                </div>
            </div>
            <b-field>
                <b-select placeholder="Select a genre" v-model="genre" @input="getDetail(child_id,date,genre)" expanded>
                    <option
                        v-for="option in options.genre"
                        :value="option.id"
                        :key="option.id">
                        {{ option.name }}
                    </option>
                </b-select>
            </b-field>
           <div style="position:relative;">
                <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
                <card v-for="(detail, index) in records.detail" 
                    :key="index"
                    :sentence="detail.sentence" 
                    :result="detail.result"
                    @select="showDatail(index)">
                </card>
            </div>
        </div>
        <app-footer></app-footer>
        <under-tab :index='1'></under-tab>
        <b-modal :active.sync="isComponentModalActive" has-modal-card>
            <modal-form :detail="modalData"></modal-form>
        </b-modal>
    </div>
</template>
<script>
import moment from "moment"
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
import Card from '../modules/solvedCard.vue'
import ModalForm from '../modules/detailModal.vue'

export default {
    name :"solved-list",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
        Card,
        ModalForm
    },
    data() {
        return {
            title: "記録",
            child_id: "",
            date: "",
            genre: "",
            records: {},
            isComponentModalActive: false,
            isLoading: false,
            modalData:{},
            options: {
                genre:[
                    {id: null, name: '全て'},
                    {id: 1, name: '算数'},
                    {id: 2, name: '社会'}
                ]
            }
        }
    },
    methods:{  
        getDetail(child_id, date, genre){
            this.isLoading = true
            http.getDetail(child_id, date, genre)
                .then((response)=>{
                    console.log(response.data)
                    this.isLoading = false
                    this.records = response.data.records
                })
                .catch((err)=>{
                    this.isLoading = false
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                            icon: 'times-circle',
                            iconPack: 'fa'
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
        showDatail(index){
            this.modalData = this.records.detail[index]
            this.isComponentModalActive = true
        }
    },
    created() {
        this.child_id = localStorage.getItem('child_id')
        this.date = moment(this.$route.params.date,'MMDD').format('YYYY-MM-DD')
        this.getDetail(this.child_id,this.date)
    }
}
</script>

<style>
    
</style>