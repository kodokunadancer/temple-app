<template lang="pug">
  div
    v-tabs(v-model="tab" centered)
      v-tab(href="#login") ログイン
      v-tab(href="#register") 会員登録
    v-card(width="400px" class="mx-auto mt-5")
      v-card-text
        v-form
          v-text-field(
            prepend-icon="mdi-account-circle" 
            label="名前" 
            v-model="form.name" 
            v-if="tabStatus"
          )
          v-text-field(
            prepend-icon="mdi-email-outline"
            label="メールアドレス" 
            v-model="form.email" 
          )
          v-text-field(
            v-bind:type="showPassword ? 'text' : 'password'"
            @click:append="showPassword = !showPassword"
            prepend-icon="mdi-lock" 
            v-bind:append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            label="パスワード"
            v-model="form.password"
          )
          v-card-actions
            v-btn.info.mx-auto(@click="register" v-if="tabStatus") 登録
            v-btn.info.mx-auto(@click="login" v-if="!tabStatus") ログイン
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'
import * as auth from '../store/modules/auth'
import { namespace } from 'vuex-class'

const Auth = namespace(auth.name)

interface Form {
  name?: string
  email: string
  password: string
}

@Component
export default class LoginComponent extends Vue {
  @Auth.Action setUser
  tab: string = 'login'
  form: Form = {
    name: '',
    email: '',
    password: ''
  }
  showPassword: boolean = false

  get tabStatus(): boolean {
    return this.tab === 'register' ? true : false
  }

  async register() {
    console.log(this.form)
    const response: any = await axios.post('http://localhost:1324/api/signup', {
      params: { form: this.form }
    })
    this.setUser(response.data)
  }

  async login() {
    const response: any = await axios.post('http://localhost:1324/api/login', {
      params: { form: this.form }
    })
    this.setUser(response.data)
  }
}
</script>

<style scoped>
</style>
