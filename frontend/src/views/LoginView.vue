<script setup>
import {reactive} from 'vue'
import {useRouter} from 'vue-router'
import {SubmitLogin} from '../../wailsjs/go/main/App'
const router = useRouter()
const data = reactive({
    name: "",
    passwd: "",
    result: 1,
})
function submitForm(data) {
    SubmitLogin(data.name, data.passwd).then(result => {
        console.log("1111result==", result)
        data.result = result
        console.log("2222result==", data.result)
    })
    if (data.result == 0) {
        console.log("result == 0")
    } else {
        console.log("result == 1")
    }
}
</script>

<template>
<div class="login">
    <el-card>
        <h2>欢迎登陆</h2>
        <el-form :model="data" label-width="80px">
        <el-form-item label="用户名称">
            <el-input v-model="data.name"></el-input>
        </el-form-item>
        <el-form-item label="密码">
            <el-input v-model="data.passwd" type="password"></el-input>
        </el-form-item>
        <el-form-item >
            <el-button type="primary" @click="submitForm('data')">确定</el-button>
            <el-button @click="cancel">取消</el-button>
        </el-form-item>
        </el-form>
        <!-- <el-input v-model="data.name" placeholder="用户名" ></el-input>
        <el-input v-model="data.passwd" type="password" placeholder="密码"></el-input>
        <el-button type="primary" @click="login">登陆</el-button> -->
        <!-- <el-input 
        v-model="data.name" 
        placeholder="用户名"
        prefix-icon="fas fa-user"></el-input> -->
    </el-card>
</div>
</template>

<style>
.el-card {
    height: 100%;
}
</style>

