<template>
    <div class="card">
        <div class="card-content" @click="select">
            <div class="media">
                <div class="media-left">
                    <figure class="image is-48x48">
                        <b-icon 
                            :icon="this.child.sex ? 'gender-female'  : 'gender-male'"
                            :type="this.child.sex ? 'is-danger'      : 'is-info'"
                            size="is-large">    
                        </b-icon>
                    </figure>
                </div>
                <div class="media-content">
                    <p class="title is-4">{{child.nickname}}</p>
                </div>
                <b-icon v-if="selected==child.child_id" icon="check"></b-icon>
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
                this.$emit('remove', this.child.child_id, this.child.nickname )
            },
            select(){
                this.$emit('select', this.child.child_id )
            }
        },
        created(){
        },
        props:['selected','child']
    }
</script>