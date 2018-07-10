<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <p class="modal-card-title">子ども追加</p>
        </header>
        <section class="modal-card-body">
            <b-field label="ニックネーム">
                <b-input
                    type="text"
                    v-model="data.nickname"
                    placeholder="ニックネーム"
                    icon="account"
                    required>
                </b-input>
            </b-field>
            <b-field label="誕生日">
                <b-datepicker
                    placeholder="Click to select..."
                    icon="calendar-today"
                    v-model="data.birthday"
                    required>
                </b-datepicker>
           </b-field>
           <b-field label="性別">
                <b-radio v-model="data.sex" native-value=0>
                    男
                </b-radio>
                <b-radio v-model="data.sex" native-value=1>
                    女
                </b-radio>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">キャンセル</button>
            <button class="button is-primary" @click="add">追加</button>
        </footer>
    </div>
</template>

<script>
    export default {
        data(){
            return{
                data:{
                    device_id: "",
                    nickname: "",
                    birthday: new Date(),
                    sex: 0
                }
            }
        },
        methods:{
            add(){
                this.$emit('add', this.data)
            }
        }
    }
</script>