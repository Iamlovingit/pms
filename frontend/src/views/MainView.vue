<template>
    <div>
        <!-- <el-container>
        <el-header width="200px">
            <el-button type="primary">人员维护</el-button>
            <el-button type="primary">公积金维护</el-button>
        </el-header>
        <el-main>Main</el-main>
    </el-container> -->
    <el-tabs type="card" style="height: 200px" class="main-tabs">
        <el-tab-pane label="人员维护">
            <el-button type="primary" @click="addAccount" style="float: left">增加人员</el-button>
            <el-button type="primary" @click="addAccount" style="float: left">刷新</el-button>
            
            <el-table :data="userData" style="width: 100%">
                <el-table-column type="index" ></el-table-column>
                <el-table-column prop="name" label="姓名" ></el-table-column>
                <el-table-column prop="id" label="身份证" ></el-table-column>
                <el-table-column prop="dep" label="部门" ></el-table-column>
                <el-table-column prop="bank" label="银行卡" ></el-table-column>
                <el-table-column prop="percent" label="缴纳率" ></el-table-column>
            </el-table>
        </el-tab-pane>
        <el-tab-pane label="公积金">

        </el-tab-pane>
        <el-tab-pane label="设置">

        </el-tab-pane>
    </el-tabs>
    </div>
</template>
<script setup>
import {reactive} from 'vue'
const userData =reactive([
    {
        name: "zhang3",
        id: "110224198611111111",
        dep: "政府",
        bank: "1234567890",
        percent: "5%",
    },
    {
        name: "li4",
        id: "110224198611111111",
        dep: "政府",
        bank: "1234567890",
        percent: "12%",
    },
])
</script>

<style>
.main-tabs > .el-tabs__content {
  padding: 16px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.el-tabs--right .el-tabs__content,
.el-tabs--left .el-tabs__content {
  height: 100%;
}
</style>