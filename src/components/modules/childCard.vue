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
                    <p class="title is-4">{{nickname}}</p>
                    <p class="subtitle is-6">ID:{{id}}</p>
                </div>
                <b-icon v-if="selected==id" icon="check"></b-icon>
                <div class="media-right">
                    <button class="button is-danger" @click="remove">
                        <b-icon icon="delete"></b-icon>
                    </button>
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
            remove(){
                this.$emit('remove', this.id, this.nickname )
            },
            select(){
                this.$emit('select', this.id )
            }
        },
        created(){
            this.icon = ( this.sex ) ? {
                icon: "gender-female",
                type: "is-danger"
            } : {
                icon: "gender-male",
                type: "is-info"
            }
        },
        props:['id','nickname','sex','selected']
    }
</script>