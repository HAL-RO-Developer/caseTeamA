<template>
    <div style="position:relative;">
        <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
        <graph :chartData='chartData' :options='options' :width="900" :height="750"></graph>
    </div>
</template>
<script>
import http from '../../service/service';
import Graph from '../modules/graph.vue'
import Fab from '../modules/fab.vue'

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
                rate:[]
            },
        }
    },
    methods:{
        fillData () {
            var problem_data = this.values.problem
            var solved_data = this.values.solved
            var correct_data = this.values.correct    
            var rate_data = this.values.rate          
            var datasets = [
                /*
                {
                    label: '回答数',
                    type: 'bar',
                    data: solved_data,
                    borderColor: "rgba(254,97,132,0.8)",
                    backgroundColor: "rgba(254,97,132,0.5)",
                    yAxisID: "count-axis"
                },
                {
                    label: '正答数',
                    type: 'line',
                    data: correct_data,
                    borderColor: "rgba(54,164,235,1.0)",
                    pointBackgroundColor: "rgba(54,164,235,1.0)",
                    fill: false,
                    lineTension: 0,
                    yAxisID: "count-axis"
                }
                */
                {
                    label: '正答率',
                    type: 'bar',
                    data: rate_data,
                    borderColor: "rgba(254,97,132,0.8)",
                    backgroundColor: "rgba(254,97,132,0.5)",
                    yAxisID: "count-axis"
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
                scales:{
                    yAxes:[
                        {
                            id: "count-axis",
                            position: "left",
                            ticks:{
                                min: 0,
                                max: 100,
                            }
                        }
                    ],
                },
            }
            this.$emit('isLoading')
        },
        aggregate(records){
            this.records = this.records
            for( var i = 0; i < this.values.genre.length; i++){
                this.values.solved[i] = 0
                this.values.correct[i] = 0
                this.values.rate[i] = 0
                records.forEach((record)=>{
                    if(record.genre==this.values.genre[i].genre_name){
                        this.values.problem[i] = Number(record.num_problem)
                        this.values.solved[i] = Number(record.num_ans)
                        this.values.correct[i] = Number(record.num_corr)
                    }
                })
                this.values.rate[i] = this.values.correct[i] / this.values.solved[i] * 100               
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
