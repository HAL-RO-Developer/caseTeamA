<template>
    <div id="fab"  @click="click">
        <b-icon :icon="icon"></b-icon>
    </div>
</template>

<script>
    export default {
    name :"fab",
    components:{
    },
    props:['icon'],
    data() {
        return {
        }
    },
    methods:{
       click(){
            this.$emit('click')
       }
    }
}
</script>

<style>
    #fab {
        position: fixed;
        font-size: 21px;
        line-height: 55px;
        text-align: center;
        color: white;
        background-color: cyan;
        width: 55px;
        height: 55px;
        position:fixed;
        right: 5vw;
        bottom: 10vh;
        border-radius: 50%;
        z-index: 10;
        box-shadow:0px 3px 10px rgba(0,0,0, 0.3);
        -webkit-box-shadow: 0px 3px 10px rgba(0,0,0, 0.3);
        transition: transform 0.3s ease-out;
    }
    
    #fab:active{
        -webkit-transform: rotate(-360deg);
        -ms-transform: rotate(-360deg);
        transform: rotate(-360deg);
    }

    #fab b-icon{
        color: white;
        line-height: 56px;
        vertical-align: middle;
        font-size: 18px;
    }

    
</style>