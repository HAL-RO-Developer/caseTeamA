<template>
    <div style="position:relative;">
        <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
        <graph :chartData='chartData' :options='options' :width="900" :height="680"></graph>
        <b-collapse :open="false">
            <button class="button is-success" style="margin-bottom:1vh;" type="button" slot="trigger">
                <b-icon icon="note"></b-icon>
                <span style="padding-left:.5vh;">グラフの解説</span>
            </button>
            <b-notification type="is-success" :closable="false">
                <h3 class="title is-6"><b>回答率：</b></h3>
                <p class="subtitle is-6">　全問題数あたりの回答数の割合(%)</p>
                <h3 class="title is-6"><b>正答率：</b></h3>
                <p class="subtitle is-6">　回答問題数あたりの正答数の割合(%)</p>
            </b-notification>
        </b-collapse>
        </div>
</template>
<script>
import http from '../../../service/service';
import Graph from './graph/radar.vue'
import Fab from '../../modules/fab.vue'

export default {
    components:{
        Graph,
    },
    data() {
        return {
            child_id: "",
            chartData:{},
            options:{},
            values:{
                genre:[],
                problem:[],
                solved:[],
                correct:[],
                prate:[],
                srate:[]
            },
        }
    },
    methods:{
        fillData () {
            var problem_data = this.values.problem
            var solved_data = this.values.solved
            var correct_data = this.values.correct    
            var prate_data = this.values.prate
            var srate_data = this.values.srate          
            var datasets = [
                {
                    label: '回答率',
                    type: 'radar',
                    data: prate_data,
                    borderColor: "rgba(54,164,235,0.8)",
                    backgroundColor: "rgba(54,164,235,0.2)",
                },
                {
                    label: '正答率',
                    type: 'radar',
                    data: srate_data,
                    borderColor: "rgba(254,97,132,0.8)",
                    backgroundColor: "rgba(254,97,132,0.2)",
                },
                
            ]

            var labels=[];
            this.values.genre.forEach((item)=>{
                labels.push(item.genre_name)
            })

            this.chartData = {
                labels: labels,
                datasets: datasets,
            }

            this.options = {
                scale:{
                    ticks:{
                        min: 0,
                        max: 100,
                        beginAtZero: true
                    }
                },
            }
            this.$emit('isLoading')
        },
        aggregate(records){
            this.records = this.records
            for( var i = 0; i < this.values.genre.length; i++){
                this.values.solved[i] = 0
                this.values.correct[i] = 0
                this.values.srate[i] = 0
                this.values.prate[i] = 0
                records.forEach((record)=>{
                    if(record.genre==this.values.genre[i].genre_name){
                        this.values.problem[i] = Number(record.num_problems)
                        this.values.solved[i] = Number(record.num_ans)
                        this.values.correct[i] = Number(record.num_corr)
                    }
                })
                this.values.srate[i] = Math.round( this.values.correct[i] / this.values.solved[i] * 10000 ) / 100
                this.values.prate[i] = Math.round( this.values.solved[i] / this.values.problem[i] * 10000 ) / 100              
            }
            this.fillData()
        }, 
        getGenre(){
            http.getGenre()
                .then((response)=>{
                    Array.prototype.push.apply(this.values.genre, response.data.genre)
                })
                .catch((err)=>{
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                            icon: 'times-circle',
                            iconPack: 'fa'
                        })
                    }
                })
        }            
    },
    created() {
        this.child_id = localStorage.getItem('child_id')
        this.getGenre()
    },
    props: ["isLoading"]
}
</script>

<style>
</style>
