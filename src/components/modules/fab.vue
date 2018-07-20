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
        right: 3vw;
        bottom: 20vh;
        display: block;
        width: 55px;
        height: 55px;
        background-color: cyan;
        border-radius: 50%;
        box-shadow:0px 3px 10px rgba(0,0,0, 0.3);
        -webkit-box-shadow: 0px 3px 10px rgba(0,0,0, 0.3);
        transition: transform 0.3s ease-out;
    }
    
    #fab:active{
        -webkit-transform: rotate(-360deg);
        -ms-transform: rotate(-360deg);
        transform: rotate(-360deg);
    }

    #fab .icon{
        display: flex;
        color: white;
        height:100%;
        width: 100%;
        justify-content: center;
        align-items: center;
    }

    
</style>