<template>
  <div class="supply_liquidity">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="supply_liquidity_container d-flex flex-row align-items-center justify-content-start">
            <div class="supply_liquidity_form_container">
              <div class="my-liquidity-title">Supply Liquidity</div>
              <table class="table">
                <thead>
                  <th>Smart Vaults</th>
                  <th>Fee Tier</th>
                  <th>Current APR</th>
                  <th>Capacity</th>
                  <th>TVL</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr>
                    <td>
                      <img src="images/usdc.png"
                           width="30"
                           height="30" />&nbsp;
                      <img src="images/weth.png"
                           width="30"
                           height="30" />&nbsp;
                      <span>USDC / WETH</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">35.1%</button></td>
                    <td>$2,366,149</td>
                    <td><button type="button"
                              class="btn btn-primary btn-sm">Supply</button></td>
                  </tr>
                  <tr>
                    <td>
                      <img src="images/usdc.png"
                           width="30"
                           height="30" />&nbsp;
                      <img src="images/usdt.png"
                           width="30"
                           height="30" />&nbsp;
                      <span>USDC / USDT</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">35.1%</button></td>
                    <td>$2,366,149</td>
                    <td><button type="button"
                              class="btn btn-primary btn-sm">Supply</button></td>
                  </tr>
                  <tr>
                    <td>
                      <img src="images/wbtc.png"
                           width="30"
                           height="30" />&nbsp;
                      <img src="images/usdt.png"
                           width="30"
                           height="30" />&nbsp;
                      <span>WBTC / USDT</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">35.1%</button></td>
                    <td>$2,366,149</td>
                    <td><button type="button"
                              class="btn btn-primary btn-sm">Supply</button></td>
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
  name: 'SupplyLiquidity',
  methods: {

  }
}
</script>
