<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <b-icon :icon="res.icon" :type="res.type"></b-icon>
            <span class="title is-6"> {{res.msg}}</span>
        </header>
        <section class="modal-card-body">
            <p class="title is-6">{{detail.sentence}}</p>
            <p class="is-6">{{nickname}}の答え：{{detail.user_answer}}</p>
            <p class="is-6">解答：{{detail.correct}}</p>
        </section>
        <footer class="modal-card-foot">
            <button class="button is-right" type="button" @click="$parent.close()">OK</button>
        </footer>
    </div>
</template>

<script>
    export default {
        data(){
            return{
                data:{
                    res : {}
                }
            }
        },
        methods:{
        },
        created() {
            this.res = ( this.detail.result ) ? {
                icon: "circle-outline",
                type: "is-info",
                msg: "正解"
            } : {
                icon: "close",
                type: "is-danger",
                msg: "不正解"
            }
        },
        props:['detail', 'nickname']
    }
</script>

<style>
</style>
