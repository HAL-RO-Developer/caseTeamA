<template>
    <div>
        <b-field grouped group-multiline>
            <div class="is-left"><span class="button" @click="prev">＜</span></div>
            <div class="is-center"><b-input type="text" v-model="values.week_range" readonly></b-input></div>
            <div class="is-right"><span class="button" @click="next">＞</span></div>
        </b-field>
        <div style="position:relative;">
            <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
            <graph :chartData='chartData' :options='options' :width="900" :height="750"></graph>
        </div>
    </div>
</template>
<script>
import moment from "moment";
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
            records:[],
            values:{
                week:[],
                week_range:"",
                solved:[],
                correct:[]
            },
        }
    },
    methods:{
        fillData () {
            var solved_data = this.values.solved
            var correct_data = this.values.correct           
            var datasets = [
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
            ]

            this.chartData = {
                labels: this.values.week,
                datasets: datasets                 
            }

            this.options = {
                scales:{
                    yAxes:[
                        {
                            id: "count-axis",
                            position: "left",
                            ticks:{
                                min: 0,
                                max: 30,
                            }
                        }
                    ],
                },
                onClick: 
                    (e, el)=>{
                        this.showSolvedList(e, el)
                    } 
            }
            //this.isLoading = false
            this.$emit('isLoading')
        },
        showSolvedList(e,el){
            if (! el || el.length == 0) return;
            var date = moment(el[0]._model.label,"MM/DD").format("MMDD")
            this.$router.push({path: '/records/date/' + date})
        },
        aggregate(records,date){
            this.records = records
            var m = moment(date).day(1)
            for( var i = 0; i < 7; i++, m.add(1,"day") ){
                this.values.week[i] = m.format('MM/DD')
                this.values.solved[i] = 0
                this.values.correct[i] = 0
                records.forEach((record)=>{
                    if( moment(record.date).isSame(m,'day') ){
                        this.values.solved[i] = Number(record.num_ans)
                        this.values.correct[i] = Number(record.num_corr)
                    }
                })                
            }
            this.values.week_range = this.values.week[0] + " ~ " + this.values.week[i-1]
            this.fillData()
        },
        prev(){
            this.aggregate(this.records,moment(this.values.week[0],'MM/DD').add(-7,'day').toDate())
        },
        next(){
            var date = moment(this.values.week[0],'MM/DD').add(7,'day')
            if(date.isBefore(moment())){
                this.aggregate(this.records,date.toDate())
            }
        },
    },
    created() {
        this.child_id = localStorage.getItem('child_id')
        //this.aggregate()
    },
    
    props: ["isLoading"]
}
</script>

<style>
</style>
