<script setup>
import {reactive} from 'vue'
import {GetRecords} from '../../wailsjs/go/main/App'

const data = reactive({
  records: "",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}
function getRecord() {
  GetRecords().then(result => {
    data.records = result
  })
}
</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>

    <div id="table">
      <table>
        <thead>
          <tr>
            <th></th>
            <th>姓名</th>
            <th>部门</th>
            <th>身份证号</th>
            <th>银行账号</th>
            <th>1月</th>
            <th>2月</th>
            <th>3月</th>
            <th>4月</th>
            <th>5月</th>
            <th>6月</th>
            <th>7月</th>
            <th>8月</th>
            <th>9月</th>
            <th>10月</th>
            <th>11月</th>
            <th>12月</th>
            <th>奖金</th>
            <th>平均</th>
            <th>总计</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="r in data.records">
            <td></td>
            <td>{{r.name}}</td>
            <td>{{r.dep}}</td>
            <td>{{r.id}}</td>
            <td>{{r.bankid}}</td>
            <td><button>编辑</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

table {
	width: 90%;
	background: rgb(39, 36, 36);
	margin: 10px auto;
	border-collapse: collapse;
	/*border-collapse:collapse合并内外边距
(去除表格单元格默认的2个像素内外边距*/
}
th,
td {
	height: 25px;
	line-height: 25px;
	text-align: center;
	border: 1px solid #ccc;
}
th {
	background: rgb(39, 36, 36);
	font-weight: normal;
}
tr {
	background: rgb(39, 36, 36);
}
tr:hover {
	background: #cc0;
}
td a {
	color: #06f;
	text-decoration: none;
}
td a:hover {
	color: #06f;
	text-decoration: underline;
}
</style>
