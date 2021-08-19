<template>
  <div>
    <!-- Header -->
    <header class="header">
      <div class="container">
        <div class="row">
          <div class="col">
            <div class="
                header_content
                d-flex
                flex-row
                align-items-center
                justify-content-start
              ">
              <div class="logo">
                <a href="#">
                  <div style="vertical-align: middle;">
                    <img src="images/logo.png"
                         width="30"
                         height="33" />&nbsp;&nbsp;SwapStream
                  </div>
                  <div></div>
                </a>
              </div>
              <nav class="main_nav">
                <ul class="d-flex flex-row align-items-center justify-content-start">
                  <li><a href="#">Dashboard</a></li>
                  <li><a href="#">Vote</a></li>
                </ul>
              </nav>
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
                     @click="centerDialogVisible = true">Connect Wallet</a>
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
              <div class="hamburger ml-auto">
                <i class="fa fa-bars"
                   aria-hidden="true"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Menu -->

    <div class="menu">
      <div class="background_image"
           style="background-image: url(images/menu.jpg)"></div>
      <div class="
          menu_content
          d-flex
          flex-column
          align-items-center
          justify-content-center
        ">
        <ul class="menu_nav_list text-center">
          <li><a href="index.html">Docs</a></li>
          <li><a href="about.html">Vote</a></li>
        </ul>
        <div class="menu_review"><a href="#">Launch App</a></div>
      </div>
    </div>

    <div class="home">
      <div class="parallax_background parallax-window"
           data-parallax="scroll"
           data-speed="0.8"></div>
      <div class="home_container">
        <div class="container">
          <div class="row align-items-center my-financial">
            <div class="col-md-3 text-center">
              <ul>
                <li class="title">My Liquidity</li>
                <li class="value">$0.000000000000000</li>
              </ul>
            </div>
            <div class="col-md-6 text-center apy-container">
              <div class="apy-style">
                <ul>
                  <li class="apy-title">Net APY</li>
                  <li class="apy-content">...</li>
                </ul>
              </div>
            </div>
            <div class="col-md-3 text-center">
              <ul>
                <li class="title">My Revenue</li>
                <li class="value">$0.00000000000000</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
    <MyLiquidity />
    <SupplyLiquidity />
  </div>
</template>

<script>
import { Component, Vue } from 'vue-property-decorator'
import MyLiquidity from '@/components/MyLiquidity.vue'
import SupplyLiquidity from '@/components/SupplyLiquidity.vue'

export default {
  components: { MyLiquidity, SupplyLiquidity },
  data () {
    return {
      centerDialogVisible: false
    }
  },
  methods: {
    lanuchApp () {
      this.centerDialogVisible = false
      this.$router.push({ path: '/dashboard' })
      // console.log(this.$store.state.name)
    }
  }
}
</script>

<style scoped>
.my-financial ul li {
  line-height: 30px;
}
.my-financial .title {
  font-size: 16px;
}
.my-financial .value {
  font-size: 20px;
  color: white;
}
.my-financial .apy-title {
  margin-top: 50px;
  font-size: 20px;
  color: white;
}
.my-financial .apy-content {
  font-size: 20px;
  color: white;
}
.apy-container {
  position: relative;
}
.apy-style {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 200px;
  height: 200px;
  border: 5px solid #9900ff;
  border-radius: 50%;
}
</style>
