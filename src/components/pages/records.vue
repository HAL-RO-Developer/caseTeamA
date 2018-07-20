<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <b-field class="form">
                <b-select placeholder="Select a child" v-model="child_id" @input="getRecords"> 
                    <option v-for="option in options.children" :key="option.child_id" :value="option.child_id">{{option.nickname}}</option>
                </b-select>
                <b-select placeholder="Select a filter" v-model="filter" @input="getRecords"> 
                    <option v-for="option in options.filter" :key="option.value" :value="option.value">{{option.name}}</option>
                </b-select>
                <button class="button message-button is-success" @click="$router.push({path: '/pages/messages'})">
                    <b-icon icon="message-outline"></b-icon>
                </button>
            </b-field>
            
            <by-date ref="date" v-if="filter=='date'" :isLoading="isLoading" @isLoading="isLoading=false"></by-date>
            <by-genre ref="genre" v-if="filter=='genre'" :isLoading="isLoading" @isLoading="isLoading=false"></by-genre>
        </div>
        <fab icon="sync" @click="getRecords"></fab>    
        <app-footer></app-footer>
        <under-tab :index='1'></under-tab>
    </div>
</template>
<script>
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
import Fab from '../modules/fab.vue'
import ByDate from '../modules/records/byDate.vue'
import ByGenre from '../modules/records/byGenre.vue'

export default {
    name :"records",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
        Fab,
        ByDate,
        ByGenre
    },
    data() {
        return {
            title: "記録",
            child_id: "",
            filter: "date",
            records: [],
            options:{
                filter:[
                    {name: '日付別', value: 'date'},
                    {name: '分野別', value: 'genre'}
                ],
                children:[]
            },
            isLoading: false
        }
    },
    methods:{
        getRecords(){
            this.isLoading = true
            localStorage.setItem('child_id', this.child_id)
            http.getRecords(this.child_id,this.filter)
                .then((response)=>{
                    this.records = response.data.records
                    if(this.filter=="date"){
                        this.$refs.date.aggregate(this.records)
                    }else if(this.filter=="genre"){
                        this.$refs.genre.aggregate(this.records)
                    }
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
    created() {
        this.child_id = localStorage.getItem('child_id')
        this.getChild()
        this.getRecords()
    }
}
</script>

<style>
    .contents .form{
        width: auto;
    }

    .message-button{
            left: 3vw;
    }
</style>
