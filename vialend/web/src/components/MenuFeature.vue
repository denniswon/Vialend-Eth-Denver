<template>
  <div class="
      header_extra
      d-flex
      flex-row
      align-items-center
      justify-content-start
      ml-auto
    ">
    <div class="book_button trans_200">
      <a href="#"
         @click="centerDialogVisible = true">Launch App</a>
    </div>
    <el-dialog title="DeFi revenue streams on AutoPilot"
               :visible.sync="centerDialogVisible"
               :append-to-body="true"
               width="30%"
               center>
      <span>Maximize youre return on liquidity with SwapStream Smart Vaults.a
        Uniswap V3 concentrated liquidity management tool that will earn revenue
        in the most efficient,secure and automated way.</span>
      <span slot="footer"
            class="dialog-footer">
        <el-button type="primary"
                   @click="lanuchApp">Lanuch App</el-button>
        <el-button @click="centerDialogVisible = false">Learn More</el-button>
        <router-link to="/dashboard">Go to Dashboard</router-link>
        <router-view></router-view>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data () {
    return {
      centerDialogVisible: false
    }
  },
  methods: {
    lanuchApp () {
      this.centerDialogVisible = false
      this.$router.push({ path: '/dashboard' })
      console.log(this.$store.state.name)
    }
  }
}
</script>
