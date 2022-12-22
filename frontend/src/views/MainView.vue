<template>
    <div>
        <el-container>
        <el-aside width="200px">
            <el-menu
            default-active="1">
                <el-menu-item index="1">
                    <template #title>
                        人员维护
                    </template>
                </el-menu-item>
                <el-menu-item index="2">
                    <template #title>
                        公积金录入
                    </template>
                </el-menu-item>
                <el-menu-item index="3">
                    <template #title>
                        设置
                    </template>
                </el-menu-item>
            </el-menu>
        </el-aside>
        <el-main>Main</el-main>
    </el-container>
    </div>
</template>