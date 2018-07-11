<template>
    <div class="card">
        <div class="card-content" @click="select">
            <div class="media">
                <div class="media-left">
                    <figure class="image is-48x48">
                        <b-icon :icon="icon.icon" size="is-large" :type="icon.type"></b-icon>
                    </figure>
                </div>
                <div class="media-content">
                    <p class="title is-6">{{detail.sentence}}</p>
                    <p class="subtitle is-6">回答日時：{{date}}</p>   
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import moment from "moment"
    export default {
        data() {
            return {
                icon: {},
                date: null
            }
        },
        methods:{
            select(){
                this.$emit('select')
            }
        },
        
        created(){
            this.icon = ( this.detail.result ) ? {
                icon: "circle-outline",
                type: "is-info"
            } : {
                icon: "close",
                type: "is-danger"
            }
            this.date = moment(this.detail.date).format("YYYY/MM/DD HH:mm:ss")
        },
        props:['detail']
    }
</script>

<style>
    .card{
        position: relative;
    }
    .media-content{
        overflow: auto;
        width: auto;
    }
    .media-content .title{
        white-space: nowrap;
        text-overflow: ellipsis;
    }
    .media-content .subtitle{
        text-align: right;
    }
</style>