<template>
  <div class="myliquidity_box">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="myliquidity_box_container d-flex flex-row align-items-center justify-content-start">
            <div class="my_liquidity_container">
              <div class="my-liquidity-title">My Value Locked</div>
              <div class="add-liquidity">
                <el-button type="primary"
                           icon="el-icon-plus">Add Liquidity</el-button>
              </div>
              <table class="table">
                <thead>
                  <th>Smart Vaults</th>
                  <th>Fee Tier</th>
                  <th>Current APR</th>
                  <th>Status</th>
                  <th>Current Value</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr>
                    <td>
                      <img src="images/usdc.png"
                           width="30"
                           height="30" />
                      <img src="images/weth.png"
                           width="30"
                           height="30" />
                      <span>USDC / WETH</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">In range</button></td>
                    <td>$33,500</td>
                    <td>Remove <button type="button"
                              class="btn btn-primary btn-sm"
                              disabled="true">Increase</button></td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MyLiquidity',
  methods: {

  }
}
</script>

<style scoped>
.myliquidity_box {
  width: 100%;
  background: transparent;
}
.my-liquidity-title {
  display: inline;
  font-size: 20px;
  font-weight: bold;
  color: #000000;
  margin-bottom: 20px;
}
.add-liquidity {
  display: inline;
  float: right;
}
.myliquidity_box_container {
  width: calc(100% + 1px);
  left: 0px;
  background: #ffffff;
  margin-top: -100px;
  z-index: 1;
  padding: 50px 66px 0 83px;
  border: 2px solid #000000;
  border-radius: 10px;
}
.my_liquidity_container {
  width: 100%;
}
.my_liquidity_container .table {
  margin-top: 20px;
}
.my_liquidity_container .table th {
  font-size: 14px;
  font-weight: bold;
  color: #000000;
  text-align: left;
}
.my_liquidity_container .table td {
  /* line-height: 50px; */
  font-size: 12px;
  font-weight: bold;
  color: #000000;
  text-align: left;
}
</style>
