<script setup>
import {reactive} from 'vue'
import {useRouter} from 'vue-router'
import {SubmitLogin} from '../../wailsjs/go/main/App'
const router = useRouter()
const data = reactive({
    name: "",
    passwd: "",
    loginResult: 2,
})
function submitForm() {
    SubmitLogin(data.name, data.passwd).then( (result) => {
        console.log("result=", result)
        data.loginResult = result
        if (data.loginResult == 0) {
            router.push("/main")
        }
    });
  
}
</script>

<template>
<div class="login">
    <el-container>
        <el-main>
            <el-alert title="用户名或密码错误" type="error" v-if="data.loginResult==1" />
            <h2>欢迎登陆</h2>
            <el-form :model="data" label-width="80px">
                <el-form-item label="用户名称">
                    <el-input v-model="data.name" required="required"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="data.passwd" type="password" required="required"></el-input>
                </el-form-item>
                <el-form-item >
                    <el-button type="primary" @click="submitForm()">确定</el-button>
                    <el-button @click="cancel">取消</el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
</div>
</template>

<style>
</style>

