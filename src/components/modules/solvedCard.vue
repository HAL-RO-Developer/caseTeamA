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
                    <p class="title is-6">{{sentence}}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                icon:{}
            }
        },
        methods:{
            select(){
                this.$emit('select')
            }
        },
        
        created(){
            this.icon = ( this.result ) ? {
                icon: "circle-outline",
                type: "is-info"
            } : {
                icon: "close",
                type: "is-danger"
            }
        },
        props:['sentence','result']
    }
</script>