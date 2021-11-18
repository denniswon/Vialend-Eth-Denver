<template>
  <div class="token_exchange_form">
    <div style="margin-bottom:20px;">Currency Converter</div>
    <el-form :inline="true"
             class="demo-form-inline">
      <el-form-item>
        <el-input v-model="currTokenVal"
                  placeholder="1" />
      </el-form-item>
      <el-form-item>
        <el-select v-model="currTokenId"
                   placeholder="请选择"
                   @change="exchangeTokenChange">
          <el-option v-for="item in tokensList"
                     :key="item.nameid"
                     :label="item.symbol"
                     :value="item.symbol" />
        </el-select>
      </el-form-item>
      =
      <el-form-item>
        <el-input v-model="usdTokenVal" />
      </el-form-item>USD
    </el-form>
  </div>
</template>

<script>
export default {
  data () {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
        currTokenId: 0,
        currTokenVal: ''
      }
    }
  },
  methods: {
    onSubmit () {
      this.$message('submit!')
    },
    onCancel () {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>
