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
                         height="33" />&nbsp;&nbsp;Vialend
                  </div>
                  <div></div>
                </a>
              </div>
              <nav class="main_nav">
                <ul class="d-flex flex-row align-items-center justify-content-start">
                  <li><a href="blog.html">Docs</a></li>
                  <li><a href="contact.html">Vote</a></li>
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
                     @click="centerDialogVisible = true">Launch App</a>
                </div>
                <el-dialog title="Automated Uniswap V3 Vault"
                           :visible.sync="centerDialogVisible"
                           :append-to-body="true"
                           width="650px"
                           center>
                  <div class="dialog_intro">Earn higher yield by automating concentrated liquidity while<br> avoiding permanent loss. Using a unique combination of<br> DeFi lending Pools and Uniswap V3.</div><br>
                  <div class="block">
                    <span class="intro_icon"><img src="images/token_icon.png"
                           width="30" /></span>
                    <span class="intro_cnt">Keep your tokens earning interest in both lending pools & V3</span>
                  </div>
                  <div class="block">
                    <span class="intro_icon">$</span>
                    <span class="intro_cnt">Automatically rebalance with no swap fees</span>
                  </div>
                  <div class="block">
                    <span class="intro_icon">50/50</span>
                    <span class="intro_cnt">Maintain a "50-50 v2 style" position in v3 while avoiding permanent loss</span>
                  </div>
                  <span slot="footer"
                        class="dialog-footer">
                    <el-button type="primary"
                               @click="lanuchApp">Lanuch App</el-button>
                    <el-button @click="centerDialogVisible = false">Learn More</el-button>
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
      <div class="background_image"></div>
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
          <div class="row">
            <div class="col">
              <div class="home_content text-center">
                <div class="home_title">
                  <h1>Welcome to Vialend</h1>
                </div>
                <div class="home_text">It's your best choice</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Component, Vue } from 'vue-property-decorator'

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
      // console.log(this.$store.state.name)
    }
  }
}
</script>
<style scoped>
.dialog_intro {
  font-size: 20px;
  text-align: center;
  color: black;
}
.dialog_detail {
  font-size: 14px;
  line-height: 36px;
}
.dialog_detail li {
}
.dialog_detail li span {
  font-size: 20px;
}
.block {
  display: block;
  padding: 10px 50px;
  overflow: hidden;
}
.intro_icon {
  font-size: 16px;
  color: black;
  line-height: 20px;
  font-weight: bold;
}
.intro_cnt {
  float: right;
  width: 85%;
  margin-right: 20px;
  font-size: 12px;
}
</style>
